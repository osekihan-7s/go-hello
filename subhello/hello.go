package subhello

import "fmt"

func Hello(name string) {
	fmt.Println(Shello(name))
}

func Shello(name string) string {
	return fmt.Sprint("Hello, ", name, " from subhello")
}

func Shello2(name string) string {
	return fmt.Sprint("Hello, ", name, " from subhello2")
}

func Shello3(name string) string {
	return fmt.Sprint("Hello, ", name, " from subhello3")
}
