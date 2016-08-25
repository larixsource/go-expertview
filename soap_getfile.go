package expertview

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
)

func parseGetFile(r []byte) ([]byte, error) {
	env := &soapEnvelope{}
	err := xml.Unmarshal(r, env)
	if err != nil {
		return nil, err
	}

	fault := env.Body.Fault
	if fault != nil {
		detail := fault.Detail
		switch {
		case detail != nil && detail.AuthenticationException != nil:
			return nil, ErrAuthentication
		case detail != nil && detail.UnexpectedException != nil:
			return nil, ErrUnexpected
		case detail != nil && detail.FirmwareNotSelectableException != nil:
			return nil, ErrFirmwareNotSelectable
		case detail != nil && detail.NoSuchFileException != nil:
			return nil, ErrAuthentication
		default:
			return nil, errors.New("unknown error")
		}
	}

	soapGetFile := env.Body.GetFileResponse
	switch {
	case soapGetFile == nil:
		return nil, errors.New("getFileResponse not found")
	case len(soapGetFile.Return) == 0:
		return nil, errors.New("empty getFileResponse")
	}

	// decode return (base64)
	dl := base64.StdEncoding.DecodedLen(len(soapGetFile.Return))
	buf := make([]byte, dl)
	n, decErr := base64.StdEncoding.Decode(buf, soapGetFile.Return)
	if decErr != nil {
		return nil, fmt.Errorf("error decoding getFileResponse b64 data: %s", decErr)
	}
	return buf[:n], nil
}
