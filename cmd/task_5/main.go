// Что выведет программа?
//
// Объяснить вывод программы.
////////////////////////////////

package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error

	err = test()
	// test() возвращает *customError, но err имеет тип err
	// Go приводит *customError к интерфейсу error
	// Интерфейс это пара (type, value), или в этом случае: err = (type=*customError, value=nil)
	// Интерфейс НЕ равен nil, если у него есть тип, даже если значение nil.

	if err != nil {
		println("error") // отработает эта строка
		return
	}
	println("ok") // это не отработает
}
