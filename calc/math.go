package calc

// Add returns sum of two integer
func Add(numbers ...int) (sum int) {
	sum := 0

	for _, num := range numbers {
		sum += num
	}
	
	return 
}
