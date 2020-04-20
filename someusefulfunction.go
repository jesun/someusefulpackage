package someusefulpackage

import (
	"fmt"

	"gopkg.in/jesun/somearbitrarypackage.v0"
)

// SomeUsefulFunction demonstrates the importing of code from github
func SomeUsefulFunction() {
	// Do something useful...
	fmt.Println("Inside SomeUsefulFunction()...")
	//...by calling...
	somearbitrarypackage.SomeArbitraryFunction()
}
