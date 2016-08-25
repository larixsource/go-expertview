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

func TestExpertView_GetFile(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		ok := bytes.Contains(rb, []byte("<sq:getFile><login>demo</login><password>fe01ce2a7fbac8fafaed7c982a04e229</password><version>2.5.0</version><filename>D9984527582012022715184747.dcf</filename></sq:getFile>"))
		if !ok {
			t.Errorf("invalid soap call: %s", rb)
			t.FailNow()
		}

		f, err := os.Open("testdata/getFileResponse.xml")
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
	f, err := ev.GetFile(Credentials{
		Login:    "demo",
		Password: "demo",
	}, "D9984527582012022715184747.dcf")
	require.Nil(t, err)

	assert.Equal(t, []byte("hello"), f)
}

func TestExpertView_GetFileAuthEx(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		ok := bytes.Contains(rb, []byte("<sq:getFile><login>demo</login><password>fe01ce2a7fbac8fafaed7c982a04e229</password><version>2.5.0</version><filename>D9984527582012022715184747.dcf</filename></sq:getFile>"))
		if !ok {
			t.Errorf("invalid soap call: %s", rb)
			t.FailNow()
		}

		f, err := os.Open("testdata/getFileResponseAuthEx.xml")
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
	_, err := ev.GetFile(Credentials{
		Login:    "demo",
		Password: "demo",
	}, "D9984527582012022715184747.dcf")
	assert.Equal(t, ErrAuthentication, err)
}

func TestExpertView_GetFileNotSelectableEx(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		ok := bytes.Contains(rb, []byte("<sq:getFile><login>demo</login><password>fe01ce2a7fbac8fafaed7c982a04e229</password><version>2.5.0</version><filename>D9984527582012022715184747.dcf</filename></sq:getFile>"))
		if !ok {
			t.Errorf("invalid soap call: %s", rb)
			t.FailNow()
		}

		f, err := os.Open("testdata/getFileResponseFwNotSelectableEx.xml")
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
	_, err := ev.GetFile(Credentials{
		Login:    "demo",
		Password: "demo",
	}, "D9984527582012022715184747.dcf")
	assert.Equal(t, ErrFirmwareNotSelectable, err)
}
