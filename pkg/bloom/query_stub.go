//go:build amd64 && !purego
// +build amd64,!purego

package bloom

//go:generate go run asm.go -out query_amd64.s
//go:noescape
func queryCore(r *bitrow, bits []bitrow, hashes []uint32)
