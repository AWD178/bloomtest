package store

import (
	"bloom/pkg/bloom"
	"fmt"
	"hash/crc32"
	"strings"
)

type StoreSearcher struct {
	stores []Store
	Index  *bloom.Index
}

const RateCount = 128

func NewStore(stores []Store) *StoreSearcher {
	s := &StoreSearcher{
		stores,
		bloom.NewIndex(len(stores)*RateCount, len(stores)*RateCount, 64),
	}

	for _, v := range s.stores {
		catKey := strings.ReplaceAll(strings.Join(v.Categories, ""), " ", "")
		geoKey := fmt.Sprintf("%.4f_%.4f", v.Geo.Latitude, v.Geo.Latitude)
		resultKey := strings.ReplaceAll(v.Name, " ", "") + " " + catKey + " " + geoKey
		tokens := make([]uint32, 0)
		for _, t := range strings.Fields(resultKey) {
			tokens = append(tokens, crc32.ChecksumIEEE([]byte(t)))
		}
		s.Index.AddDocument(tokens)
	}

	return s
}

func (s *StoreSearcher) GetStoreByName(name string) ([]Store, error) {
	tokens := []uint32{
		crc32.ChecksumIEEE([]byte(strings.ReplaceAll(name, " ", ""))),
	}

	ids := s.Index.Query(tokens)
	if len(ids) == 0 {
		return nil, ErrNotFound
	}

	result := make([]Store, 0)
	for _, id := range ids {
		//if s.stores[id].Name == name {
		result = append(result, s.stores[id])
		//}
	}

	return result, nil
}

func (s *StoreSearcher) GetStoreByGeo(geo Geo) ([]Store, error) {
	tokens := []uint32{
		crc32.ChecksumIEEE([]byte(fmt.Sprintf("%.4f_%.4f", geo.Latitude, geo.Latitude))),
	}

	ids := s.Index.Query(tokens)
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
	catKey := strings.ReplaceAll(strings.Join(cat, ""), " ", "")
	tokens := []uint32{
		crc32.ChecksumIEEE([]byte(catKey)),
	}

	ids := s.Index.Query(tokens)
	if len(ids) == 0 {
		return nil, ErrNotFound
	}

	result := make([]Store, 0)
	for _, id := range ids {
		for _, c := range cat {
			if inArray(s.stores[id].Categories, c) != -1 {
				result = append(result, s.stores[id])
				break
			}
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
