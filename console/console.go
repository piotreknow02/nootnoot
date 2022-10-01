package console

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func GetTermSize() (uint16, uint16) {
	width, height, err := terminal.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		panic(fmt.Errorf("Cannot get term size: %v", err))
	}
	return width, height
}

func Veriant1() {

}

func Variant2() {

}

func Variant3() {

}
