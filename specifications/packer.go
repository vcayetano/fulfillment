package specifications

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Package struct {
	Size     int `json:"size"`
	Quantity int `json:"quantity"`
}
type Packages []Package

type Packer interface {
	CalculateMinPacks(orderSize int, packageSizes []int) Packages
}

// TotalItems returns the total number of items in the packages.
func (d Packages) TotalItems() int {
	total := 0

	if len(d) == 0 {
		return total
	}

	for _, pack := range d {
		total += pack.Size * pack.Quantity
	}

	return total
}

// PackerSpecification is a test that should be run against any Packer implementation.
func PackerSpecification(t *testing.T, packer Packer) {
	packages := packer.CalculateMinPacks(500000, []int{23, 31, 53})
	assert.Equal(t, packages.TotalItems(), 500000)
}
