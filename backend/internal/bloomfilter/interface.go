package bloomfilter

type IBloomFilter interface {
	Add(key string) error
	CheckMembership(key string) (bool, error)
}
