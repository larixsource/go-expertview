package expertview

import (
	"github.com/lestrrat/go-libxml2/dom"
	"github.com/lestrrat/go-libxml2/types"
)

func createGetFileList(cred Credentials, version string) (*dom.Document, error) {
	doc, _, body, err := createEnvelope()
	if err != nil {
		return doc, err
	}
	node, err := createBaseNode(doc, "getFileList", cred, version)
	if err != nil {
		return doc, err
	}
	body.AddChild(node)
	return doc, nil
}

func createEnvelope() (doc *dom.Document, header types.Element, body types.Element, err error) {
	doc = dom.CreateDocument()

	envelope, err := doc.CreateElementNS("http://schemas.xmlsoap.org/soap/envelope/", "S:Envelope")
	if err != nil {
		return
	}
	err = envelope.SetAttribute("xmlns:sq", "http://webservice.expertview.squarell.com/")
	if err != nil {
		return
	}

	header, err = doc.CreateElement("S:Header")
	if err != nil {
		return
	}
	err = envelope.AddChild(header)
	if err != nil {
		return
	}

	body, err = doc.CreateElement("S:Body")
	if err != nil {
		return
	}
	err = envelope.AddChild(body)
	if err != nil {
		return
	}

	err = doc.SetDocumentElement(envelope)
	if err != nil {
		return
	}
	return
}

func createBaseNode(doc *dom.Document, name string, cred Credentials, version string) (node types.Element, err error) {
	node, err = doc.CreateElement("sq:" + name)
	if err != nil {
		return
	}

	loginNode, err := doc.CreateElement("login")
	if err != nil {
		return
	}
	err = loginNode.AppendText(cred.Login)
	if err != nil {
		return
	}
	err = node.AddChild(loginNode)
	if err != nil {
		return
	}

	passwordNode, err := doc.CreateElement("password")
	if err != nil {
		return
	}
	err = passwordNode.AppendText(cred.PasswordMD5Sum())
	if err != nil {
		return
	}
	err = node.AddChild(passwordNode)
	if err != nil {
		return
	}

	versionNode, err := doc.CreateElement("version")
	if err != nil {
		return
	}
	err = versionNode.AppendText(version)
	if err != nil {
		return
	}
	err = node.AddChild(versionNode)
	if err != nil {
		return
	}
	return
}
