package specifications

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPack(t *testing.T) {

	t.Run("It returns the correct number of items for the edge cases", func(t *testing.T) {
		testPacks := []int{23, 31, 53}
		dp := &DefaultPacker{}

		tests := []struct {
			name               string
			orderCount         int
			expectedItemsCount int
		}{
			{
				name:               "returns correct pack count when order count is 500000 (edge case)",
				orderCount:         500000,
				expectedItemsCount: 500000,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {

				got := dp.CalculateMinPacks(tt.orderCount, testPacks)

				if !assert.Equal(t, tt.expectedItemsCount, got.TotalItems()) {
					t.Errorf("Expected %v but got %v", tt.expectedItemsCount, got)
				}
			})
		}
	})

	t.Run("it returns the correct number of items to use for the order count", func(t *testing.T) {

		testPacks := []int{250, 500, 1000, 2000, 5000}

		dp := DefaultPacker{}

		tests := []struct {
			name               string
			orderCount         int
			expectedItemsCount int
		}{
			{
				name:               "returns correct pack count when order count is 0",
				orderCount:         0,
				expectedItemsCount: 0,
			},

			{
				name:               "returns correct pack count when order count is 1",
				orderCount:         1,
				expectedItemsCount: 250,
			},
			{
				name:               "returns correct pack count when order count is 250",
				orderCount:         250,
				expectedItemsCount: 250,
			},
			{
				name:               "returns correct pack count when order count is 251",
				orderCount:         251,
				expectedItemsCount: 500,
			},

			{
				name:               "returns correct pack count when order count is 500",
				orderCount:         500,
				expectedItemsCount: 500,
			},

			{
				name:               "returns correct pack count when order count is 501",
				orderCount:         501,
				expectedItemsCount: 750,
			},

			{
				name:               "returns correct pack count when order count is 700",
				orderCount:         700,
				expectedItemsCount: 750,
			},

			{
				name:               "returns correct pack count when order count is 1000",
				orderCount:         1000,
				expectedItemsCount: 1000,
			},

			{
				name:               "returns correct pack count when order count is 1001",
				orderCount:         1001,
				expectedItemsCount: 1250,
			},

			{
				name:               "returns correct pack count when order count is 1500",
				orderCount:         1500,
				expectedItemsCount: 1500,
			},

			{
				name:               "returns correct pack count when order count is 2000",
				orderCount:         2000,
				expectedItemsCount: 2000,
			},

			{
				name:               "returns correct pack count when order count is 4000",
				orderCount:         4000,
				expectedItemsCount: 4000,
			},

			{
				name:               "returns correct pack count when order count is 12500",
				orderCount:         12500,
				expectedItemsCount: 12500,
			},

			{
				name:               "returns correct pack count when order count is 12001",
				orderCount:         12001,
				expectedItemsCount: 12250,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {

				got := dp.CalculateMinPacks(tt.orderCount, testPacks)

				if !assert.Equal(t, tt.expectedItemsCount, got.TotalItems()) {
					t.Errorf("Expected %v items but got %v items", tt.expectedItemsCount, got)
				}
			})
		}

	})
}
