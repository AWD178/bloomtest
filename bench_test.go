package bloom

import (
	"bloom/internal/store"
	"math/rand"
	"strconv"
	"testing"
)

const totalStores = 10000

func BenchmarkSimpleMapSearcher_GetStoreByName(b *testing.B) {

	stores := generateStores(totalStores)
	var found []store.Store
	var err error
	s := store.NewStore(stores)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found, err = s.GetStoreByName("store1")
	}

	_ = err

	_ = found
	PrintMemUsage()
}

func BenchmarkSimpleMapSearcher_GetStoreByGeo(b *testing.B) {

	stores := generateStores(totalStores)
	var found []store.Store
	var err error
	s := store.NewStore(stores)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found, err = s.GetStoreByGeo(store.Geo{Latitude: 0.5, Longitude: 0.5})
	}

	_ = err

	_ = found
	PrintMemUsage()
}

func BenchmarkSimpleMapSearcher_GetStoreByCategories(b *testing.B) {

	stores := generateStores(totalStores)
	var found []store.Store
	var err error

	categories := randCategories()

	s := store.NewStore(stores)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found, err = s.GetStoreByCategories(categories)
	}

	_ = err

	_ = found
	PrintMemUsage()
}

func generateStores(i int) []store.Store {

	stores := make([]store.Store, i)

	for j := 0; j < i; j++ {
		stores[j] = randSore()
	}

	return stores
}

func randSore() store.Store {
	return store.Store{
		Name:       randName(),
		Geo:        randGeo(),
		Categories: randCategories(),
	}
}

func randName() string {
	return "store " + strconv.Itoa(rand.Intn(1000000))
}

func randCountCategories() int {
	return rand.Intn(10) + 1
}

func randCategoryName() string {
	return "category " + strconv.Itoa(rand.Intn(1000))
}

func randCategories() []string {
	count := randCountCategories()
	categories := make([]string, count)
	for i := 0; i < count; i++ {
		categories[i] = randCategoryName()
	}
	return categories
}

func randGeo() store.Geo {
	return store.Geo{
		Latitude:  rand.Float32() + 1,
		Longitude: rand.Float32() + 3,
	}
}
