// Code generated by go-bindata.
// sources:
// template/ccd.file.tmpl
// template/client.ovpn.tmpl
// template/dh4096.pem.tmpl
// template/iptables.tmpl
// template/server.conf.tmpl
// DO NOT EDIT!

package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateCcdFileTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xc1\x6a\xc3\x30\x0c\x86\xef\x7e\x0a\x91\x9d\x23\xec\x40\xc6\xf2\x04\x63\x87\x8d\x91\x4b\xcf\x6e\x2c\x27\xa6\xe0\x04\xdb\xa1\x0d\xc2\xef\x5e\x62\x1a\xda\x1e\x04\xfa\x05\xff\xa7\xcf\xd9\x61\xf6\xd6\x8d\xf5\xb2\xc6\x09\x98\x01\x7f\xfe\x21\xe7\xb2\xfd\x51\xfa\xd5\xf1\x02\x39\x8b\x0f\x17\xe6\x35\x11\xa8\xae\x41\xf5\xf9\x85\x9d\x44\x09\x4d\xdb\xe2\x31\x52\x30\x3b\x0b\xd8\x93\x71\x81\x86\xf4\x7d\xda\x6b\x05\x5a\x85\xc7\xad\x1e\x75\xa2\xab\xde\xc0\x90\x55\x70\xde\x16\x1d\x63\x6d\xa6\x61\xa9\x04\x33\x90\x37\x7b\x45\x30\x07\xed\x47\x02\xec\xf7\x8f\xf1\x49\x29\x02\xcc\xce\x1b\xba\x01\x82\x2c\x96\x47\x52\x6f\xa9\xc9\xf9\x95\x79\x0f\x00\x00\xff\xff\x5e\x42\x13\x8e\xe5\x00\x00\x00")

func templateCcdFileTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateCcdFileTmpl,
		"template/ccd.file.tmpl",
	)
}

func templateCcdFileTmpl() (*asset, error) {
	bytes, err := templateCcdFileTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/ccd.file.tmpl", size: 229, mode: os.FileMode(420), modTime: time.Unix(1503944154, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateClientOvpnTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xce\x31\x4b\xc5\x30\x10\x07\xf0\xfd\x3e\xc5\x81\x8b\x0e\x79\x1d\xdc\xa4\x08\xe2\xa0\x20\xea\x9b\x74\x10\x87\x34\xbd\xbe\x86\x97\xde\x85\xe4\x5a\x88\x8f\x7e\x77\x49\x2b\x6e\x77\xbf\x3b\xf8\xff\xaf\x50\x47\x9f\x51\x96\xc8\x38\xf8\x40\xe8\x33\xda\x59\x65\xb2\xea\x9d\x0d\xa1\xe0\x89\x98\x92\x55\xea\xb1\x2b\xf8\xf5\xfe\x71\x7c\xfd\xbe\x1e\x55\x63\xbe\x6b\x9a\x93\xd7\x71\xee\x0e\x4e\xa6\xc6\xd9\xbe\x91\x25\x4e\x37\x00\x2e\x78\x62\x85\x9e\x16\xd4\x99\x21\x26\x51\xc1\xcb\x05\x0f\xc7\x6d\x5a\x57\x48\x34\x89\xd2\x66\xcf\x92\x95\xed\x44\xb8\xae\xfb\x8f\x24\xdd\x5f\xb2\x84\xc5\x24\xd2\x54\xd0\xf3\xe0\xd9\x2b\x01\x67\xe3\x28\xa9\xd1\x12\x09\x33\xa5\x85\x12\xb0\x74\x9e\x7b\x88\x94\xb2\xcf\x6a\xce\x54\xfe\xe7\x1a\xef\x64\x8a\x26\xfc\x08\x2c\x94\x3a\xbc\x05\x3b\xeb\x68\x58\x9c\x75\x23\x01\xb4\xce\xde\x43\xcd\x7d\x7c\xc0\x75\x6d\x9b\xba\xb6\x35\xe2\x4f\x69\x6b\xd3\x36\x3b\xb5\x67\x2a\xfb\xe1\x85\xca\xe6\x1b\x54\xf1\x03\x1e\xde\xe4\xe9\x73\xeb\x2e\xb3\x92\x61\x89\x73\x08\xf5\x46\xdc\x57\xfe\x0d\x00\x00\xff\xff\x94\xe5\xf3\xfc\x6c\x01\x00\x00")

func templateClientOvpnTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateClientOvpnTmpl,
		"template/client.ovpn.tmpl",
	)
}

