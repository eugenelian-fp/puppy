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

func v2() string {
	fmt.Println("v1.2.0")
}
