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

func TestExpertView_GetInstallationRecords(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		ok := bytes.Contains(rb, []byte("<sq:getInstallRecords><login>demo</login><password>fe01ce2a7fbac8fafaed7c982a04e229</password><version>2.5.0</version></sq:getInstallRecords>"))
		if !ok {
			t.Errorf("invalid soap call: %s", rb)
			t.FailNow()
		}

		f, err := os.Open("testdata/getInstallRecordsResponse.xml")
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

	ev, err := NewExpertView(server.URL, DefaultVersion, Credentials{
		Login:    "demo",
		Password: "demo",
	})
	require.Nil(t, err)
	fl, err := ev.GetInstallationRecords()
	require.Nil(t, err)

	assert.Len(t, fl, 2)
	assert.Equal(t, "296930501", fl[0].SerialNumber)
	assert.Equal(t, "667769", fl[0].ID)
	assert.Equal(t, "Larix Ltda", fl[0].Telematic)
	assert.Equal(t, "FLX12", fl[0].HardwareProf)
	assert.Equal(t, "FLEX-256-V2", fl[0].SoftwareProf)
	assert.Equal(t, "SQU-8000-TRKS-000000-131001CL.DCF", fl[0].DCF)
	assert.Equal(t, "8000-01V114R048.BIN", fl[0].Firmware)
	assert.Equal(t, "293230583 FLX12 FLEX-256-V2 hjashj dshj dahy", fl[0].Key)
	assert.Equal(t, "asdf", fl[0].Username)
}

func TestExpertView_GetInstallationRecordsAuthEx(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		ok := bytes.Contains(rb, []byte("<sq:getInstallRecords><login>demo</login><password>fe01ce2a7fbac8fafaed7c982a04e229</password><version>2.5.0</version></sq:getInstallRecords>"))
		if !ok {
			t.Errorf("invalid soap call: %s", rb)
			t.FailNow()
		}

		f, err := os.Open("testdata/getInstallRecordsResponseAuthEx.xml")
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

	ev, err := NewExpertView(server.URL, DefaultVersion, Credentials{
		Login:    "demo",
		Password: "demo",
	})
	require.Nil(t, err)
	_, err = ev.GetInstallationRecords()
	assert.Equal(t, ErrAuthentication, err)
}
