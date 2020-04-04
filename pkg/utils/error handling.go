//
//Functions relating to handling errors
//

package utils

import (
	"github.com/sirupsen/logrus"
)

//CritErrorCheck - A oneliner function for error checking, exits the program and logs a message to the terminal if there is an error
func CritErrorCheck(err error, errorMessage string) {
	if err != nil {
		logrus.Warn(errorMessage)
		logrus.Fatal(err)
	}
}
