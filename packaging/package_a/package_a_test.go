package package_a

import "testing"

func TestA(t *testing.T) {
	t.Run("TestA", func(t *testing.T) {
		A()
	})
}
