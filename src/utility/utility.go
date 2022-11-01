package utility

import "log"

// CheckErrors checks the errors
func CheckErrors(err error) {
	log.Fatal(err)
}
