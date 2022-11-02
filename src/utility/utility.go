package utility

import "log"

// CheckErrors checks the errors
func CheckErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
