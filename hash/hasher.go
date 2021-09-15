package hash

type Hasher interface {
	Hash64(data string) (uint64, error)
}
