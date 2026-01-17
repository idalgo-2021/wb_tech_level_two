// Что выведет программа?
//
// Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.
///////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err) // nil - err это интерфейс(type = *os.PathError, value = nil), т.е. реализует error.
	// Ф-я печати в вызывает Error() только если value != nil, но value = nil, поэтому -> <nil>

	fmt.Println(err == nil) // false - Интерфейс считается nil только если (type == nil && value == nil),
	// но type = *os.PathError. Таким образом -> false
}
