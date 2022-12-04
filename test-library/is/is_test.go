package is

import (
	"github.com/matryer/is"
	"strings"
	"testing"
)

func Binary(b string) (bool, error) {
	return true, nil
}
func TestSomething(t *testing.T) {
	is := is.New(t)

	b, err := Binary("0")

	is.NoErr(err)
	is.Equal(b, true)
	is.Equal([]string{"a", "b"}, []string{"a", "b"})
	got := "anuchito is gopher"
	is.True(strings.Contains(got, "anuchito"))
}
