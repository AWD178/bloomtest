package bloom

import (
	"bloom/internal/store"
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
}

func TestSimpleMapSearcher_GetStoreByName(t *testing.T) {

	tests := []struct {
		name      string
		stores    []store.Store
		storeName string
		want      []store.Store
		wantErr   error
	}{
		{
			name: "test 1",
			stores: []store.Store{
				{
					Name: "store 1",
					Geo: store.Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 2"},
				},
				{
					Name: "store 2",
					Geo: store.Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 404"},
				},
			},
			storeName: "store 1",
			want: []store.Store{
				{
					Name: "store 1",
					Geo: store.Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 2"},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.NewStore(tt.stores)
			got, err := s.GetStoreByName(tt.storeName)

			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("SimpleMapSearcher.GetStoreByCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleMapSearcher.GetStoreByCategories() = %v, want %v", got, tt.want)
			}

		})
	}
	PrintMemUsage()
}

func TestSimpleMapSearcher_GetStoreByGeo(t *testing.T) {

	tests := []struct {
		name    string
		stores  []store.Store
		geo     store.Geo
		want    []store.Store
		wantErr error
	}{
		{
			name: "test 1",
			stores: []store.Store{
				{
					Name: "store 1",
					Geo: store.Geo{
						Latitude:  0.6,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 2"},
				},
				{
					Name: "store 2",
					Geo: store.Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 404"},
				},
			},
			geo: store.Geo{
				Latitude:  0.5,
				Longitude: 0.5,
			},
			want: []store.Store{
				{
					Name: "store 2",
					Geo: store.Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 404"},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.NewStore(tt.stores)
			got, err := s.GetStoreByGeo(tt.geo)

			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("SimpleMapSearcher.GetStoreByCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleMapSearcher.GetStoreByCategories() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestSimpleMapSearcher_GetStoreByCategories(t *testing.T) {

	tests := []struct {
		name       string
		stores     []store.Store
		categories []string
		want       []store.Store
		wantErr    error
	}{
		{
			name: "test 1",
			stores: []store.Store{
				{
					Name: "store 1",
					Geo: store.Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 2"},
				},
				{
					Name: "store 2",
					Geo: store.Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 404"},
				},
			},
			categories: []string{"category 1", "category 2"},
			want: []store.Store{
				{
					Name: "store 1",
					Geo: store.Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 2"},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.NewStore(tt.stores)
			got, err := s.GetStoreByCategories(tt.categories)

			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("SimpleMapSearcher.GetStoreByCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleMapSearcher.GetStoreByCategories() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestSimpleMapSearcher_GetStoreByCategories2(t *testing.T) {

	tests := []struct {
		name       string
		stores     []store.Store
		categories []string
		want       []store.Store
		wantErr    error
	}{
		{
			name: "test 1",
			stores: []store.Store{
				{
					Name: "store 1",
					Geo: store.Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 2", "cat"},
				},
				{
					Name: "store 2",
					Geo: store.Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 404", "cat"},
				},
			},
			categories: []string{"category 1", "cat"},
			want: []store.Store{
				{
					Name: "store 1",
					Geo: store.Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 2", "cat"},
				},
				{
					Name: "store 2",
					Geo: store.Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 404", "cat"},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.NewStore(tt.stores)
			got, err := s.GetStoreByCategories(tt.categories)

			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("SimpleMapSearcher.GetStoreByCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleMapSearcher.GetStoreByCategories() = %v, want %v", got, tt.want)
			}

		})
	}
}
