package package_a

import "fmt"

var (
	a string
)

type AA struct {
	b string
}

func A(arg any) {
	j := AA{b: "local b"}
	b := "local b"
	a = "Hello from package_a"
	fmt.Println("A", a, b, arg, j)
}
