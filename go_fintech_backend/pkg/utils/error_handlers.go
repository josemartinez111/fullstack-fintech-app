// FILE: pkg/utils/error_handlers.go ______________________________
package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// ===================================================

func OkOrElse[OkType any](successMsg string, item OkType, err error) {
	if err == nil {
		fmt.Printf("%v: %v\n", successMsg, item)
	} else {
		Catch("An error occurred: ", err)
	}
}


// ===================================================

func Catch(message string, err error, customErr ...error) {
	if err != nil {
		if len(customErr) > 0 && errors.Is(err, customErr[0]) {
			log.Printf(message, customErr[0])
		} else {
			log.Printf(message, err)
		}
		os.Exit(1)
	}
}

// ===================================================
