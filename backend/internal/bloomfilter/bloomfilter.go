package bloomfilter

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/VarthanV/bloom-filter-usecase/backend/internal/intutils"
)

type bloomFilter struct {
	capacity int
	boolArr  []bool
}

func New(capacity int) IBloomFilter {
	return &bloomFilter{
		capacity: capacity,
		boolArr:  make([]bool, capacity),
	}
}

func (b *bloomFilter) hash(input string) (int, error) {

	h := sha256.New()
	_, err := h.Write([]byte(input))
	if err != nil {
		fmt.Println("error in hashing input ", err)
		return -1, err
	}

	sumBytes := h.Sum(nil)
	dataInt := binary.BigEndian.Uint64(sumBytes)
	return intutils.AbsInt(int(dataInt) % b.capacity), err
}

// Add implements IBloomFilter
func (b *bloomFilter) Add(key string) error {
	idx, err := b.hash(key)
	if err != nil {
		fmt.Println("error in adding key to filter ", err)
		return err
	}
	b.boolArr[idx] = true
	return nil
}

// CheckMembership implements IBloomFilter
func (b *bloomFilter) CheckMembership(key string) (bool, error) {
	idx, err := b.hash(key)
	if err != nil {
		fmt.Println("error in adding key to filter ", err)
		return false, err
	}
	return b.boolArr[idx], nil
}
