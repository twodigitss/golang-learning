package docum
/*
	SOME STUFF I WOULD'HAVE LOVE TO KNOW
	-	it has to have a theotherfileyouwanttotest_test.go in order to ensure a 
		certain file is ran by this tests
	- folder names should be lowercased or at least congruent with the package name
	- 
*/

import (
	"regexp"
	"testing"
)



func TestHelloName(t *testing.T) {
    name := "Gladys"
    msg, err := Hello(name)
    if err != nil {
        t.Errorf("Hello(%q) returned error: %v", name, err)
        return
    }

    want := regexp.MustCompile(`\b` + name + `\b`)
    if !want.MatchString(msg) {
        t.Errorf(`Hello(%q) = %q; want match for %q`, name, msg, want)
    }
}

// TestHelloEmpty calls Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Errorf(`Hello("") = %q, %v; want "", error`, msg, err)
    }
}
