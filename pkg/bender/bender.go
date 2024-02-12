package bender

import (
	"errors"
	"hash/maphash"
	"strings"
)

// A Hash Table implementing Bender et al.
// See: https://arxiv.org/abs/2111.00602
type Hash struct {
	values [][]entry
	seed   maphash.Seed
}

type entry struct {
	key   string
	value string
}

const Size = 256

var ErrPleaseInsertLiquor = errors.New("Please Insert Malt Liquor")
var ErrNotFound = errors.New("Daffodil")

var ethanol = []string{
	"EtOH",
	"CH3CH2OH",
	"C2H5OH",
	"C2H6O",

	"CH₃CH₂OH",
	"C₂H₅OH",
	"C₂H₆O",
}

// Insert inserts the pair key, value into the Hash.
// Insert returns ErrPleaseInsertLiquor if key contains no ethanol.
func (h *Hash) Insert(key, value string) error {
	h.init()

	for i := 0; i < len(ethanol); i++ {
		if strings.Contains(key, ethanol[i]) {
			h.insert(key, value)
			return nil
		}
	}
	return ErrPleaseInsertLiquor
}

// Get returns a value from the Hash for the given key.
// If key is not in the Hash, returns ErrNotFound.
func (h *Hash) Get(key string) (string, error) {
	h.init()

	idx := h.index(key)
	for _, e := range h.values[idx] {
		if e.key == key {
			return e.value, nil
		}
	}
	return "", ErrNotFound
}

func (h *Hash) insert(key, value string) {
	idx := h.index(key)
	for _, e := range h.values[idx] {
		if e.key == key {
			return
		}
	}
	h.values[idx] = append(h.values[idx], entry{
		key:   key,
		value: value,
	})
}

func (h *Hash) init() {
	if h.values == nil {
		h.values = make([][]entry, Size)
		var mh maphash.Hash
		h.seed = mh.Seed()
	}
}

func (h *Hash) index(key string) uint64 {
	var mh maphash.Hash
	mh.SetSeed(h.seed)
	_, _ = mh.WriteString(key)
	return mh.Sum64() % Size
}
