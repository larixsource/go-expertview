package expertview

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const defaultTimeout = 10 * time.Second

type soapCli struct {
	Timeout            time.Duration
	Endpoint           string
	InsecureSkipVerify bool
}

func (sc *soapCli) call(r io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", sc.Endpoint, r)
	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Set("User-Agent", "go-expertview/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: sc.InsecureSkipVerify,
		},
		Dial: func(network, addr string) (net.Conn, error) {
			timeout := sc.Timeout
			if timeout == 0 {
				timeout = defaultTimeout
			}
			return net.DialTimeout(network, addr, timeout)
		},
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
