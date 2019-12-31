package calc

// Add returns sum of array elements of type integer
func Add(numbers ...int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	
	return 
}
