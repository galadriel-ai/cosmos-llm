package types

const (
	// ModuleName defines the module name
	ModuleName = "inference"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_inference"

	InferenceKey      = "Inference/value/"
	InferenceCountKey = "Inference/count/"

	GpuNodeKey = "GpuNode/value/"
)

var (
	ParamsKey = []byte("p_inference")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
