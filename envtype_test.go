package xenditgo_test

import (
	"testing"

	"github.com/cheekybits/is"
	xenditgo "github.com/reyvinandrian/xendit-client"
)

func TestEnvironmentType(t *testing.T) {
	is := is.New(t)
	is.Equal("https://api.xendit.co", xenditgo.Sandbox.String())
	is.Equal("https://api.xendit.co", xenditgo.Production.String())
}
