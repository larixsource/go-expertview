package expertview

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
)

type installRecordsXml struct {
	XMLName xml.Name           `xml:"INSTALLATIONS"`
	Records []installRecordXml `xml:"RECORD"`
}

type installRecordXml struct {
	SN           string `xml:"SN,attr"`
	ID           string `xml:"ID"`
	Telematic    string `xml:"TELEMATIC"`
	HardwareProf string `xml:"HARDWAREPROF"`
	SoftwareProf string `xml:"SOFTWAREPROF"`
	DCF          string `xml:"DCF"`
	Firmware     string `xml:"FIRMWARE"`
	Key          string `xml:"KEY"`
	Username     string `xml:"USERNAME"`
}

func parseGetInstallRecords(r []byte) ([]InstallationRecord, error) {
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
		default:
			return nil, errors.New("unknown error")
		}
	}

	soapGetInstallRecords := env.Body.GetInstallRecordsResponse
	switch {
	case soapGetInstallRecords == nil:
		return nil, errors.New("getInstallRecordsResponse not found")
	case len(soapGetInstallRecords.Return) == 0:
		return nil, errors.New("empty getInstallRecordsResponse")
	}

	// decode return (base64)
	dl := base64.StdEncoding.DecodedLen(len(soapGetInstallRecords.Return))
	buf := make([]byte, dl)
	n, decErr := base64.StdEncoding.Decode(buf, soapGetInstallRecords.Return)
	if decErr != nil {
		return nil, fmt.Errorf("error decoding getInstallRecordsResponse b64 data: %s", decErr)
	}

	// decode xml
	var xmlRecords installRecordsXml
	decErr = xml.NewDecoder(bytes.NewReader(buf[:n])).Decode(&xmlRecords)
	if decErr != nil {
		return nil, fmt.Errorf("error decoding getInstallRecordsResponse xml data: %s", decErr)
	}

	// to []InstallationRecord
	instRecords := make([]InstallationRecord, 0, len(xmlRecords.Records))
	for _, rec := range xmlRecords.Records {
		instRecords = append(instRecords, InstallationRecord{
			SerialNumber: rec.SN,
			ID:           rec.ID,
			Telematic:    rec.Telematic,
			HardwareProf: rec.HardwareProf,
			SoftwareProf: rec.SoftwareProf,
			DCF:          rec.DCF,
			Firmware:     rec.Firmware,
			Key:          rec.Key,
			Username:     rec.Username,
		})
	}
	return instRecords, nil
}