func templateClientOvpnTmpl() (*asset, error) {
	bytes, err := templateClientOvpnTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/client.ovpn.tmpl", size: 364, mode: os.FileMode(420), modTime: time.Unix(1504303885, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateDh4096PemTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd4\x35\x12\xa4\x08\x00\x40\xd1\x9c\x53\x4c\x4e\x6d\xe1\x16\xe2\xee\xde\x19\xd0\x38\x8d\xfb\xe9\xb7\x76\xe2\xfd\xe1\x3f\xc0\xfb\xe7\xbf\x38\x51\x56\xad\x3f\x82\xf2\xc7\x61\x3d\xd6\x14\x03\xd1\xf3\xff\x7e\xc0\x54\x55\x51\x70\x75\x9e\xad\x45\xb6\x2f\xb3\xc9\xd2\xe7\x37\x56\x43\x1c\xa4\x1b\x2d\xac\xa4\x52\xd1\xb4\xbe\xa5\x86\xf7\xd6\x99\x21\xd2\xde\x29\x40\x6c\xec\xf4\x04\x64\x3e\x11\x13\xee\x7e\x38\xf0\xc5\x28\x74\x5d\xa8\x62\x55\x98\x61\x1a\xea\x16\x8e\x9e\xaf\x9b\xd5\x15\x48\x63\x88\xf6\xd1\xa8\x1f\x03\x39\x69\x95\x7a\x63\x88\xe6\x15\x81\x75\xb8\x4a\xdb\x34\x5f\x4b\x1c\xdf\xc5\x06\x57\x17\xc0\xf7\x75\x64\x4a\xc5\x2d\x99\x2d\x90\x54\x89\x30\x0f\x94\x8f\x86\xce\xa5\xda\x3e\x8b\xaa\x8d\x2f\x7c\x14\x4d\xab\xeb\x97\x43\x29\x3e\xd4\xd9\x9e\x62\xf5\x10\x0d\x6a\xcb\xbf\x7c\x5f\xfa\xb5\x77\x01\x52\x8d\x90\x43\x92\xb3\xfa\x3a\xdd\x57\xdd\x33\x6c\xbb\x15\x03\x5d\xe9\xe4\x5e\xb0\x56\x9e\x92\x31\x26\xb3\x57\x2e\x0a\xf2\x73\x82\xf0\x0c\xd7\x9f\xd8\xef\xa7\xfc\x40\xe5\xf0\x9b\xa9\x9b\x42\x02\xdd\xf1\x83\xa0\xe0\x73\x4e\xc7\xd5\x0b\x62\x43\xc2\xc7\x2e\x11\x97\x73\x4d\x8b\xf9\x94\x4c\xd6\x84\x76\x79\xbf\x20\xb5\xab\xd4\x51\x22\xd0\xb0\x96\x60\x63\x64\x61\xec\xb6\xf5\x2b\x96\x04\x99\x03\x95\x47\xc4\xe1\xc7\x6b\x2e\xab\x66\xde\x60\xac\x67\x29\x30\x92\x96\x55\x31\x46\xa3\x6f\x25\x0f\xdb\x2f\xd6\x9f\x10\xef\x0b\x0e\xbf\x90\x49\x13\x9a\xf1\xc3\x3f\x06\xb7\xf9\xfe\xb7\x40\xd4\x6e\x01\x66\x8f\xa4\x08\x8e\x5b\x8c\xcb\x9f\xf0\xc9\xb6\xa0\x09\x09\x6a\x33\x15\x30\x08\xbb\xfd\xac\x4f\xfc\xd0\x38\x7d\x5b\x48\x07\xe3\x3b\x2a\x63\x6b\xd0\x7a\x03\x72\x20\x5a\xd6\xec\xa7\xcc\x9e\x30\x07\xe6\x25\x22\x07\x29\x4c\xa4\x09\x49\xba\x7e\x88\xdd\xad\x58\xd0\xfa\x3a\xf4\xf3\xa2\x12\xfe\x0b\x3b\x63\x9a\x7d\xe1\xf5\x36\xa2\x1c\xca\x8c\x27\xd5\xdf\x13\x2f\x56\x83\xf0\xbc\x0d\xba\xed\xa2\x02\x6c\x54\xc0\xa0\x09\x1f\xbe\xd8\x18\xbb\x44\xcb\x11\xad\x82\x53\xbc\xcc\x0e\x63\xe6\x41\x70\x27\x10\xb5\x3c\xda\x12\xfd\xab\x2d\xf9\x26\x71\xc7\xe4\xc0\x98\x70\xcb\xfe\xfd\x40\x59\xf6\xdb\x18\x1d\x30\xe1\x62\x4f\x14\xea\x8c\xbe\xee\x59\x80\x2b\x4d\xea\xc2\x0e\xa9\x6e\x32\xcd\xc7\x34\x25\x88\x8e\xf7\xbf\x66\x09\xbb\x3e\xd6\xd6\x6e\x87\xc3\x33\xf2\x66\x74\xeb\xd5\xd4\xb7\xc4\x72\xe8\x58\x18\xb0\x8c\x92\x87\x7a\x38\xe5\xe0\xd5\x8a\x3f\x37\x16\x7a\x35\x96\xf8\xd0\x09\x4f\x9a\x0d\x3a\x3b\x64\x13\x21\xda\x1e\xd3\xac\x7c\x66\xf4\x88\x1b\xc6\x28\x5d\x70\x69\x37\xbe\xae\x55\x56\xdb\x7f\x27\x80\x2b\x14\x4a\xd9\x50\x26\x14\xfb\x83\x14\xb4\x8e\xa1\xb3\x2e\x3f\x3f\xd0\x90\xb9\xfb\x66\xc6\xe3\x30\x83\x6e\xce\x08\x0b\x8a\x2f\xe8\xda\x95\xc2\x8d\xf5\x68\x42\x94\x6f\x3f\xca\x13\x67\x27\x05\xc0\x64\xa4\x4a\x11\x14\x5e\x69\xd2\x71\x89\xaf\x5b\xe2\x39\x8e\x47\x69\x0f\xee\x20\x33\x1b\xf7\xc6\x05\xd0\x8a\x91\x6e\x8b\x94\xd2\xb3\xa6\xb5\xd9\x4c\xad\x51\x77\x55\xad\x21\x7b\xb6\x26\x22\x03\x28\x1a\xd9\x19\xda\x0f\x12\x54\x7d\xba\xaf\x83\x48\x3c\x25\xbc\x90\x54\xa3\x74\xa1\x5b\x2e\x44\xfd\xc8\xa5\x29\xc6\x4d\x78\x6f\xcf\xe9\x3c\xf2\xd9\x3e\x7c\xab\x38\xf9\x6f\x3d\x2d\xed\xd5\x46\x07\xa0\xda\x1a\x19\xcd\x3c\x57\x5a\x88\x1d\x40\x87\xd1\x91\x6a\x5a\x3d\x96\x00\x23\xbb\x22\xaa\x33\x70\xf6\x85\xfc\xdd\x93\x4b\xd0\x44\xc1\xaf\x46\x8a\x89\x8b\x5b\x82\x95\x33\x84\x6f\xc1\x19\x50\x0e\x54\x66\x34\x5e\x0c\x9c\x53\x9b\x1b\xa9\x8c\x28\x26\xf7\x3c\x13\xe8\xa1\xa6\xcc\xbd\x07\xef\x56\x65\x3a\x26\x9f\xd8\xfc\x0b\x4a\x49\x2c\x8b\xe7\x9b\xa1\x90\x7d\xd1\x93\x9c\xbe\x53\x3e\x12\x3e\x0b\xd8\xcc\x2d\xce\x05\xfa\xb3\xc2\xc1\xb6\x02\x7c\x6f\x38\xa7\x81\x30\x2b\x75\xe4\xcf\xa8\xb6\x42\x6a\xd1\xa7\xd2\xe9\xad\x3f\xa6\x45\xea\x6d\x79\x3c\xdf\x67\x20\x7e\x3c\x4b\xec\x55\xb2\x21\xbe\x11\x40\x63\xe2\xb9\x59\xa6\x3d\xc9\xf4\xf7\xd0\xdc\x66\x38\xdc\xbb\xcf\x60\xbb\x91\x6a\x38\x0c\xe4\xa6\x2c\x10\x24\xd9\xe0\x36\xd8\x5b\x9a\x3f\x47\x7b\xd6\x97\xa7\x46\xcf\x64\xb0\x47\x58\x65\x70\x07\x50\xa3\x82\xbd\x64\x78\x9e\x8d\x91\x76\x1e\x6c\xe3\xd6\xd4\x26\x14\x89\x98\xfd\x7d\x9a\x45\x27\x28\xca\x67\xe0\xfd\xbe\xb7\x03\x0c\x19\x22\x5d\x64\x0f\x3c\x30\xb9\x55\xfc\x34\xcc\x8d\x8e\x6a\x04\xb0\x8a\x7c\xbb\x51\xd0\x97\xcb\x4d\x1f\x54\xcf\x28\x35\x8c\xee\x09\xa3\xfc\xc8\x8b\xce\x24\x89\x8a\x79\xc7\xd0\x34\xa3\x04\xcd\x58\x19\xec\x7d\xc6\x79\xb2\x9c\x21\x47\x1a\x0b\x52\x10\xfd\x89\x01\xea\x98\x67\x1d\x43\x35\xa9\xbd\xc8\x99\xd6\xf7\x63\xee\x77\xd4\x7e\xbf\x44\xd3\xc0\xd6\x46\x9e\x40\x01\x43\xf8\xf7\x78\x8e\x78\x71\x15\x4d\xbd\xd7\xdb\x64\x47\xe8\x60\x85\x3b\xcd\x70\x74\x5f\x01\xe9\x6e\xca\xfc\x55\xbd\x3a\x07\xc7\x1b\xe6\xc7\x18\x92\x3d\x07\xd3\xbd\x42\xfd\xa8\x2b\xd8\x6e\xdd\xd2\x23\x19\x65\xfe\x45\x17\xf8\xeb\xb0\x68\x09\xff\xa7\xf3\xbf\x01\x00\x00\xff\xff\x2c\x8e\x5a\x0b\xbc\x05\x00\x00")

