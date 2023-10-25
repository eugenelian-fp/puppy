package puppy

import (
	"fmt"

	"github.com/eugenelian-fp/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof Woof Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func V2() {
	fmt.Println("v1.2.0")
}

func V3() {
	fmt.Println("v1.2.1")
}
