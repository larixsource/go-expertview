package expertview

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

const (
	DefaultEndpoint = "https://playground.expertview-live.com:443/ExpertWebservice-play/Webservice"
	DefaultVersion  = "2.5.0"
)

var (
	ErrAuthentication = errors.New("authentication error")
)

type Credentials struct {
	Login    string
	Password string
}

func (c *Credentials) PasswordMD5Sum() string {
	hasher := md5.New()
	hasher.Write([]byte(c.Password))
	return hex.EncodeToString(hasher.Sum(nil))
}

type DeviceType struct {
	Description   string
	ProductNumber string
}

type RecordKind string

const (
	DCFKind      RecordKind = "DCF"
	FirmwareKind RecordKind = "Firmware"
)

type Record struct {
	Kind RecordKind
	Name string
	File string
}

type FileList struct {
	DeviceTypes []DeviceType
	Records     []Record
}

// ExpertView is a client for the Squarell Expert View webservice.
type ExpertView struct {
	version string
	cli     *soapCli
}

// NewExpertView returns an *ExpertView, pointing to the given endpoint and using the specified API version. If no
// version or endpoint are given, the default values are used (DefaultEndpoint and DefaultVersion)
func NewExpertView(endpoint string, version string) *ExpertView {
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}
	if version == "" {
		version = DefaultVersion
	}
	return &ExpertView{
		version: version,
		cli: &soapCli{
			Endpoint:           endpoint,
			InsecureSkipVerify: true,
		},
	}
}

func (ev *ExpertView) GetFileList(c Credentials) (FileList, error) {
	doc, err := createGetFileList(c, ev.version)
	if err != nil {
		return FileList{}, fmt.Errorf("error building xml request: %s", err)
	}

	reqBody := strings.NewReader(doc.String())
	resp, err := ev.cli.call(reqBody)
	if err != nil {
		return FileList{}, err
	}

	if len(resp) == 0 {
		return FileList{}, errors.New("empty response")
	}
	return parseGetFileList(resp)
}
