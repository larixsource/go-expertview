package expertview

import "encoding/xml"

type soapEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body soapBody
}

type soapBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault                     *soapFault                 `xml:",omitempty"`
	GetFileListResponse       *getFileListResponse       `xml:"getFileListResponse,omitempty"`
	GetFileResponse           *getFileResponse           `xml:"getFileResponse,omitempty"`
	GetInstallRecordsResponse *getInstallRecordsResponse `xml:"getInstallRecordsResponse,omitempty"`
}

type soapFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string      `xml:"faultcode,omitempty"`
	String string      `xml:"faultstring,omitempty"`
	Actor  string      `xml:"faultactor,omitempty"`
	Detail *soapDetail `xml:"detail,omitempty"`
}

type soapDetail struct {
	AuthenticationException        *authenticationException        `xml:"AuthenticationException,omitempty"`
	UnexpectedException            *unexpectedException            `xml:"UnexpectedException,omitempty"`
	FirmwareNotSelectableException *firmwareNotSelectableException `xml:"FirmwareNotSelectableException,omitempty"`
	NoSuchFileException            *noSuchFileException            `xml:"NoSuchFileException,omitempty"`
}

type authenticationException struct {
	XMLName xml.Name `xml:"http://webservice.expertview.squarell.com/ AuthenticationException"`

	Message string `xml:"message,omitempty"`
}

type unexpectedException struct {
	XMLName xml.Name `xml:"http://webservice.expertview.squarell.com/ UnexpectedException"`

	Message string `xml:"message,omitempty"`
}

type getFileListResponse struct {
	XMLName xml.Name `xml:"http://webservice.expertview.squarell.com/ getFileListResponse"`

	Return []byte `xml:"return"`
}

type firmwareNotSelectableException struct {
	XMLName xml.Name `xml:"http://webservice.expertview.squarell.com/ FirmwareNotSelectableException"`

	Message string `xml:"message,omitempty"`
}

type noSuchFileException struct {
	XMLName xml.Name `xml:"http://webservice.expertview.squarell.com/ NoSuchFileException"`

	Message string `xml:"message,omitempty"`
}

type getFileResponse struct {
	XMLName xml.Name `xml:"http://webservice.expertview.squarell.com/ getFileResponse"`

	Return []byte `xml:"return"`
}

type getInstallRecordsResponse struct {
	XMLName xml.Name `xml:"http://webservice.expertview.squarell.com/ getInstallRecordsResponse"`

	Return []byte `xml:"return"`
}
