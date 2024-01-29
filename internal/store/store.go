package store

import (
	"bloom/pkg/bloom"
	"fmt"
	"hash/crc32"
	"strings"
)

type StoreSearcher struct {
	stores    []Store
	NameIndex *bloom.Index
	CatIndex  *bloom.Index
	GeoIndex  *bloom.Index
}

const RateCount = 12

func NewStore(stores []Store) *StoreSearcher {
	s := &StoreSearcher{
		stores,
		bloom.NewIndex(len(stores)*RateCount, len(stores)*RateCount, 12),
		bloom.NewIndex(len(stores)*RateCount, len(stores)*RateCount, 12),
		bloom.NewIndex(len(stores)*RateCount, len(stores)*RateCount, 12),
	}

	for _, v := range s.stores {
		geoKey := fmt.Sprintf("%.4f %.4f", v.Geo.Latitude, v.Geo.Latitude)

		tokensCat := make([]uint32, 0)
		tokensName := make([]uint32, 0)
		tokensGeo := make([]uint32, 0)

		for _, t := range v.Categories {
			tokensCat = append(tokensCat, crc32.ChecksumIEEE([]byte(t)))
		}

		for _, t := range strings.Fields(v.Name) {
			tokensName = append(tokensName, crc32.ChecksumIEEE([]byte(t)))
		}

		for _, t := range strings.Fields(geoKey) {
			tokensGeo = append(tokensGeo, crc32.ChecksumIEEE([]byte(t)))
		}

		s.CatIndex.AddDocument(tokensCat)
		s.NameIndex.AddDocument(tokensName)
		s.GeoIndex.AddDocument(tokensGeo)

	}

	return s
}

func (s *StoreSearcher) GetStoreByName(name string) ([]Store, error) {
	tokens := []uint32{
		crc32.ChecksumIEEE([]byte(name)),
	}

	ids := s.NameIndex.Query(tokens)
	if len(ids) == 0 {
		return nil, ErrNotFound
	}

	result := make([]Store, 0)
	for _, id := range ids {
		if s.stores[id].Name == name {
			result = append(result, s.stores[id])
		}
	}

	return result, nil
}

func (s *StoreSearcher) GetStoreByGeo(geo Geo) ([]Store, error) {
	tokens := []uint32{
		crc32.ChecksumIEEE([]byte(fmt.Sprintf("%.4f %.4f", geo.Latitude, geo.Latitude))),
	}

	ids := s.GeoIndex.Query(tokens)
	if len(ids) == 0 {
		return nil, ErrNotFound
	}

	result := make([]Store, 0)
	for _, id := range ids {
		if s.stores[id].Geo.Latitude == geo.Latitude && s.stores[id].Geo.Longitude == geo.Longitude {
			result = append(result, s.stores[id])
		}
	}

	return result, nil
}

func (s *StoreSearcher) GetStoreByCategories(cat []string) ([]Store, error) {
	tokens := make([]uint32, 0)
	for _, c := range cat {
		tokens = append(tokens, crc32.ChecksumIEEE([]byte(c)))
	}
	ids := s.CatIndex.Query(tokens)
	if len(ids) == 0 {
		return nil, ErrNotFound
	}

	result := make([]Store, 0)
	for _, id := range ids {
		idx := 0
		for _, c := range cat {
			if inArray(s.stores[id].Categories, c) != -1 {
				idx++
			}
		}
		if idx == len(cat) {
			result = append(result, s.stores[id])
		}
	}

	return result, nil
}

func inArray(stack []string, need string) int {
	for index, s := range stack {
		if s == need {
			return index
		}
	}
	return -1
}