func templateDh4096PemTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDh4096PemTmpl,
		"template/dh4096.pem.tmpl",
	)
}

func templateDh4096PemTmpl() (*asset, error) {
	bytes, err := templateDh4096PemTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/dh4096.pem.tmpl", size: 1468, mode: os.FileMode(420), modTime: time.Unix(1503944154, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateIptablesTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func templateIptablesTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateIptablesTmpl,
		"template/iptables.tmpl",
	)
}

func templateIptablesTmpl() (*asset, error) {
	bytes, err := templateIptablesTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/iptables.tmpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1503944154, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateServerConfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x5a\x5f\x6f\xdb\xb8\xb2\x7f\xd7\xa7\x18\xd8\x0f\x4d\x00\x5b\x71\xdc\x76\x4f\xbb\xb9\x7f\xe0\x4d\xdc\xad\xb1\xad\xe3\xdb\x38\xdd\x5d\x60\x5f\x68\x69\x6c\x11\x96\x48\x2d\x49\xd9\xf5\x39\xd8\xef\x7e\x31\x43\x52\x92\x93\x74\x0f\x4e\x81\xa0\xb6\x44\x0e\xe7\xff\xfc\x66\xe8\x9b\x5a\x1b\x07\xd7\xd7\xef\xdf\x24\xfc\xe9\x5f\xff\x82\x74\x45\x1f\xfe\xfa\x2b\x49\x86\xb0\xbe\x5d\x81\x36\xf0\x78\xb7\x02\x8b\xe6\x80\xe6\x7f\x93\x9b\xda\x68\xa7\xc1\x65\x75\xe2\x3f\xf1\x16\xfe\xe4\xf7\x0c\x72\x3c\x80\x6b\xd4\x00\x8e\xb2\x2c\x21\x33\x28\x1c\x82\x00\xa3\x1b\x87\x39\x2c\x56\xf4\x52\x61\x39\x6a\xd7\x8a\xfa\xc9\x5a\x05\xe8\x0a\x34\x0a\x5d\x58\x9b\x26\x43\x78\xb4\xd8\xae\x9f\x0c\x40\x6e\xe1\xa4\x1b\x10\x06\xbb\xc5\x1b\x23\xf3\x9d\x54\xbb\x64\x08\x42\xe5\x50\x88\x03\x42\x6d\xd0\x93\xcd\x41\xf0\x56\x38\x48\xe3\x1a\x51\x82\x54\x0e\xcd\x56\x64\x18\x96\xf3\x6e\xcc\x41\x3a\x38\x4a\x57\x10\x79\xd3\xd1\x6e\x57\x13\x2f\x0b\x7f\xf8\x51\x28\x07\x4e\x43\xa6\x95\x33\xba\x04\x91\x65\x68\x2d\xd4\xba\x94\x99\x44\x9b\x0c\x41\x1f\xd0\x80\x2b\x10\xbe\xae\x96\x23\xde\x53\x35\xd6\x45\x41\xb7\xd2\xe0\x51\x94\x65\x32\x04\xd3\x94\x68\x61\xab\xfd\x6a\xfa\x5b\x3f\x2e\xaf\xd6\xb3\xd5\xf9\xc1\xf7\x0a\x94\x56\xe3\x5f\xa5\xca\xf5\xd1\x82\x3d\x59\x87\x95\xf5\x94\x33\xa1\x60\x27\x0f\x5e\x1c\xc0\x6f\x35\x71\xe1\xa0\x51\xd2\x81\x6a\xaa\x0d\x9a\x11\xd8\x26\x2b\x40\x58\x52\xeb\x24\xd0\x0b\xb4\x46\xd0\x04\x05\x8f\x95\xce\x71\x10\x78\x91\x36\x2c\xab\xb4\x75\xdd\x79\x41\x24\x6f\x36\xa5\x1d\x6c\x1b\x95\x39\xa9\x55\x32\x84\x46\x95\xa4\x05\x62\xa9\x16\xc6\x49\x51\x96\x27\xf2\xa2\x6d\x43\x1f\x72\x69\xc5\xa6\x24\x26\x89\x46\xd4\x40\x2b\xf9\x0b\x52\xdf\x04\xa3\x27\xc1\xb1\xc8\xc9\xa2\x02\x14\x62\x6e\xfd\xc6\xd9\x8a\xd4\xf2\x7a\x0a\x22\x17\xb5\x43\x03\x4a\x54\x74\xcc\xd6\xe8\x8a\x57\x2c\xd1\x1d\xb5\xd9\xc3\xad\x56\x0a\x99\x59\x0b\xb5\x50\x58\x06\x5f\x4a\x86\xde\x63\x2a\x6d\xc8\x02\x42\x81\x56\x98\x02\x09\xff\xdb\x0a\x1e\x56\x53\x12\xa2\x90\xbb\x02\x0d\xf9\x2e\x1b\x53\x9c\x98\x05\x72\x02\x8b\x25\x51\x3d\x60\x27\x24\x1d\xdb\x63\xf6\xb9\xb0\xb3\x55\xe4\x96\xb4\xbc\x7c\x6e\x59\x68\x6c\xc3\xfa\xcb\xb5\x7a\xe5\xc2\x59\x6c\x94\x9b\x68\x29\xf8\x7c\x5a\x8b\x9a\x94\xf2\xf0\xf0\xe9\x6a\xfd\xe9\x01\x8c\xd6\x0e\x32\x34\x4e\x6e\x65\x46\x7e\x76\x91\x89\xcb\x51\xff\x49\x32\x84\x0b\xfa\x7a\x39\x62\xcf\xaf\x8d\x3c\xd0\xba\x3d\x9e\xe0\x62\x8f\xa7\xcb\x14\x60\x2e\xb2\x02\xb2\x52\xa2\x72\x21\x3e\x88\x63\x9f\x05\xbc\x13\xb3\xb2\x5c\x81\xd2\x80\x3e\x2a\x26\x4f\xeb\x92\x21\xd3\xd9\xca\x92\x94\xb7\xee\x36\x11\x0d\x92\xde\x13\xb5\xec\x3b\xe4\x2e\xd6\xfb\xbb\x15\x15\x42\x26\xfc\xc6\x64\x48\xf2\xa0\x7f\x33\x40\x61\x4f\x63\x63\xc5\x00\x72\x69\x30\x73\xda\x9c\x58\x87\x82\x48\x87\x40\xdb\x82\xcd\x8c\xac\x9d\x0f\xa2\x1d\x2a\x34\xc2\x49\xb5\x83\x2f\x0f\xb3\xbe\xe8\x36\x48\xd3\x93\xd9\xa6\x00\x5f\xb0\x42\x0a\x11\xb2\x64\x63\x39\x86\x28\x72\xfe\x6c\x10\x6e\x75\x55\x69\x05\x4b\xe2\x2f\x5a\xce\x8b\x14\x48\x21\xa9\x4a\x6f\xf9\x85\x17\xee\xec\x40\x2f\xcc\x4c\x9d\xe0\xb7\xb7\x93\xf7\xac\x9c\x4a\x28\xb1\xc3\x8a\x56\x7a\x43\x73\xf0\x6e\x90\x8e\xce\x39\xe0\x6a\x54\x14\x5f\xf4\x58\x94\x96\x79\x02\x01\xab\x5f\x6e\x1f\x60\x78\x3d\x25\x3e\x2a\xe1\x28\xa9\x45\x5d\x93\x49\x2d\x22\x0c\xea\x7d\x66\xaf\xa7\x51\x55\xf2\x80\x20\x15\x1d\x08\xb5\xd8\xe1\x65\x9a\xdc\x64\x02\xa2\x42\xaf\x48\xfa\xab\x4c\xa4\x99\x71\xc9\x0d\x5b\xf0\xfc\x95\x97\xd3\xbf\xa6\x93\x5e\x7c\x4b\x2f\x60\x08\xeb\x42\x5a\x66\x05\x6c\xa1\x9b\x32\x27\x79\xf6\x58\x3b\xb0\x94\x81\x5d\x92\x64\x82\x8b\xc5\xed\xec\x16\x8d\x5b\x09\x57\x50\xc5\xe0\x43\xf9\x71\xef\x21\x51\xa4\x67\xbf\xe0\x29\x3e\x4a\x86\x70\x27\xb7\x5b\x89\x50\x60\x59\x7a\x79\x8c\xa8\xd0\xa1\xe1\x14\xf5\xb3\xb7\x38\xfa\xb4\x4d\x0e\x49\x49\xfc\xc7\x64\x08\x00\xba\x46\x65\x6d\x09\x79\xc1\x7b\x60\xac\x1b\x07\x79\x71\x3d\x99\xbe\x49\x6b\xac\x80\x3e\x90\xbf\x35\x1b\xeb\xa4\x6b\x1c\xc2\x74\xf2\xe6\x1d\x1b\x9b\x5e\xf5\x6b\x4d\x63\x7d\x7d\xe1\x05\x1b\xe9\xbc\xff\x24\xc3\xbc\xe8\x11\x4c\x6e\xf2\xe2\x89\xaa\xf2\x82\x76\xf0\xcb\xbc\x60\xd9\xee\x3e\xae\x88\x19\xdb\x13\x30\x66\x28\xa7\x6b\x5d\xea\xdd\x89\x78\x6a\x55\x69\x9b\x0d\xd5\xa1\x0b\x91\xe7\x06\x2d\xb1\x01\x07\x29\x60\xb1\xba\xec\x72\x6e\xcc\x1e\x31\xc4\x0e\xd3\x74\x92\xbe\x67\x27\x2d\xf5\x11\x4d\x88\x59\x9d\x0c\x3d\xc5\x9a\x8a\x3e\xe6\x70\xe1\x0a\x54\xa0\xd0\xbd\x9e\x8c\x40\xa6\x98\x82\x80\xab\xd7\x13\xa8\xd1\x04\x5a\x74\xc8\x1d\x6e\x45\x53\x3a\x4b\x31\xc2\x6b\xe1\x82\x52\xbf\xc1\x4c\x57\x15\xaa\x1c\xf3\xcb\x24\xb2\x1e\xd8\x25\xa9\x6e\xb5\xda\xca\x5d\x63\xba\xec\x41\x59\x8b\x78\x22\x06\xca\x13\x08\x2e\x25\x61\xc3\x90\xd5\x1e\xfd\xdf\x69\xc8\x8d\x38\xc6\xb0\x0a\xb2\x53\xa5\x34\xba\x22\xb3\xf7\xf2\x0b\xd7\x22\x27\xf6\x08\xd7\x93\xf4\x5d\x3a\x49\xaf\x99\x94\x74\x16\xcb\xed\x28\x14\x1c\x83\xd6\xf9\x95\x1b\x84\x4a\x10\x23\x07\x21\x4b\x9f\xad\x75\x54\x1c\x51\xee\x25\xc0\x76\x43\x5c\x66\x38\xe4\xcf\x12\x81\x56\xed\xb1\x29\xa7\x0c\xda\x47\xa9\x1a\x4a\xa9\x10\xc8\xe3\x3a\x37\x4a\x86\xcf\x41\x4b\xda\xe6\xbb\x18\xac\xcc\x3e\x57\x23\xa9\xb6\x3a\x4d\x6e\x82\xa0\xe1\x9c\x09\x4c\xdf\xbe\x4d\xe3\xdf\x24\x09\x6f\xc9\xb7\x96\x48\x08\x8e\x3f\x7e\x16\x76\x1f\xdc\xeb\xb3\x90\xca\x09\xa9\x08\x8c\x61\xa6\x4d\x4e\x49\x2b\x48\xf8\x5f\xe3\xff\x69\x81\xd1\x62\x15\x15\x4d\x09\xce\x5a\x9d\x49\xe1\xcb\xa5\x54\x5e\xa4\x90\xdc\x17\xdb\xd6\x50\x3b\x8d\x16\x72\x0a\x3b\x4d\xda\x90\x96\x35\x2d\xc8\xbb\x46\x7c\x9a\xaf\xb9\x6a\xd7\xfa\x66\xc8\x78\xc2\x5a\xb9\x53\x98\x07\xfb\x70\x15\x78\xce\x48\x57\xc5\x6b\xad\x4b\xaa\xce\x0e\x8e\x82\xf8\xab\x0d\x1e\xa4\x6e\x2c\xf9\x51\x20\x95\x26\x72\x9b\xb1\xd3\x8d\x69\xf5\xb8\x46\x63\xa5\x75\x20\xeb\x3a\x75\xdf\xfe\xc6\x27\x49\xdf\xcf\xed\x92\x0c\xe1\xf7\x08\xdc\xb6\xd2\x58\xc7\xc9\x98\xd3\xcc\xfd\xc3\x2b\xdb\x2e\x84\x4c\xd4\x62\x23\x4b\xe9\x28\x70\x9d\x0e\x80\xb2\xad\xf4\x2d\xa4\xf1\xf0\x92\x1e\xc7\xc3\x28\xf6\x17\xb7\x3d\xd0\xc3\x85\x53\x75\x80\xb1\x12\xca\xe3\x00\x4b\x98\x98\x51\xc5\x62\x75\xa5\xd0\x55\x64\x5e\xad\x98\x5a\x38\xaf\xa5\x32\x82\x02\x0d\xc2\x11\xbd\x19\x9b\xaa\x8d\x8c\x37\x57\x67\xae\x93\x02\x7c\x90\x8a\xe9\xf3\x62\x3e\x92\x4e\x12\x56\x72\xa8\x92\x25\x8c\x50\x4c\xdc\x7b\x40\x1b\xad\x17\x6c\xe5\xff\x0e\x84\xdf\x4e\x00\x55\x1e\xbf\x5d\x4f\x26\x97\xa4\x08\x51\x96\x3a\xc0\x0e\x8f\x96\x9f\xf8\x42\x0a\xf0\x09\x3d\x9a\x88\xf1\xe2\x53\x8a\x63\xbf\xa0\xd8\xe9\x81\xca\x17\x21\x7f\x1b\x1d\xe3\xa0\x85\x28\xe9\x79\x90\x40\xc7\x67\xc7\xe3\x7f\xe8\x10\x8c\x59\x88\x7b\x01\x77\x1f\x6f\x57\xe3\xda\xe8\x6f\xa7\x11\x1c\x59\xd9\xd1\xbb\x9d\x28\xf7\x5e\x5c\xb2\x4c\x8c\x92\xc0\x22\x6b\x95\xf6\x76\xe9\x83\x93\x4a\x86\xb2\x85\x54\x3d\xdf\x0f\xea\xf3\xe8\x9a\xf2\xe6\xdd\xf2\xa1\x05\x55\x31\x23\xa6\xf0\x82\x97\x7a\x98\xfa\x7d\x3f\x7d\xee\xa5\x14\xbb\xdf\xf7\xd3\x27\x5e\xca\x98\xd5\xe1\x8f\xde\x6e\xac\x31\xad\xc8\x89\xb4\xd9\x5b\xf2\xca\xa8\x8d\x8b\xd0\x77\x74\x60\xf8\x32\x2a\xac\x03\x4d\x5e\x2d\x3d\x4c\x0c\x92\x36\x6c\x74\xa3\x18\x60\x7b\x75\x87\xc5\x4f\xed\x4d\x36\x5c\x35\xb6\xf0\x6d\xa6\x8d\x8a\x0f\xd9\x2d\xf8\xe0\x11\xa4\x8b\xba\x66\xcc\x46\x62\xb5\x48\xd0\xbb\xb4\x85\x0d\x16\x52\xb5\xf9\xc8\x43\x9c\x3e\x40\xa4\xdc\xe3\x0a\x64\xed\x3e\xdd\xcb\x85\x82\x01\x1b\x41\x75\x7f\xd6\x5e\xe9\x23\x9f\x49\x9c\x9d\xb9\x43\x07\xaf\x83\xa5\x39\xb9\x5d\xc4\xfc\x7e\x1e\xa4\x54\x81\x37\x22\xdb\xbf\xec\x53\x69\x32\xf4\xf4\xaf\xdf\x4f\xd3\xeb\x1f\xde\xa5\xef\x9f\x15\x88\xe8\xf1\xd3\x67\x4b\xaf\xbf\xb7\xf4\xed\xb3\xa5\xd3\xef\x2d\x7d\xfd\x6c\xe9\xeb\xef\x2d\x7d\x93\x0c\x6b\x32\xd5\x20\x6c\xf8\x07\x6d\xa0\x40\x7c\xba\x61\xf0\x64\x61\xa0\xfc\xc3\x4b\x0b\x61\xad\x43\x01\x00\x5b\x63\x46\xe0\xbb\x17\x43\xde\x23\xe2\x8b\x64\xd8\x7a\x26\x01\x84\x2d\x88\xe7\x49\x09\x0a\x61\x41\x44\x03\x27\xc3\x88\xbf\xbc\x77\x80\x74\xde\x11\x02\xd4\x65\x93\x33\xb8\x22\x93\xf8\x39\xc0\xa8\xdf\xdc\x34\x9b\xae\x75\x19\x64\x59\xee\xfb\xeb\xe8\xf8\x3d\xbe\x42\x1e\xe2\x70\xe7\x3a\x6b\x3d\xae\x7f\x19\x14\x5c\xa6\xe4\xf8\xf3\xdf\x66\x9f\x57\x9f\xe6\x3f\xc2\x03\x21\x3a\xdb\x0f\x2a\xdf\xd4\x92\x58\xfc\xac\xd7\x14\x66\xbe\xad\xa1\x2e\x19\x06\xeb\x02\x4b\xad\xa8\x8e\x92\x2a\x83\x34\x24\xbf\xad\xa8\x65\x3b\x97\x9d\x82\xbd\xd3\x17\xd5\x0b\x91\x15\x52\x61\x37\x5e\x88\x96\x7a\x33\x49\xaf\xa7\xef\xce\xfc\x78\xfa\xe6\x1d\xe5\x8d\x0f\x94\xa2\x46\xd0\xa8\x90\xe8\x39\xc9\x73\x54\x71\xfe\xb7\x3f\x26\xc3\xa0\x9c\x50\xc9\x73\x69\x20\xcb\xf2\x7f\xef\xe6\x1e\x14\xaa\x6e\xf0\xc4\x2d\x49\x96\xe5\x57\x9d\x90\x31\xb3\x85\x6a\xe3\xbb\x04\x79\x4e\xd9\xf3\x0e\x4f\x78\x4f\x42\x9f\x13\x02\x9d\x72\x4a\x47\xf6\xd5\x93\x7c\xe0\x91\x76\x98\x0a\x85\x91\x09\x97\x76\x69\x01\xbf\x89\xaa\x2e\xd1\x13\x6a\xb3\x26\xa5\xdf\xae\xd9\x20\x86\xa4\xda\x8d\x78\xc2\x12\xd3\x77\x00\xe8\x1d\x92\xf4\xc5\xa8\x1b\xbb\x51\x85\x18\xf8\xac\xd0\xeb\x02\xed\xcb\x9e\xd2\x1f\x63\x85\xe1\x51\x4f\x4d\xa4\xbc\x6f\x98\xb3\x53\xf7\x0a\x92\xde\x52\x28\xbf\x67\xac\x1b\x6d\xf9\x77\xa6\xbc\x79\xd9\x94\xcf\x9f\x72\x1f\x78\x7b\x17\x5b\xa2\x60\x49\x91\xe7\x3d\x64\x40\x08\xe2\xcc\x98\xc1\x7a\x2d\xe4\xa3\x8c\x11\xd9\x8b\x1f\xa6\x49\x92\x99\x72\x7c\x40\x23\xb7\xbe\xb5\xbc\xfd\xf2\xa9\x3b\xa6\x8b\x1b\xe1\xce\x54\x82\x8a\x21\x7f\x2e\xb7\x5b\x34\x3e\x9c\xda\x29\xce\x93\x69\x1f\x07\x66\xbb\x10\x76\x46\x37\x75\x98\x4c\x74\x30\x67\xcd\x35\x8f\x6c\xeb\x8e\x1a\x2a\x74\x85\xce\x59\x80\x8b\xeb\x4b\xf8\xd2\x28\xa8\x9a\xd2\x49\xf2\x8b\x98\xde\x73\x81\x95\x56\x76\x04\x5a\x05\x3c\x22\xb2\x82\x25\x06\x7f\x86\x1f\xe1\xb4\x5c\xbd\x38\x44\x0b\xeb\xe3\x76\xbf\xf1\xca\x93\x06\x51\xd7\x46\xd7\x46\x0a\x87\xe5\x89\xec\x79\x31\xbd\x84\x8b\x59\x7e\x10\x2a\xc3\xfc\x12\x6e\x63\x20\xf9\x01\x0b\xf7\x64\x27\x25\x2a\x99\x11\x62\x0c\xa4\x2b\x9d\x93\x66\xcf\x66\x7a\x52\x51\x13\x50\x6b\x65\xd9\x6a\x5e\x5f\x91\x15\xc2\xf3\x9d\xb6\x3a\x0d\x3d\xf8\x5c\x17\x96\x3d\xcf\x78\x04\x2b\x4a\x14\x46\x8d\xa3\x37\x7a\xb6\xd2\xe4\xe6\xfc\x71\x7a\xe5\x5f\x24\x7e\x5e\xeb\x2d\x99\x8f\xbc\x27\x75\xc3\x11\x3f\x78\x8e\xe8\x8f\x93\x5f\x37\xa2\x62\xa4\xe0\xd7\x06\x58\x96\xfb\x1e\x38\x19\x52\x07\xcc\xcd\xfa\x4e\x38\x3c\x0a\x12\xdd\xe8\x66\x57\x74\xd3\xde\x4c\xc4\x61\x01\x91\x5c\xac\xc0\x19\xb1\xa5\xb2\x14\xb3\xe4\x11\x37\xb0\x31\xfa\xe8\xc1\x24\x43\x8e\x88\xee\x4a\xad\xf7\x4d\xcd\x0c\xec\xf4\x53\xd2\x64\xa2\xf5\x33\x04\x10\xb3\xf0\xd9\x50\x72\x39\x5b\x93\x0b\x9a\x33\x94\xf7\xd4\x3b\x22\xa2\xe0\x07\x1e\xd5\x4b\xea\xe3\x72\x34\xed\x14\x98\xd6\xb0\xb8\xe4\x2b\x68\xca\xd3\x65\x9a\xdc\x84\xe2\x1c\x34\x34\x8e\x9a\xc8\x71\x7b\x0d\x9b\x53\x2d\xac\x1d\xe7\x45\x56\x0f\xfe\x83\x95\x04\xc5\xd1\x70\x97\x1a\x90\x62\x5b\x1d\x5b\x8d\x5b\x74\x94\x16\xc9\x95\x42\x13\x49\xe4\xbd\xc0\xc1\x72\x5d\x29\xba\x5b\x3e\x78\x0d\xfc\xba\x78\x19\x35\xdf\xce\xbe\xce\x67\x6b\x8a\xc1\xc2\xb9\xfa\xc7\xab\x2b\x5d\xa3\x3a\xd4\x2a\x55\xe8\xae\xb6\xe2\xcf\xb4\x70\x55\x39\x24\xee\x32\x71\x40\xe1\x6c\x98\x3a\x74\xc8\x62\x83\x54\x06\x0c\x6e\xfd\xf8\x90\xbb\xd4\x66\x53\x72\x3d\xef\xa0\xba\x25\xd5\x1d\x64\x8e\x39\x6c\x4e\x3c\x92\xca\x95\x4d\x33\x5d\xb5\x8a\xa4\x33\xc6\xba\xe6\xca\x4f\xfb\xa6\x93\x77\xe9\x0f\xff\x48\xa7\xd3\x29\xfd\x0d\x92\x97\x57\xf1\x2c\x69\xf9\x00\x7f\xfd\xc5\xda\x7b\x6c\xd3\xf0\x13\x47\x6f\x31\x70\x3f\x99\xf5\x1c\xbd\x37\xde\x18\x58\xc4\x01\x74\x10\x99\x92\xc2\x4f\xa7\xe8\xfb\xa3\xb3\x09\xae\xaf\x5d\x36\x8c\x2e\x5a\x30\x4a\x70\x6c\xab\x4d\x86\xfd\x23\x5e\x58\x3a\x0a\x93\xf7\x73\xe0\xcc\xcc\xf6\xd3\xd2\x59\x7e\x23\x30\xc6\x9b\x5f\xd9\x97\xee\x0b\x42\x59\x71\x7a\x1c\x20\xd0\xdf\x69\x45\x6e\xbb\x94\x1b\x38\x25\x40\x23\x77\x85\x8b\x18\xa7\x6b\x83\xfc\x88\xba\xc3\x50\x57\x7b\x3c\x71\x45\x20\x94\x46\x70\xae\x83\x54\x36\xd6\x79\x9e\x84\xb4\x03\x32\x1e\x14\x95\x7e\x7a\xed\xd0\x32\xde\xac\x1b\x43\xc5\x87\x76\x7c\xd0\xd4\x8e\xe8\xbc\xe1\xeb\x09\x42\x8f\x04\x22\xb1\x37\x89\x0a\x78\x93\xa1\xa6\x24\xfc\x1a\x26\xef\xe7\x1c\xd5\x42\x1a\x3f\x76\x5e\x7c\x80\xdf\xef\x1f\xe1\xe3\xec\xeb\x1c\x96\xf7\x6b\xf8\x79\xbe\x9c\x7f\x99\xad\xe7\x77\xb0\x58\xde\x2d\xbe\x2e\xee\x1e\x67\x9f\x28\xe6\xe6\x5f\xd6\x8b\x0f\x8b\xdb\xd9\x7a\x7e\xf5\xcb\xfc\x77\x58\xcd\x16\x5f\x1e\xe0\xc3\xfd\x17\x98\xcf\x6e\x3f\xc2\xed\xa7\xc5\x7c\xb9\x26\x5e\xf8\xeb\xc7\xd9\xd7\xc5\xf2\x67\x58\xac\x1f\xe0\xfe\xd7\x25\x3c\x2e\x17\xff\xf7\x38\x87\xc1\xed\xfd\xe7\xcf\xf7\x4b\x58\xce\x3e\xcf\x07\xb4\xf6\x71\x49\x4f\xe6\xcb\x35\xac\x3f\x2e\x1e\xe0\xd3\x62\x39\x87\xfb\xc7\x75\x9a\xdc\xe4\x4d\x5d\x32\xbf\xe3\x8c\xaf\x73\x28\x9e\xf6\x88\xb5\x28\xc9\x26\x9d\x75\x28\x85\xa2\x85\x5a\xaa\xdd\xb8\x94\x7b\x9e\x4d\xa0\xb5\x62\x87\xd1\x65\x2d\x29\x85\xdb\x22\xae\x82\xda\xb8\x82\x2f\xdd\x42\xfb\x56\x4a\xb5\x07\xab\x7d\x61\x67\x3d\x72\x93\x49\x2d\x99\xa5\x06\x54\x85\x75\xbe\x13\xe4\x77\x84\x7b\x77\x54\x6c\x73\x7d\x54\xe4\xc9\x2b\x32\x12\x1e\xd0\x9c\xe0\x7a\x02\x16\x33\xad\x72\x3b\x8a\x13\x15\xa6\x6c\xb0\xd2\xdc\x24\xd4\xc8\x6d\xab\x9f\x84\xc9\x2d\x28\xcd\xcc\xc7\xde\x3e\x87\xbc\x31\xa1\x26\xc0\xf5\x34\x52\x03\x27\x2b\x84\x1a\x8d\xd4\x79\x9a\x74\x8a\xb8\x9e\xd0\x22\x52\x10\x39\x06\x7e\x73\x46\xd0\x8e\xc6\x50\xeb\xbe\xc1\x13\x6f\xa5\xf3\x63\x66\xa1\x16\xf1\x14\xef\x81\x46\xbd\xfb\xd4\xc1\xc7\xcf\xb3\xdb\x36\x88\x06\xbe\x31\x2d\xb0\xac\x61\x53\xea\x6c\x0f\x77\xfa\x01\x84\x73\x22\xdb\x5b\x56\xe4\xe3\xdd\x0a\xf8\x42\x78\x5b\x6a\x9d\xfb\x29\x58\x7f\xc8\x7e\x3e\x5b\x3f\xd4\x0a\xc6\xe3\x1d\xaa\x3d\x9e\x60\x3c\xf6\xf3\x7e\x70\x22\x65\x6f\x3c\x9f\xd2\xb6\x17\x26\xc1\xa1\xdb\xab\x24\xd6\x49\xa6\xeb\x93\xbf\x4a\x91\x16\xf6\x78\xea\x46\xbc\xac\xa7\x76\xe4\xdf\xbb\x62\x78\x35\x79\xe5\xc7\xaf\xee\xfc\x90\x57\xd7\xaf\xe2\xd3\x76\xaa\x7b\xe3\x4a\x3b\x16\x0d\x85\x33\x33\x07\x93\xb3\x7b\x0b\x69\xdb\xbb\x8a\x21\x3c\xf0\x7d\x1e\xf1\x64\x4e\xb5\xd3\x3b\x23\xea\x42\x66\x90\xc9\x3a\x64\xc4\x75\xe8\x83\xb6\x72\x07\xd2\x61\xe5\x45\xd9\x50\x67\x55\x4b\xce\x61\xc1\xb9\xe2\x9d\x90\x5f\xca\x27\x71\xf5\x2f\xcb\x34\xb9\xf1\xf4\xe0\xa7\x0f\xe3\xdb\x9f\x6e\x21\xfc\x1b\xc2\x4f\xa5\x3e\x6e\xa5\x2d\xe0\x22\x64\xdd\xcb\x76\xe9\x6c\xfe\x30\xbe\x9e\xbe\x0b\xeb\x87\xf4\xbd\x7d\x77\x37\x7f\x18\xcf\xef\xe6\xaf\xfd\xcb\x21\xac\x0d\x65\xb6\xf1\xdd\xfc\x81\xd1\xbf\x47\xb3\x99\xae\x6a\xbe\x42\xd0\x2a\x6a\x88\x60\x04\xc5\x4a\xef\x72\x3b\x40\x5f\xe9\x78\x8a\xd8\xbb\xbc\xa6\x2c\x4d\x59\xa9\x7d\x2f\xd5\x77\xc4\x4c\x13\x3a\x6a\x5c\xfe\x53\xc7\x20\xaf\xc4\x37\x59\x35\x55\xb8\x8f\x66\x6c\xac\x55\xd6\x18\x2a\x47\xe5\x29\x66\x5c\x76\xe4\xb6\xc8\x60\x8b\xc6\xb9\x7e\xa5\xc9\x4d\x25\xbe\x8d\xe3\xeb\x30\xbf\x5b\xb8\x57\xd4\xae\xec\xb4\xce\x41\xe6\x28\x02\x74\x6b\xb2\xb3\x89\x4b\x32\x0c\x78\xfa\x95\xe5\x56\x4d\x96\x48\xa9\x44\x6c\x79\xda\xa4\xa4\x93\xa2\x94\xff\xe4\xde\xdb\x3b\xfc\xef\xe1\x5e\xbd\x39\xaf\x1f\xd4\xe1\xf0\x34\xee\x85\xcb\xf8\x34\x69\x2c\x89\xe6\xe3\x22\x61\xa8\xdd\x7e\x0b\x6a\x88\x13\x69\x5f\xc6\x43\x21\x75\x86\xe7\x71\xe2\xa0\x65\xde\xb6\x8d\x3c\x90\x08\x80\xc8\xa0\xd5\x8d\xc9\x90\x47\x6b\x61\xbc\xce\x2e\x26\x9c\x07\x7d\x1a\x4a\xad\x76\x68\xb8\x96\xfb\xed\x64\xa2\x0d\x72\x1e\xf5\xbd\x08\xa3\x93\x28\x3a\xe7\xa9\x9d\x11\x39\xa6\x49\xe0\x69\x4c\x31\x1b\x3f\x87\xdb\xf6\xfb\xc6\xd5\x0d\x85\x82\x2d\x28\x27\x58\x27\x5c\xd3\x5d\xf5\x1d\x7d\x42\x0b\x56\x6c\x27\x03\xdc\xb4\x38\xd3\x28\xca\xf3\x11\xdb\x1a\x3c\x1a\xe9\x1c\xaa\x90\x50\x2b\xa9\x1a\x87\x69\x12\x48\x06\x35\x8d\xfd\xd7\xb4\xd4\xbb\xe4\x09\xf4\x28\xf5\xae\x2b\x01\xac\x36\xc2\xc7\x1e\x74\xd9\x93\xa5\xd7\x17\x3a\xdc\xc7\xb4\xbf\x6c\x90\x5b\x30\x8d\x52\x0c\xb3\xad\xbf\x35\x3e\xc8\x0c\xf9\x67\x0c\xa7\x1e\x95\x10\xaf\x83\x3f\x56\x86\xe2\xbd\x82\x0f\x54\xd9\xff\x08\xce\xf3\x47\xa9\x77\xbd\x1b\xe8\xcb\xf8\xa3\x14\x3a\x53\x1b\xfa\x6f\x2c\x6a\x82\x75\x0c\x75\x0e\x68\x0c\x95\x13\x8f\x37\x3c\xfb\xb4\x63\xc0\x54\x82\xc1\xbd\x72\x7c\xa5\xd2\x21\x35\x68\xd5\x01\x7b\xb2\x70\x53\x53\x2d\x3d\x16\xf4\x6e\xd0\x1d\x12\x68\x74\x27\x4a\x97\x02\xf3\xa3\x15\x7a\xc4\xdb\x15\xb6\x8b\x4d\xe3\xfc\x10\x41\xbb\x82\x90\x3b\x9d\x16\xff\x45\xbc\x4b\xda\xbe\xe9\x49\x71\xf6\x82\x13\x22\x37\x41\x7d\x60\x06\x25\x1e\xb0\x24\xb7\xa2\x35\x1e\x09\xc1\x01\xcd\x46\x5b\xe9\x4e\x3e\x84\x26\x9c\x56\x65\x89\xca\x8d\x00\xbf\x65\x58\x3b\x86\x3f\x5b\xe1\x44\x09\x68\x8c\x36\x04\xb9\xde\x78\xa4\x24\xac\xf6\x79\xa5\xbb\xc5\x2f\xa1\x21\x73\x27\x43\x78\xcb\x3e\xf4\x03\x47\x24\x17\x2f\x6a\x45\x71\xd3\xec\x7a\x4e\x47\xb5\x70\x53\x62\x45\x34\xdf\x03\x0f\x5a\x9c\xc1\x8a\x30\xa4\x67\x0c\x13\xfa\x1f\x5e\xb3\x48\xc4\x56\x86\x60\xb0\x46\xff\x6b\x81\xe8\x5b\x29\xc0\xcc\xf9\x1f\xbd\x4c\x27\x0c\x38\xff\x6c\x50\x51\x7e\xe8\xdc\x2f\x44\x13\xa3\xc2\xf0\x90\x5b\x12\x87\x3b\x6d\x4e\xed\xa5\xa1\xf6\xd1\x13\x9c\xb4\xd4\x3b\xca\x60\xfe\x92\x39\xf9\xff\x00\x00\x00\xff\xff\xe2\xd0\xd7\x91\x80\x25\x00\x00")

func templateServerConfTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateServerConfTmpl,
		"template/server.conf.tmpl",
	)
}

func templateServerConfTmpl() (*asset, error) {
	bytes, err := templateServerConfTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/server.conf.tmpl", size: 9600, mode: os.FileMode(420), modTime: time.Unix(1504442070, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"template/ccd.file.tmpl": templateCcdFileTmpl,
	"template/client.ovpn.tmpl": templateClientOvpnTmpl,
	"template/dh4096.pem.tmpl": templateDh4096PemTmpl,
	"template/iptables.tmpl": templateIptablesTmpl,
	"template/server.conf.tmpl": templateServerConfTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"template": &bintree{nil, map[string]*bintree{
		"ccd.file.tmpl": &bintree{templateCcdFileTmpl, map[string]*bintree{}},
		"client.ovpn.tmpl": &bintree{templateClientOvpnTmpl, map[string]*bintree{}},
		"dh4096.pem.tmpl": &bintree{templateDh4096PemTmpl, map[string]*bintree{}},
		"iptables.tmpl": &bintree{templateIptablesTmpl, map[string]*bintree{}},
		"server.conf.tmpl": &bintree{templateServerConfTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

