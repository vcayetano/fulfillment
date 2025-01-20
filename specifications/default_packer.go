package specifications

import (
	"math"
	"sort"
)

type DefaultPacker struct {
	//packSizes []int
}

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

	// initialize the minPacks array with the maximum value - MaxInt32 or MaxInt64 depending on intSize.
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

		// iterate through the pack sizes and calculate the minimum number of packs needed
		for currentQuantity := packSize; currentQuantity <= orderQuantity; currentQuantity++ {
			// check if the current quantity to current pack difference is less than the max int
			quantityPackSizeDifference := currentQuantity - packSize
			packQuantityIsNotMaxInt := minPacks[quantityPackSizeDifference] != math.MaxInt

			// if the current quantity is less than the max int and the current quantity is less than the current min pack
			if packQuantityIsNotMaxInt && minPacks[quantityPackSizeDifference]+1 < minPacks[currentQuantity] {
				minPacks[currentQuantity] = minPacks[quantityPackSizeDifference] + 1
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
	// iterate through the packagePacks array and reconstruct the results
	for remainingQuantity > 0 {
		packSize := packagePacks[remainingQuantity]
		//
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
