package expertview

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
)

type deviceTypeXml struct {
	Description   string `xml:"DESCRIPTION,attr"`
	ProductNumber string `xml:"PRODUCTNUMBER,attr"`
}

type recordXml struct {
	Kind string `xml:"KIND,attr"`
	Name string `xml:"NAME"`
	File string `xml:"FILE"`
}

type getFileListResponseXml struct {
	XMLName     xml.Name        `xml:"RESPONSE"`
	DeviceTypes []deviceTypeXml `xml:"DEVICETYPES>DEVICETYPE"`
	Files       []recordXml     `xml:"FILES>RECORD"`
}

func parseGetFileList(r []byte) (FileList, error) {
	env := &soapEnvelope{}
	err := xml.Unmarshal(r, env)
	if err != nil {
		return FileList{}, err
	}

	fault := env.Body.Fault
	if fault != nil {
		detail := fault.Detail
		switch {
		case detail != nil && detail.AuthenticationException != nil:
			return FileList{}, ErrAuthentication
		case detail != nil && detail.UnexpectedException != nil:
			return FileList{}, ErrUnexpected
		default:
			return FileList{}, errors.New("unknown error")
		}
	}

	soapGetFileList := env.Body.GetFileListResponse
	switch {
	case soapGetFileList == nil:
		return FileList{}, errors.New("getFileListResponse not found")
	case len(soapGetFileList.Return) == 0:
		return FileList{}, errors.New("empty getFileListResponse")
	}

	// decode return (base64)
	dl := base64.StdEncoding.DecodedLen(len(soapGetFileList.Return))
	buf := make([]byte, dl)
	_, decErr := base64.StdEncoding.Decode(buf, soapGetFileList.Return)
	if decErr != nil {
		return FileList{}, fmt.Errorf("error decoding getFileListResponse b64 data: %s", decErr)
	}

	// decode xml
	var gflr getFileListResponseXml
	decErr = xml.NewDecoder(bytes.NewReader(buf)).Decode(&gflr)
	if decErr != nil {
		return FileList{}, fmt.Errorf("error decoding getFileListResponse xml data: %s", decErr)
	}

	// to FileList
	deviceTypes := make([]DeviceType, 0, len(gflr.DeviceTypes))
	for _, dt := range gflr.DeviceTypes {
		deviceTypes = append(deviceTypes, DeviceType{
			Description:   dt.Description,
			ProductNumber: dt.ProductNumber,
		})
	}
	records := make([]Record, 0, len(gflr.Files))
	for _, rec := range gflr.Files {
		records = append(records, Record{
			Kind: RecordKind(rec.Kind),
			Name: rec.Name,
			File: rec.File,
		})
	}
	return FileList{
		DeviceTypes: deviceTypes,
		Records:     records,
	}, nil
}
