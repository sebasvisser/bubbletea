package main

import (
	"fmt"
)

// getPairs returns all pairs (a, b) where 2 <= a <= b <= 99 and a + b = sum
// Changed constraint: both numbers must be > 1 (classic version)
func getPairs(sum int) [][2]int {
	var pairs [][2]int
	for a := 2; a <= sum/2; a++ {
		b := sum - a
		if b >= 2 && b <= 99 && a <= b {
			pairs = append(pairs, [2]int{a, b})
		}
	}
	return pairs
}

// getFactorizations returns all pairs (a, b) where a * b = product and 2 <= a <= b <= 99
// Changed constraint: both numbers must be > 1 (classic version)
func getFactorizations(product int) [][2]int {
	var pairs [][2]int
	for a := 2; a*a <= product; a++ {
		if product%a == 0 {
			b := product / a
			if b >= 2 && b <= 99 && a >= 2 && a <= 99 && a <= b {
				pairs = append(pairs, [2]int{a, b})
			}
		}
	}
	return pairs
}

// isProductAmbiguous checks if a product can be formed by multiple pairs
func isProductAmbiguous(product int) bool {
	factorizations := getFactorizations(product)
	return len(factorizations) > 1
}

// Step 1: Find all sums where X can say "I know you don't know"
// This means: for ALL pairs (a,b) that sum to S, their product is ambiguous
func findValidSumsForX() map[int]bool {
	validSums := make(map[int]bool)

	// Check all possible sums
	for sum := 2; sum <= 198; sum++ { // max sum is 99 + 99 = 198
		pairs := getPairs(sum)
		if len(pairs) == 0 {
			continue
		}

		// Check if ALL products from this sum are ambiguous
		allAmbiguous := true
		for _, pair := range pairs {
			product := pair[0] * pair[1]
			if !isProductAmbiguous(product) {
				allAmbiguous = false
				break
			}
		}

		if allAmbiguous {
			validSums[sum] = true
		}
	}

	return validSums
}

// Step 2: For each product, check if exactly one valid sum is consistent with it
func solve() {
	fmt.Println("Solving the riddle...")
	fmt.Println()

	// Find all sums where X could say "I know you don't know"
	validSums := findValidSumsForX()

	fmt.Printf("Found %d valid sums where X can say 'I know you don't know'\n", len(validSums))

	// Display all valid sums sorted
	allValidSums := []int{}
	for sum := range validSums {
		allValidSums = append(allValidSums, sum)
	}
	// Simple bubble sort
	for i := 0; i < len(allValidSums); i++ {
		for j := i + 1; j < len(allValidSums); j++ {
			if allValidSums[i] > allValidSums[j] {
				allValidSums[i], allValidSums[j] = allValidSums[j], allValidSums[i]
			}
		}
	}
	fmt.Println("All valid sums:", allValidSums)
	fmt.Println()

	// For each possible product, check if Y can uniquely determine A and B
	candidateSolutions := make(map[int][][2]int)

	for product := 1; product <= 99*99; product++ {
		factorizations := getFactorizations(product)

		// Y must have multiple factorizations initially (otherwise Y would know immediately)
		if len(factorizations) <= 1 {
			continue
		}

		// After X's statement, Y knows the sum must be in validSums
		// Check which factorizations have sums in validSums
		validFactorizations := [][2]int{}
		for _, pair := range factorizations {
			sum := pair[0] + pair[1]
			if validSums[sum] {
				validFactorizations = append(validFactorizations, pair)
			}
		}

		// Y can determine A and B if exactly one factorization has a valid sum
		if len(validFactorizations) == 1 {
			candidateSolutions[product] = validFactorizations
		}
	}

	fmt.Printf("Found %d products where Y can determine A and B after X's statement\n", len(candidateSolutions))
	fmt.Println()

	// Step 3: After Y says "Now I know", X should also be able to determine A and B
	// For each valid sum, check if exactly one of its pairs leads to a product in candidateSolutions
	finalSolutions := [][2]int{}

	for sum := range validSums {
		pairs := getPairs(sum)
		validPairs := [][2]int{}

		for _, pair := range pairs {
			product := pair[0] * pair[1]
			if _, ok := candidateSolutions[product]; ok {
				validPairs = append(validPairs, pair)
			}
		}

		// X can determine A and B if exactly one pair from this sum is in candidateSolutions
		if len(validPairs) == 1 {
			finalSolutions = append(finalSolutions, validPairs[0])
		}
	}

	// Display final solutions
	if len(finalSolutions) > 0 {
		fmt.Println("Final Solution(s) where both X and Y can determine A and B:")
		for _, pair := range finalSolutions {
			sum := pair[0] + pair[1]
			product := pair[0] * pair[1]
			fmt.Printf("\nA = %d, B = %d\n", pair[0], pair[1])
			fmt.Printf("Sum (known by X) = %d\n", sum)
			fmt.Printf("Product (known by Y) = %d\n", product)
			fmt.Println("\nVerification:")
			fmt.Printf("  1. Y sees product %d with factorizations: %v\n", product, getFactorizations(product))
			fmt.Printf("     Y says 'I don't know' (multiple factorizations exist)\n")
			fmt.Printf("  2. X sees sum %d, all pairs sum to %d: %v\n", sum, sum, getPairs(sum))
			fmt.Printf("     X verifies all these products are ambiguous, so X says 'I know you don't know'\n")
			fmt.Printf("  3. Y realizes only the pair (%d, %d) has sum %d which is valid for X's statement\n", pair[0], pair[1], sum)
			fmt.Printf("     Y says 'Now I know': A = %d, B = %d\n", pair[0], pair[1])
			fmt.Printf("  4. X can now determine which pair from sum %d is the answer\n", sum)
		}
	} else {
		fmt.Println("No unique solution found!")
	}
}

func getSampleSums(sums map[int]bool, limit int) []int {
	sample := []int{}
	for sum := range sums {
		sample = append(sample, sum)
		if len(sample) >= limit {
			break
		}
	}
	return sample
}

func main() {
	solve()
}
