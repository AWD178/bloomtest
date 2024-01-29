//go:build amd64 && !purego
// +build amd64,!purego

package bloom

func queryCore(r *bitrow, bits []bitrow, hashes []uint32)
