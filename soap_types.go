package expertview

import "encoding/xml"

type soapEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body soapBody
}

type soapBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault               *soapFault           `xml:",omitempty"`
	GetFileListResponse *getFileListResponse `xml:"getFileListResponse,omitempty"`
}

type soapFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string      `xml:"faultcode,omitempty"`
	String string      `xml:"faultstring,omitempty"`
	Actor  string      `xml:"faultactor,omitempty"`
	Detail *soapDetail `xml:"detail,omitempty"`
}

type soapDetail struct {
	AuthenticationException *authenticationException `xml:"AuthenticationException,omitempty"`
}

type authenticationException struct {
	XMLName xml.Name `xml:"http://webservice.expertview.squarell.com/ AuthenticationException"`

	Message string `xml:"message,omitempty"`
}

type getFileListResponse struct {
	XMLName xml.Name `xml:"http://webservice.expertview.squarell.com/ getFileListResponse"`

	Return []byte `xml:"return"`
}
