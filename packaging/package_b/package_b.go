package package_b

import (
	"github.com/adiludmer/gopractice/packaging/package_a"
)

func B() {
	package_a.A("arg")
}
