package crypto

import (
	"errors"
	"fmt"

	"github.com/onflow/flow-go/crypto/hash"
	"github.com/onflow/flow-go/model/flow"
)

// prefixedHashing, embeds a crypto hasher
type prefixedHashing struct {
	hash.Hasher
	tag [flow.DomainTagLength]byte
}

// paddedDomainTag converts a string into a padded byte array
func paddedDomainTag(s string) ([flow.DomainTagLength]byte, error) {
	var tag [flow.DomainTagLength]byte
	if len(s) > flow.DomainTagLength {
		return tag, fmt.Errorf("domain tag %s cannot be longer than %d characters", s, flow.DomainTagLength)
	}
	copy(tag[:], s)
	return tag, nil
}

// NewPrefixedHashing returns a new hasher that prefixes the tag for all
// hash computations.
// Only SHA2 and SHA3 algorithms are supported.
func NewPrefixedHashing(shaAlgo hash.HashingAlgorithm, tag string) (hash.Hasher, error) {

	var hasher hash.Hasher
	switch shaAlgo {
	case hash.SHA2_256:
		hasher = hash.NewSHA2_256()
	case hash.SHA3_256:
		hasher = hash.NewSHA3_256()
	case hash.SHA2_384:
		hasher = hash.NewSHA2_384()
	case hash.SHA3_384:
		hasher = hash.NewSHA3_384()
	default:
		return nil, errors.New("hashing algorithm is not a supported for prefixed algorithm")
	}

	paddedTag, err := paddedDomainTag(tag)
	if err != nil {
		return nil, fmt.Errorf("prefixed hashing failed: %w", err)
	}

	return &prefixedHashing{
		Hasher: hasher,
		tag:    paddedTag,
	}, nil
}

// ComputeHash calculates and returns the digest of input byte array prefixed by a tag.
// Not thread-safe
func (s *prefixedHashing) ComputeHash(data []byte) hash.Hash {
	s.Reset()
	_, _ = s.Write(data)
	return s.Hasher.SumHash()
}

// Reset gets the hasher back to its original state.
func (s *prefixedHashing) Reset() {
	s.Hasher.Reset()
	_, _ = s.Write(s.tag[:])
}
