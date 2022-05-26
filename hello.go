package hello

import "fmt"

func Hello(name string) {
	fmt.Println(Shello(name))
}

func Shello(name string) string {
	return fmt.Sprint("Hello, ", name)
}
