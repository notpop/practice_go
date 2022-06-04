// lib is my custom lib.
package lib

// Avarage returns the average of a series of numbers
func Avarage(s []int) int {
	total := 0
	for _, value := range s {
		total += value
	}
	return int(total / len(s))
}
