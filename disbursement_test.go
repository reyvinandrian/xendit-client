package xenditgo_test

import (
	"testing"

	"github.com/cheekybits/is"
	xenditgo "github.com/reyvinandrian/xendit-client"
)

func TestBatchDisbEnvironmentType(t *testing.T) {
	is := is.New(t)

	client := xenditgo.NewClient()
	is.Equal(xenditgo.Sandbox, client.APIEnvType)

}
