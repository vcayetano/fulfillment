package specifications

import (
	"math"
	"sort"
)

type DefaultPacker struct {
	//packSizes []int
}

//func NewPacker(availablePack []int) Packer {
//	return &DefaultPacker{
//		packSizes: availablePack,
//	}
//}

// CalculateMinPacks - Calculate the minimum number of packs needed to fulfill the order
func (d *DefaultPacker) CalculateMinPacks(orderSize int, packageSizes []int) Packages {

	if len(packageSizes) == 0 {
		return nil
	}

	// Sort pack sizes in descending order to prioritize larger packs
	sort.Sort(sort.Reverse(sort.IntSlice(packageSizes)))

	minPacks := initializeMinPacks(orderSize)
	packagePacks := make([]int, orderSize+1)

	populateMinPacks(orderSize, packageSizes, minPacks, packagePacks)

	return reconstructResults(orderSize, packageSizes, packagePacks)
}

// initializeMinPacks - initialize the minPacks array with the maximum value
func initializeMinPacks(orderQuantity int) []int {
	minPacks := make([]int, orderQuantity+1)

	for i := 1; i <= orderQuantity; i++ {
		minPacks[i] = math.MaxInt
	}

	// if the order quantity is 0 then we need 0 packs
	minPacks[0] = 1
	return minPacks
}

// populateMinPacks - populate the minPacks array with the minimum number of packs needed
func populateMinPacks(orderQuantity int, packSizes []int, minPacks, packagePacks []int) {
	for _, packSize := range packSizes {

		for currentQuantity := packSize; currentQuantity <= orderQuantity; currentQuantity++ {
			packQuantityIsNotMaxInt := minPacks[currentQuantity-packSize] != math.MaxInt

			if packQuantityIsNotMaxInt && minPacks[currentQuantity-packSize]+1 < minPacks[currentQuantity] {
				minPacks[currentQuantity] = minPacks[currentQuantity-packSize] + 1
				packagePacks[currentQuantity] = packSize
			}
		}
	}
}

// reconstructResults - reconstruct the raw results
func reconstructResults(orderQuantity int, packSizes []int, packagePacks []int) Packages {
	result := make(map[int]int)
	packages := Packages{}

	remainingQuantity := orderQuantity
	for remainingQuantity > 0 {
		packSize := packagePacks[remainingQuantity]
		if packSize == 0 {
			smallestPackSize := packSizes[len(packSizes)-1]
			result[smallestPackSize]++
			remainingQuantity -= smallestPackSize
		} else {
			result[packSize]++
			remainingQuantity -= packSize
		}
	}

	for packSize, quantity := range result {
		packages = append(packages, Package{Size: packSize, Quantity: quantity})
	}

	return packages
}
