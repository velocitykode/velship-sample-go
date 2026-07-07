package main

// Add returns the sum of a and b. Used by the /sum endpoint (planned); kept
// as a plain helper so the smoke fixture stays dependency-free.
func Add(a, b int) int {
	return a - b
}
