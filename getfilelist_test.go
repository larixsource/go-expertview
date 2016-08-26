package expertview

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpertView_GetFileList(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		ok := bytes.Contains(rb, []byte("<sq:getFileList><login>demo</login><password>fe01ce2a7fbac8fafaed7c982a04e229</password><version>2.5.0</version></sq:getFileList>"))
		if !ok {
			t.Errorf("invalid soap call: %s", rb)
			t.FailNow()
		}

		f, err := os.Open("testdata/getFileListResponse.xml")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		rw.Write(b)
	}))
	defer server.Close()

	ev := NewExpertView(server.URL, DefaultVersion)
	fl, err := ev.GetFileList(Credentials{
		Login:    "demo",
		Password: "demo",
	})
	require.Nil(t, err)

	assert.Len(t, fl.DeviceTypes, 3)
	assert.Equal(t, "8000-1", fl.DeviceTypes[0].ProductNumber)
	assert.Equal(t, "Normal (Solid/Flex)", fl.DeviceTypes[0].Description)
	assert.Len(t, fl.Records, 56)
	assert.Equal(t, DCFKind, fl.Records[0].Kind)
	assert.Equal(t, "INTA1-FLX12-MBA+VDO-RS-120224CL.DCF", fl.Records[0].Name)
	assert.Equal(t, "D9984527582012022715184747.dcf", fl.Records[0].File)
}

func TestExpertView_GetFileListAuthEx(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		ok := bytes.Contains(rb, []byte("<sq:getFileList><login>demo</login><password>fe01ce2a7fbac8fafaed7c982a04e229</password><version>2.5.0</version></sq:getFileList>"))
		if !ok {
			t.Errorf("invalid soap call: %s", rb)
			t.FailNow()
		}

		f, err := os.Open("testdata/getFileListResponseAuthEx.xml")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		rw.Write(b)
	}))
	defer server.Close()

	ev := NewExpertView(server.URL, DefaultVersion)
	_, err := ev.GetFileList(Credentials{
		Login:    "demo",
		Password: "demo",
	})
	assert.Equal(t, ErrAuthentication, err)
}
