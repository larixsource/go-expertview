package expertview

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

const (
	DefaultEndpoint = "https://www.expertview-live.com:443/ExpertWebservice-prod/Webservice"
	DefaultVersion  = "2.5.0"
)

var (
	ErrAuthentication        = errors.New("authentication error")
	ErrUnexpected            = errors.New("unexpected error")
	ErrFirmwareNotSelectable = errors.New("firmware not selectable error")
	ErrNoSuchFile            = errors.New("no such file error")
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

type InstallationRecord struct {
	SerialNumber string
	ID           string
	Telematic    string
	HardwareProf string
	SoftwareProf string
	DCF          string
	Firmware     string
	Key          string
	Username     string
}

// ExpertView is a client for the Squarell Expert View webservice.
type ExpertView struct {
	version     string
	cli         *soapCli
	credentials Credentials
}

// NewExpertView returns an *ExpertView, pointing to the given endpoint and using the specified API version and
// credentials. If no version or endpoint are given, the default values are used instead (DefaultEndpoint and
// DefaultVersion)
func NewExpertView(endpoint string, version string, credentials Credentials) (*ExpertView, error) {
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}
	if version == "" {
		version = DefaultVersion
	}
	if credentials.Login == "" || credentials.Password == "" {
		return nil, errors.New("invalid credentials: login and password required")
	}
	ev := &ExpertView{
		version: version,
		cli: &soapCli{
			Endpoint:           endpoint,
			InsecureSkipVerify: true,
		},
		credentials: credentials,
	}
	return ev, nil
}

func (ev *ExpertView) GetFileList() (FileList, error) {
	doc, err := createGetFileList(ev.credentials, ev.version)
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

func (ev *ExpertView) GetFile(filename string) ([]byte, error) {
	doc, err := createGetFile(ev.credentials, ev.version, filename)
	if err != nil {
		return nil, fmt.Errorf("error building xml request: %s", err)
	}

	reqBody := strings.NewReader(doc.String())
	resp, err := ev.cli.call(reqBody)
	if err != nil {
		return nil, err
	}

	if len(resp) == 0 {
		return nil, errors.New("empty response")
	}
	return parseGetFile(resp)
}

func (ev *ExpertView) GetInstallationRecords() ([]InstallationRecord, error) {
	doc, err := createGetInstallRecords(ev.credentials, ev.version)
	if err != nil {
		return nil, fmt.Errorf("error building xml request: %s", err)
	}

	reqBody := strings.NewReader(doc.String())
	resp, err := ev.cli.call(reqBody)
	if err != nil {
		return nil, err
	}

	if len(resp) == 0 {
		return nil, errors.New("empty response")
	}
	return parseGetInstallRecords(resp)
}
