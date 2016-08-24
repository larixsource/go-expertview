// +build itests

package expertview

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

var (
	testEndpoint = "https://playground.expertview-live.com:443/ExpertWebservice-play/Webservice"
	testUser     = "demo"
	testPassword = "demo"
)

func TestExpertView_GetFileList(t *testing.T) {
	ev := NewExpertView(testEndpoint, DefaultVersion)
	fl, err := ev.GetFileList(Credentials{
		Login:    testUser,
		Password: testPassword,
	})
	require.Nil(t, err)
	// print (don't known the response before hand)
	spew.Dump(fl)
}
