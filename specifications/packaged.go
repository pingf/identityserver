// Code generated by go-bindata.
// sources:
// specifications/api/itsyouonline.raml
// specifications/api/securitySchemes/oauth_2_0.raml
// DO NOT EDIT!

package specifications

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

var _itsyouonlineRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x3d\x7f\x77\xdb\x36\x92\xff\xfb\x53\x60\x73\x7b\xaf\xbb\xaf\xb6\x24\x3b\x4e\x36\xf5\x3f\x59\xc7\x71\x52\x77\x63\x37\x2f\x76\x92\xd7\xe6\x7a\x0d\x45\x42\x16\x6b\x8a\x54\x49\xca\xb6\x9a\xed\x77\xbf\x19\xfc\x20\x01\x10\x00\x49\x49\x76\xd2\xeb\xf2\x35\x4d\x44\x82\xc0\xcc\x60\x30\xbf\x30\x18\xfe\xd7\x7f\xbf\x39\x3c\x7d\x45\x76\x07\xa3\xad\x32\x2e\x13\x7a\x40\x4e\xca\x62\x99\x2d\xb2\x34\x89\x53\xba\x75\x4d\xf3\x22\xce\xd2\x03\x32\x1a\xec\x6e\x8d\x83\x82\xbe\xcd\xe3\x03\x32\x2d\xcb\x79\x71\x30\x1c\xc6\xac\xe9\x40\xb4\x2d\x68\xb8\xc8\xe3\x72\x79\x1e\x4e\xe9\x8c\x16\x07\x5b\x84\x64\xc1\xa2\x9c\xfe\xbc\xf7\xf3\xe8\x80\xfc\x25\x4e\xc3\x64\x11\x51\x62\x34\x1b\x56\x6d\x06\x79\x30\x4b\xb6\xb6\xca\xe5\x9c\xbf\xfc\x2a\x18\xd3\x04\xff\x41\x08\xde\x3b\x20\x45\x99\xc7\xe9\x25\xbb\x31\x0b\x6e\x5f\xd1\xf4\xb2\x9c\x1e\x90\x47\x23\x7e\x27\x4e\xe5\x9d\xbd\x2d\xb8\xf3\xb6\xa0\xf9\xe1\xeb\x93\x7f\xd1\x25\xef\x22\xa2\x45\x98\xc7\xf3\x92\x61\x83\x0f\x49\x31\xa7\x61\x3c\x89\x43\x02\xcd\xc8\x15\x5d\xb2\x66\xf3\x3c\x9b\xd3\xbc\x8c\x39\x08\xfc\x5a\x40\xeb\x34\x98\xe9\x10\xe0\x15\xcc\x63\x78\xcf\x72\x7b\x9e\xc4\x61\x80\x43\xc5\x51\xe3\x69\x11\x66\x88\xa0\xb8\xfd\xe1\xa7\xea\x41\xc2\xf0\xe5\x68\x23\x06\xaf\x17\x63\xe8\xa6\x42\x40\xc7\xa0\x7a\x48\xb2\x09\x09\x18\x88\xbc\x95\x0d\x01\xad\x67\x79\x73\xce\x7a\x50\xe1\x87\x47\x87\x30\x17\x59\x1e\xff\xc6\x80\x3f\x0d\xe6\xb6\xa1\xe1\xf6\x1c\x9a\x93\x31\x2d\x6f\x28\x4d\x49\x4e\x7f\x5d\xd0\xa2\xa4\x11\x1f\xa7\x20\x41\x1a\xc1\xcd\x20\x11\xbf\xdd\x70\x55\x6f\x5a\x01\xc4\x2e\x4c\x9a\x3c\x8f\x2f\xe3\x32\x48\xde\xc3\x13\x5a\x6a\xc0\x5a\x66\xb9\x27\xa0\x03\x72\x98\x14\x19\x99\x06\xf0\x80\x7c\x05\x4c\x9a\xd3\x34\x5c\x7e\x25\x21\x77\xf1\x87\x0f\x89\x06\x0a\xfc\xb6\xec\xbb\xa6\x44\x83\xc1\x0d\x96\xde\xad\x6f\xd6\x9c\xbf\xfb\xa8\x31\x65\x16\x2a\xbc\xc8\x72\x40\x94\xd0\xdb\x79\x12\xa4\xac\x11\x09\xc6\xd9\xa2\x14\x8c\xc8\x88\xc0\xfe\x39\xe3\xe4\xda\x86\x15\x4a\xab\x25\x0e\xe4\x9e\x2e\xc6\x83\x30\x9b\x89\xd5\xce\x17\xfb\x30\x8e\x68\x0a\x02\x63\x09\x7c\x07\x32\x62\x38\x4e\xb2\xf1\x70\x16\x00\x19\xf2\x61\x94\x85\x62\x55\xef\x0d\xf9\x18\x83\x59\xd4\x7f\x71\x5d\xe6\x41\x0a\x54\xbd\xc8\xea\xb6\x0e\x3a\x19\x08\x5f\x4c\x29\xc9\xf2\xcb\x20\x15\x34\x21\xe5\x34\x28\x49\x5c\xc8\x1e\x71\x56\x22\x12\x84\x21\x2d\x0a\x52\x66\xa4\xc8\x66\x14\x57\x51\x09\xef\x21\x34\x5f\x15\x24\x4e\x27\x59\x3e\x63\x6f\xd7\xab\x3a\x8a\x72\x78\x83\x16\x4f\x0f\x5c\x23\x4b\x86\x13\x9d\x09\x2e\x13\xbf\xe4\xe8\x59\x4a\xd9\xb0\x4a\x8b\xd2\x84\xb8\xe2\xa9\x81\x32\x94\x8d\x7a\x0a\x51\xcc\xa5\xab\x48\x96\xb9\x14\x16\x00\xbb\xaf\x1d\xb0\x40\x9c\xfc\x91\xf1\x9c\xc2\xa0\xe9\x62\x36\x06\xad\xf5\xc7\x83\x7e\x1c\xa4\x57\xc0\x95\xd9\x22\x2d\xff\x80\xd0\x4f\x82\x90\x8e\xb3\xec\x0a\x38\x0c\xfe\x4a\x68\x50\xaf\x1c\x2e\x43\x2c\x0f\x54\x78\x0a\xd7\x32\x57\xc6\x30\x48\xf1\x2a\x2e\x4a\xc4\x5c\xeb\x86\x21\x2a\x70\x63\x74\x52\x71\x0e\x41\x10\xa2\x74\x03\x39\x96\x93\x19\x65\x8c\x32\x8d\xe7\xd0\xc9\x00\x45\xe9\x33\x98\x81\x43\x3e\x03\x07\x4e\x81\x15\xc3\x3c\x75\x10\x49\x8a\x9c\x7e\x38\xaa\xe7\x38\x0e\xfb\xbd\xbb\x5b\xbf\xcb\xe0\xca\x97\xfd\xde\xdf\x1f\x19\x56\xc0\x56\x63\x52\x6b\xed\x2f\xb4\xeb\x21\xac\xfe\xf2\x90\x8b\x01\x17\x1d\x84\x94\x30\x06\x17\xba\x4e\xbb\x07\x8a\x27\xce\x61\x98\x28\x28\x69\x19\xcf\xa8\xa1\x03\x8b\xe5\x6c\x9c\xd5\xef\x30\xc3\xa7\x5e\xc6\x07\x4e\xe3\x61\xde\x68\xe4\x26\xc8\x3c\x28\x41\x31\x01\xc7\xfc\xef\xff\x7c\xfd\xf4\xc3\x68\xe7\x9b\x9f\xbe\xfe\xeb\x96\x83\x2a\x26\x4d\x5a\xe8\x00\x88\x80\x16\xd4\x17\x8b\x63\x56\x5c\x3c\x01\x0d\x29\x2d\xfb\x77\xf1\xa8\xee\x22\xcd\xfb\xbf\xae\x70\x56\x06\x6b\x26\x7f\xba\x16\x12\x16\xe6\xec\xd6\x87\xc2\xa0\xf3\xac\x00\xee\x0b\xb3\x88\xf6\xef\x66\xaf\x33\x9f\x13\xf2\x42\x88\xaa\x96\xa5\x0e\xb6\xbb\x61\x9f\xc5\x20\x5c\x2f\xa5\xa1\x4d\x98\xd5\xe2\xb3\xe0\xe6\x71\x58\x2e\x72\x7f\x1b\x30\xa6\xae\x1c\x0d\xe0\xee\x4b\x26\x39\x5b\xe0\x4c\xb2\xcb\x38\xf5\x0e\xd2\x86\x49\x70\x1d\x94\x41\xfe\xf3\x22\x4f\xbc\xdd\x4c\xcb\x59\xd2\xda\xc8\x43\x15\xb8\x7b\x8c\x16\x46\xcb\x92\xf2\xae\x48\x7e\x4b\x35\x54\x3c\x6b\x17\xdd\x3c\xf7\xc2\xad\x2c\xcf\x75\xf8\x9e\x75\x5e\x89\x97\x0f\xc1\xce\x6f\x28\x5e\x3e\xed\x6d\x3f\x1c\xfd\xfe\x57\xfd\x7d\xd5\x41\x95\x37\x27\x71\x5e\x94\x56\xf3\x37\x09\x1c\x0f\x6a\x63\xce\xa2\x21\xb9\xb0\x7d\xda\x90\xb6\xa0\x39\xb9\x6c\xbf\x61\x9e\xd3\x81\x4d\xd4\x3b\x0d\xc1\x03\x6d\xda\x1c\xf6\xd6\x81\x2a\xb6\x95\x36\x4a\x2f\xcd\x0e\x54\x93\xe7\x40\x55\xbf\x56\xc3\xc2\x32\x4f\xc6\x52\x6e\x98\x1c\x96\x57\xb4\x35\xb5\xb5\xc5\xa9\x16\xcc\xe6\x09\xb5\x39\x25\xe3\x6c\x6c\x9b\xac\x67\xca\xed\x7a\xa6\xbe\xcb\xa6\x69\xa1\xb8\x0c\xca\x54\x69\x80\xec\x90\xdd\x97\xef\xce\xff\xf1\x68\x7f\xbf\xbc\x7e\x7f\xba\x77\x71\xba\xfb\xf0\xfa\xec\xdb\x27\xc5\x38\x0b\x66\xdf\xfd\xf6\xea\xf1\xec\xbb\x87\x8f\x8d\x29\x45\x09\xb7\xfb\x64\x67\x77\xb4\xb3\x37\xba\xd8\x7d\x7c\xb0\xbf\x0b\xff\x0d\x46\xdf\x8c\x7e\x74\x4d\x9a\x31\xa2\x50\xcd\x37\x59\x7e\xa5\x5b\x76\xfa\x62\x42\x84\xff\x09\xde\xde\x3c\x48\x97\xe8\xf5\xd9\x7b\x99\x66\x0a\x67\xb9\x7a\x11\x54\xd5\x7a\xd1\xb9\x70\xcb\xd6\xf7\x38\x1e\x87\x59\x9c\xea\xdd\x57\x1d\x4f\xa2\xc9\xe5\x04\x24\x17\xff\x4f\x6f\xd4\x99\x52\xec\x32\x0d\x90\xf1\xd8\xce\xd7\x76\x0a\x8c\x4d\xfc\x55\x7b\x84\x3c\xf8\xfa\xe1\xde\xee\xde\xc3\x7d\xf1\xe7\x81\xbd\x0f\x7a\xe9\xef\x03\xf0\xb0\xf5\xd1\x36\xc5\x96\xc9\x61\x36\x0a\x39\x9f\xa3\xb4\x98\xc4\x34\x89\x8c\xe7\xc2\x04\x01\xe7\x22\x4e\xc5\x0f\xa3\x05\x58\x18\x64\xff\xd1\x33\xb3\x5f\xa1\xf5\xc9\x45\x36\xcb\xf2\x3c\xbb\x49\x82\xd4\xec\x5b\xd1\xea\xe4\x62\xe7\x9b\xfd\x47\xa3\xce\x9c\xc9\xc1\x3e\x0a\xe6\xc8\x33\xe4\x28\x16\x31\x98\x26\xdc\xb2\x89\x13\xf4\x07\xbb\x0f\xd6\x06\x7d\x77\x34\x1a\xd9\x25\x97\x1d\xa1\xab\x71\x68\xf4\xc7\x5c\x07\x72\xf1\x0a\xa7\xf4\x51\xfd\x3f\xa3\x15\xfa\x08\xe4\xf0\xd9\xd1\xf3\xe3\x17\x9d\x60\xde\xe2\xf2\xea\x3a\xa6\x37\x1b\x51\x76\x5a\xb4\xa0\x8a\x33\xfe\x29\xb5\x83\x7c\xa6\x3b\xaa\xb5\xda\x85\x06\xdf\x81\xac\xfa\x5e\x79\x7c\x92\x5e\x03\x2b\xd6\xb1\x38\xdb\x64\xa8\xdd\x35\xc8\x8f\x33\xd5\xb8\x99\x67\x49\x57\x3b\x05\x89\x7c\x40\x3e\x64\x37\x29\xcd\xb7\x85\x93\x5b\xd3\x2b\xcc\x29\x58\x06\x91\x6a\x21\xd8\x54\xa0\x0e\xe1\x6c\x19\x82\xe7\x5e\x64\x61\x48\xf3\x30\x59\x48\x41\xc9\x21\xad\x35\x24\x03\x92\xb0\x81\xb7\xb4\xd1\x98\x48\x7e\xbc\x33\xda\xdb\xd9\x7b\x62\x88\x64\x68\x78\x94\x01\x53\x07\x61\x79\x1e\x5f\xa6\x80\xcc\x1b\xee\xbb\x7b\x1c\x2e\xd1\xfe\xa4\x19\x5b\x9f\x07\x79\xb9\xd4\x7c\x49\x75\x6a\x5c\x5d\x5e\x26\xd9\x38\x48\x9a\x46\xb2\x3b\x1c\xfb\xb0\xbe\xa9\x86\x63\x15\x57\xa6\x69\x12\x3e\xdc\x86\x06\x95\x4d\x68\x33\x0d\x1c\x51\x0f\x18\xe2\xa4\xa4\xb3\x42\x71\x72\xa2\xb4\xcf\x6b\xbb\x8a\xdc\x6a\x0d\xa0\x90\xe7\x67\xe7\x32\x22\xc4\xa6\x72\x25\x00\x1d\x03\x21\xc7\x14\xe4\xa3\x14\x46\xc5\x47\x39\x92\x88\xc5\xf4\x1a\xaa\x0d\x2b\x3e\xd8\xcd\x34\x23\x41\x4e\xc5\x08\xe0\xf8\x90\x72\x1a\x17\x1a\xca\x12\x08\xb1\x43\xb5\x11\xd2\x72\x9e\x3a\x89\x10\x90\x62\x31\xd6\x25\x88\x75\xc9\x55\x5c\x48\x2e\x41\x89\xa5\x71\x89\x37\xa8\x87\x5d\xc0\x8a\x3c\xfc\xe6\xfd\xfb\xe9\xe3\xf8\xf0\x78\xff\xcd\x8f\x2f\xcf\xfe\x11\xdc\x2e\x1f\xdf\xfe\x38\x5e\xbc\x5f\xbc\xba\x4a\x7f\x7d\xff\x2a\x7f\x67\xe3\x97\x1d\x6d\x00\xc5\x3c\x33\xa7\x7b\x87\x91\x70\xd7\xf8\xbd\x67\xfc\x7e\xe8\x9a\x42\xfe\x78\xdf\x49\xdc\x1d\x90\xfe\xb0\xa6\x06\x0d\x68\x8c\x65\x7b\x01\xcf\x91\xdc\x6b\x2d\xdf\x70\x1a\x27\xa0\x5f\x60\x66\x6c\x5d\x73\x51\x7e\x4a\xeb\x68\x52\x73\x98\xa6\x06\xb5\x4a\x09\x73\x9f\x31\x02\x0e\xa4\x29\x09\x8a\x82\x8b\x37\x12\x48\x56\x2c\x33\xdc\x97\xd1\x19\xd1\xc6\x18\xba\x33\x62\x52\x47\xdd\xe3\xec\xb1\xfb\x17\x82\x05\x3e\x0e\xc2\xab\xb7\x6f\x5e\x3d\xed\x17\x51\xdc\xd3\x0c\x38\x45\x28\xd6\x76\x4c\x98\xc4\x34\x2d\x8f\x72\xca\x36\x89\x82\xa4\x78\x89\x71\xe9\x0b\xe8\xda\x1d\xd7\x3e\x49\x23\xdc\x33\xa5\x05\x89\x27\x7c\x89\x5e\xd1\x25\x8c\xbb\x04\x53\x1b\x49\x10\xe1\xd2\x0d\x44\xd7\xa8\x58\x64\xdf\x7c\x93\x79\x8f\x4c\x92\xec\x66\xd0\xc0\xc4\x0c\x39\xe3\xa8\x93\x60\x91\x80\xd1\x38\x81\x97\x6b\x43\xb9\xa0\xd0\x67\xb9\x02\x2d\x98\x06\x63\xfe\x92\x5b\x65\x35\x59\xd4\xd3\xbf\x43\xd3\xb0\xbe\xdb\xd4\x4a\x03\x42\x4d\x29\x39\xdc\x51\x67\xb4\xdd\x21\xda\x85\x9f\x65\xe9\xa3\x11\x75\xc0\x0b\x83\xf1\xd0\x3e\x32\x29\xee\x8c\xfe\x77\x84\x48\x95\xbd\xbe\x71\x70\x33\x6f\x3d\x84\x7d\xbd\x97\xc1\x6d\xe7\xe8\xab\xad\x8f\x46\xec\xa1\xd6\x03\x71\x12\x5f\xd1\x10\x4c\xfc\x65\xeb\x14\x82\x2e\x78\x72\xfe\xf2\xdb\x1f\xce\xe3\x27\xdf\x1d\x5d\x1f\x5e\x5f\xbc\x7d\xfc\xc3\xec\x7a\xf7\xdb\x37\xb3\xd9\xe2\xc9\xe3\xd3\x5f\xf7\x7e\x98\xff\x66\xce\x9f\xdb\x28\x73\x51\x6e\x87\xdd\xf9\x67\x0d\x97\xe6\xdb\x73\x4a\x90\x67\xc7\x8a\xd7\xca\xfe\xe0\x12\x11\x21\x85\xda\x49\xa9\x09\x25\x56\x0f\x93\x6c\x68\x03\x06\x75\xd0\xb4\xb9\x96\x50\x8a\xd2\xe8\x99\x99\x74\x81\x8c\xd7\x60\xbf\xb9\x91\x40\xe1\x9a\x9a\xc6\xee\x31\x7f\x91\x09\x20\x94\x3c\xca\xd0\x81\x27\xa0\x0b\x77\x5f\x33\xfb\xd3\x01\xba\x2b\x58\xaa\x9a\xab\xd2\x16\x76\xf5\x81\x06\x6e\xdc\xb4\x50\xd8\xb8\xaa\x81\x12\xa7\x92\x8b\xfd\x56\xda\x22\x8d\xc1\xda\x16\xf7\xcb\x7c\x51\xed\xce\x00\x1c\x20\x5f\x4d\x25\x2a\xa0\x43\x49\xde\x66\x2c\x5b\x23\xfc\x7a\x62\x02\x28\x5f\xc2\xa2\x11\xa8\x0a\xa5\x94\x9f\x64\xa0\x19\x83\xcb\x82\x8c\x97\x6c\x2b\x8f\x99\xf4\xdc\x91\xe0\xda\x53\x02\x51\xd9\x15\xd7\xb8\x5a\x50\x73\x3c\xed\x65\xb8\xd9\x81\x4a\x84\xf1\x58\x3b\x19\x5f\x15\x5c\x1d\xc9\x3b\xea\x88\x52\xdf\xd0\x5b\xa0\x56\x74\x6f\x00\x14\x3c\xa1\x03\xc6\xc4\xe4\x2c\xb2\x98\x57\x31\x47\xbe\xb4\x8b\xe6\x6e\x5b\xed\x33\xf5\x58\x0b\xff\x56\xd6\x3e\xae\x8b\xba\x17\x84\x01\x67\x67\x1a\x14\x53\xb9\xf9\x2c\x78\x73\x9b\xfd\x10\x0c\x54\xff\x90\x7c\xb3\x4d\x58\xc6\x09\xde\x66\xc0\x72\xa7\x03\xc1\x55\x95\xf7\x45\x86\x16\x4a\xb8\x48\xe0\x7e\x35\x10\x74\x16\x5c\xf1\x9f\xbf\x14\x6c\x07\x7b\x0e\xb8\xc2\x28\xbc\x0f\x01\x46\x45\xa6\x1b\x70\xe5\x31\xcf\x05\x6f\xc6\x11\xcf\x73\x91\x4b\x18\xa1\xcc\xe3\x19\x7b\x96\xd2\x1b\x4c\x69\x41\x9f\x80\x14\xf3\x20\x84\x7f\xfd\x8d\xde\x86\x74\x5e\x72\xd3\x0d\xda\x2c\x31\xa7\x8b\x06\xc2\x7d\xa0\x04\x18\x00\x7c\xd4\xbf\xb3\x3e\x69\x8a\x01\x22\x7c\x42\x16\xe5\xe4\x89\x8a\xc5\x8b\x3c\x9b\xf1\xb9\x83\x11\xc1\xe6\x10\x08\x04\xe4\xfc\xdb\xc3\xbd\x47\x8f\x81\xd7\x13\x30\x5a\x80\xe9\x81\xd7\x03\xf2\xe6\xe4\xf5\xf1\xe9\xf3\xdd\xc7\xa3\x6d\x00\xe9\xb6\x24\x98\x6e\xf7\xe8\xc9\xd1\x94\x86\x57\xc7\x7c\x0c\xbe\xb5\x8d\x3d\x31\xe4\xe0\x9d\xd1\xed\x68\x04\x62\x82\x4e\xe2\x5b\x61\x3c\x2a\x72\x0a\xd8\xa0\x92\xa7\x8c\x07\x8f\xf3\x3c\x73\xda\xb7\x94\x3d\x54\x85\xd1\x1b\x7a\x09\xcc\x98\x2f\x8f\xeb\x9d\xbd\xe6\x6b\xed\xe2\xb5\x2d\x8d\x09\x28\x21\x6e\xbf\x43\xb2\x76\x17\x2c\xbb\xa3\x3d\xd0\x2d\x2c\x9f\x90\xab\x84\x0f\x75\xc2\x21\xf9\x69\x6b\xc8\x1c\x40\xec\x0e\xe3\x78\xbc\xdb\x28\x2e\xe6\x49\xb0\x3c\x63\x72\xf7\x88\x45\x27\xde\xca\xd4\x39\x8d\xf5\xf9\x33\xa0\x30\x70\x47\x9d\x5d\x37\xce\x22\x05\x5b\x25\xcf\x6f\x88\x0c\x69\xb3\x02\x58\xef\x70\x7f\xf8\x49\x9a\xf1\xbf\xf3\x66\x97\x75\xfc\x5b\xc5\x40\xc9\x98\xfc\x54\xa5\x0a\x7e\x20\x0f\x58\xac\x85\xf9\x4b\x0f\xc8\x4f\xe4\x77\x22\x45\x8a\x86\xd0\x4b\x5a\x56\xd8\xa0\xb5\x51\x80\x60\x68\xc4\x89\xc1\x71\xd6\x6f\x34\x11\x53\x2f\x3f\x92\x0e\x84\xf1\xe7\x50\x75\x96\xd6\x41\x71\xbe\x50\x76\x0a\x34\x7c\xdf\xce\x23\x31\x81\xf8\xdb\xe1\x82\xb1\x36\x55\x72\x17\xc1\xd5\x58\x6d\xe4\xb0\x05\x2c\xf7\x6f\xb6\x5c\xc4\xf0\x93\xc0\x99\x32\xe3\xd9\xdc\xe3\x97\x73\x8b\xcf\x3a\x73\x7b\xa3\x7d\x7d\x00\x0d\xcd\xf3\x05\xcb\x67\x9b\x2c\x92\x04\x8c\x16\x86\x73\xd4\x86\x27\x01\xd5\xec\xeb\x93\x65\xc9\xa6\x59\x09\x72\x6a\xc1\xe3\xcc\x30\xab\x73\xf0\x64\x6f\xb2\x3c\xda\xf8\xcc\x5a\xe7\xf6\xb5\x18\xcd\xe5\x35\xda\x66\x77\xde\x7c\xa7\xc9\xde\x6d\x6c\xed\x9e\x55\xb9\x73\x54\x56\x94\xb0\xcf\x2e\x41\xd1\xe1\x69\xe3\x58\x9e\x96\x89\x6e\x60\x7d\x96\x55\x6f\xa3\xe6\x0c\x8c\xd6\xfb\x7b\x7b\xb6\x55\x6a\xb8\xdb\xcc\x88\xa9\x88\x45\x78\xc2\x23\xea\x40\x66\xdf\xe0\x33\x03\x4f\x4b\x97\xeb\x8b\x0d\x2e\x34\x98\x62\xe2\xfc\x65\xdf\xc5\x5c\x87\xcb\x2e\xa9\x4b\x7e\x80\xbc\x54\xf7\x2e\x68\x61\x17\x21\xd0\x0c\x14\x81\x34\xcc\x34\x5e\x63\xd0\xd6\xbb\x18\xb5\xf2\x07\xab\x3a\x5f\x82\x59\x0e\xe3\x94\xc6\x2e\xa2\xb4\x1e\x23\xeb\x06\x44\x33\x60\x61\x40\x93\xb1\xbf\x83\x04\xed\x2d\xde\x3b\x9a\xcf\x93\x38\xc1\x7f\xa1\x19\x94\x82\x04\xa8\xc6\xd0\x21\x54\x51\x74\x08\x19\x43\x39\xd8\x66\xb8\x7d\x6e\xc5\xac\xda\xf6\x85\x6a\x5d\x6c\x99\x10\x6e\x71\x80\x38\xa7\x37\xea\xcb\xae\xd5\x2f\x9b\x0b\x35\xad\xa1\x7a\x67\xab\xdf\x1a\x4e\x93\x97\xbe\x0b\xde\x67\xcd\xef\xb6\xac\x79\x89\x2b\x4c\xa9\x1f\x5b\x17\xce\xdd\x30\xef\x42\x81\x6e\xa4\xa8\xae\x16\x9a\xe0\xb5\x3f\xfa\xa6\x55\x62\xb1\x51\x98\xef\x93\x60\x02\x37\x77\xcd\xe5\x82\x1b\x7e\x62\xb0\xfc\x5e\xf7\x62\xe8\x16\xab\x76\x71\xf0\x98\x55\xc7\x14\x75\x32\x2d\x6a\xd3\x21\x08\x4a\xe6\x05\xb0\xf3\x17\xa9\x67\x36\x56\x5b\x41\x7e\xfa\xb7\x10\xbe\x95\xe2\xfd\xf8\xd0\x4a\x0f\x53\x1d\x74\x99\xc2\x0b\xee\x6d\x09\x2a\x1a\x53\xa9\xc8\xde\x04\xa4\x9a\x67\xf2\x9e\xb3\x06\xdd\x26\xef\x0d\x9d\x65\xd7\xb4\xf0\x4f\x91\x93\x1c\x0e\x65\x6c\x8c\x71\xac\x76\x0c\x9d\xe1\x88\x5a\x56\x75\x3f\x0a\xa1\x9d\xa6\x03\xcb\xf2\x96\xd1\x10\x1b\xd3\xba\xfb\xaa\xab\xa1\x14\xf6\x6a\xef\xa6\x98\x35\x48\xf8\x4e\xbc\xd2\x8d\x88\xe7\x18\xdc\x90\x3a\x05\x5d\x6d\x0e\x1d\xa8\x9d\x15\x68\x6a\xa3\xa8\x36\xda\x3b\x73\x1c\x3c\x91\x33\xc6\x03\x3c\xe0\xeb\x4b\x33\x94\x1f\xbd\xda\x88\x7d\xe0\xd5\x48\x40\x9b\xc3\x79\xfc\x2f\xba\x74\xe9\x20\x68\xc0\xb8\x8b\xef\xc3\xc8\xa4\xfb\xfa\x3c\x16\xbb\xee\x45\x03\xad\xa1\x5c\x00\x09\xa6\x57\xbc\x48\xb8\x50\xe9\x86\x90\xbc\x6a\x2f\x91\x0f\xb6\x69\x4d\x70\x49\xdd\x93\x89\x3b\xc3\x7c\x54\xa7\x49\x81\x4d\xb8\xa8\x17\xa7\x02\x8b\x6e\x14\xb6\xf8\xd4\xd6\x4d\x69\x4b\xaf\xfc\xda\x34\x5d\x2b\x83\x6b\x35\xd5\x08\x9d\x5c\x19\x93\xd3\xa2\x14\x59\xe0\x16\x7e\xc1\xca\x24\xe6\xab\xf7\xaa\x03\xbf\x68\xe5\x76\x69\x66\xfa\x9b\x0e\x49\x1b\xe1\x99\x33\x92\x4a\x3e\xc2\xc8\x61\xd2\x87\x02\x16\x3e\x35\x65\x81\x72\x1a\x56\xbf\x36\xe5\xed\x29\x8b\xbf\xaf\xd2\x6f\xa3\x8e\xa2\xee\x6d\x78\xac\xa9\xe8\x25\xd1\x6b\x1d\xcc\x95\x51\x2e\x02\xa5\x9b\xd7\x46\xa6\x2e\x42\xd2\x69\x61\x59\xbb\xc7\xca\x95\x92\x1e\xc0\x55\xa5\xfa\x57\x68\xa6\xf0\x67\xdb\x7c\xbf\x9d\x32\xc4\x0c\xae\x85\x67\x6c\xb7\x01\x90\xcd\x6f\xf2\xb8\x2c\x69\x5a\x9b\x35\xfd\xa2\x64\x7c\xe6\xed\xa0\x3b\x03\x95\x8d\xe5\x6a\x51\x5a\x08\x7a\x17\x3c\x8d\x9e\xec\xbc\xdc\x8d\x8f\x7d\xb8\xe0\x65\x95\x16\x2b\x6a\xb0\x86\xfe\x52\x39\xc0\x3e\xf9\xb5\x0a\x93\xed\x08\x05\x18\xe3\x36\xdf\xdf\x2a\x1e\xac\x4a\xcc\xd9\xef\x9d\xd0\x56\x68\xb2\x6a\x95\x0d\x3f\x01\xab\x2a\xea\xac\x29\x3a\xd6\x59\x7f\x0d\xa2\x73\xc1\xe3\x5e\x78\x2e\x09\x64\x70\xe5\x84\xef\x05\xf9\xf8\x72\x8d\x88\xa0\x3e\x56\xd3\x47\xf0\x59\x45\x62\xd7\xa0\x13\x76\x3c\x14\xd6\x0d\xb3\x41\x3b\x6a\xed\x76\x53\xc5\x6a\xcd\x65\x86\xd7\x7a\x56\x53\xfb\x42\x6e\x8f\xc4\x56\xd1\x71\xce\xa4\x3c\xb5\x58\xbe\x65\xb2\xa6\x85\xb1\xec\x09\xc7\xba\x4d\x82\x47\xee\xaf\x44\x6a\x32\x31\x9b\x3a\x49\xdb\x80\x5d\x97\x9e\xbc\x1b\xb2\x60\x9d\x53\xe9\x57\xc9\x7c\x6a\x17\x06\xeb\x2d\x2d\x0b\xfe\xae\x1c\x6d\x0d\x58\xde\x94\x9b\x97\x45\x91\x85\x31\xdf\xde\x10\xaf\x92\x60\x65\x92\x58\xc7\xb1\xf4\xcc\x89\x53\xde\x64\x13\x0c\xb9\x4e\xb3\xa8\xce\x3e\xf4\x45\x97\x2f\x6e\xb2\x17\x87\xa7\xfc\x05\x3b\x66\xb8\xa0\xd8\xe6\x79\x56\x14\xf1\x38\x01\x1c\x6f\x32\x1c\xbe\xc4\xaa\x0d\x40\x5e\x4c\x72\xe3\x7c\x4c\x66\x46\x3f\x77\x19\xc0\xf5\x9b\xde\x65\x56\xce\xed\x91\x6a\x7e\x15\x33\x6b\xc2\xff\x90\xbd\xd7\x89\x6e\xdf\x5f\xbc\x3e\x67\x09\x79\x5f\x08\xb6\x3c\x3b\xd0\x3c\xdf\xe9\x36\xd5\xce\x69\xb9\x98\x23\x16\x2b\x9a\x4b\x5d\x60\xe2\xe7\x52\xba\x6c\x1d\x36\x36\x84\xac\x5b\x41\xb2\x4f\x8d\xbe\xbe\xfd\x41\xc4\x0f\x04\x02\x60\x4a\x0a\x65\xfb\xb1\x8b\xec\xe3\x5a\x52\x23\x90\x1d\x72\xd3\x94\xd2\x37\xe9\x83\x14\xa3\x62\x5c\xdd\x71\x70\x8c\x35\x33\xa6\x61\x00\xe2\x88\x27\x61\xc4\xd2\x5d\x05\x2b\x26\xb8\xc6\x10\x29\x2e\x38\x76\x7c\x56\x2c\xae\x9e\xa8\xab\x7b\xae\x42\xe7\x6e\x69\x5b\xfd\xc3\x3a\x83\x6e\x95\xfd\x7e\x7c\xfb\xc1\x76\xbf\xcd\xff\x93\x46\xe9\x92\xcf\x93\x07\x20\x8f\x24\x19\x04\xd1\x63\x95\x0a\x51\xac\xe1\xc9\xb7\xe2\x3d\xd1\xc6\xb3\xa9\x86\x4b\xe2\x8a\x3a\xb7\xa2\xea\x14\x3d\x15\x18\x63\x9b\x71\x75\xdd\xe6\xc2\x43\xcc\x89\xb9\xbf\x78\xdf\x33\x22\xab\x77\x68\x5b\x71\xaa\xf4\x72\x6d\xc2\x29\xc0\x57\xab\xda\xbd\xff\xa6\x87\x81\xfb\xe4\xca\xe8\x00\xba\xa9\x64\xb8\x84\xfe\x0d\xb2\x4d\x6c\x04\x5a\x01\x6b\x8a\xa5\xae\xce\x9d\x19\x8b\xf3\x2a\x41\x85\xf8\xcf\x96\x66\x69\xac\xfb\x64\x20\x1b\x0d\xda\x53\x37\x9a\xbc\xd3\xa0\x94\x92\xbd\xe1\xde\x58\xbb\x8d\x79\x21\x17\x31\x9d\x83\xbb\x98\xcf\x3e\x31\x88\x0e\x01\xc3\xf6\x00\x40\x97\x60\x61\xd3\xaf\x75\xf8\xa5\xad\xbb\xe4\x55\x4c\xac\xb9\x24\xd6\xf0\x39\x0f\xcd\x0d\x2f\x53\xbc\x5a\x0e\x7b\xdf\x89\x88\xd5\x8a\xc2\x7d\x5e\x11\xeb\x2e\xab\xd0\x45\xdc\x5a\xde\xee\x20\x76\x05\x99\xd1\x33\x42\xef\x7c\x55\x21\xec\x44\xe0\x4b\x12\xc7\xad\x50\xde\x93\x6c\xb6\x00\xf0\x25\xc8\x68\x37\x5d\xda\xe5\xb5\x8f\xa6\x5f\x8a\xdc\xf6\xc3\xf8\x47\x95\xe1\x3d\x28\x7f\x7f\xb2\x5c\x17\xe5\xcd\x8a\x18\x77\x66\x2c\xb3\xb8\xc1\x19\x1f\xaa\xcd\xee\xaf\xd2\xdc\x54\xf4\xba\xd4\x88\xac\x92\xe8\xf4\xae\xbd\x99\x74\x2a\x0d\x7c\x32\x51\x5b\xdf\xab\xc8\xb8\x0a\x87\x66\x08\xa5\xab\xd1\xae\xbc\xd9\x41\x83\xcc\x1b\xad\xfb\x27\xba\x37\x47\x74\x44\x6c\x76\x3d\x6e\x75\x43\x69\x30\xc8\x88\xd6\xad\x8b\xae\x9d\x28\xeb\x82\x15\x2f\xcb\x8a\xbf\x33\xbb\x5e\x81\xe0\xf3\xea\x0d\x37\x49\xba\xa4\xdd\xdb\x08\xb9\xba\xa6\x50\xd8\x70\xd5\x6d\x46\xdf\x0c\x77\xe3\xc7\x56\xd5\xd0\xc6\x26\xab\xa8\x05\x87\x61\xdf\x4a\xde\x7a\xb3\xa9\xb9\x84\x57\xde\x19\x78\xad\xac\x39\x33\x9b\x6d\x18\x84\x65\x7c\xad\xe5\x9a\x79\xd3\xa7\x64\x20\xc7\xbe\xe2\xfc\x49\x66\x25\x1e\x82\x02\x81\xac\xca\x80\x4e\xea\xde\xbf\x89\xa9\xa4\x97\x09\x1c\x59\x5a\x99\xfe\xc6\x3a\x5b\x98\x5d\x52\x67\x6b\x34\x6d\x05\xc4\x7d\xc6\xda\x3b\x9a\xc7\x93\xa5\xa2\x24\x5d\xf4\x64\x0d\xe3\x9a\x39\x9a\x52\x74\xb3\x59\x69\xc5\xac\xb0\x46\xa4\xbb\xa1\xbc\x19\x6e\xbd\xe6\x38\x1b\x96\x5c\xf3\x38\x84\xd6\x83\x38\xcf\xa9\x03\x88\x07\x21\x8c\x60\xb8\x9d\x64\x5d\xc8\x86\x97\x7a\xd0\xc1\x30\xae\xb0\xce\xd1\xda\x56\x95\xd3\x32\xa8\x4f\xb9\x29\x15\x94\x6c\x46\x81\x38\xf2\x86\xb2\x0b\x21\x32\x76\xd5\xfa\x58\x04\x3e\x36\xe1\x74\x68\x82\xd2\xee\x5e\xae\x62\x4a\xb5\xad\x45\x3b\x34\xad\xea\x5b\x79\x61\x95\xa4\x8a\x8d\x86\x1c\xcc\xba\x58\x0d\xce\xd2\x8d\x93\x8d\xd8\xed\xad\x54\xf9\x32\x8c\x1a\x73\x5e\xbb\x98\x34\x4d\xce\xb4\x9b\x34\xaa\xe1\xa2\x14\x2a\x13\x27\xed\xf0\x73\x02\x4e\x13\xa6\x7b\x9c\xc3\x04\xa6\xa3\xdd\xd0\x8a\x83\xd8\x4f\x0f\x2c\xeb\xd0\xb9\x11\xe7\xdb\x0e\x53\xba\x51\x0e\x0c\xd6\x8d\xfd\x9b\x69\xea\xdb\x1c\x43\x73\x1b\x0d\xfa\xc4\x2f\x75\xa8\xa5\x42\xee\x88\x8f\xcf\xd4\x91\xec\xc4\x93\x5b\xf6\xf2\x40\x98\x06\x1c\x3b\xf0\x5e\x50\x56\x72\x6a\x0e\x76\x05\xb2\x47\x5c\x95\x84\x63\x67\xcc\x61\xf2\xf3\x0c\x94\xcd\xe7\x93\x1c\x9d\x0e\xf7\x10\xa2\x02\x7e\xe0\x29\x74\x67\x94\x50\xb1\xc0\xc5\xf1\x5d\xa3\x0f\x79\xa8\x5f\x14\x87\x2b\x0e\x1c\x55\xe3\x78\x21\x27\x95\x75\xe4\x9b\x77\x1d\xb6\x90\xf0\x58\x23\xb6\x92\x65\x2a\x60\xb0\xb2\x40\xae\x9c\x5b\x05\x37\x61\xd7\xa8\xa3\x30\x20\xdf\xe7\x11\x3c\xc2\x8e\x04\x23\x8d\x97\x5a\xa1\x04\x67\x58\x44\xd4\xda\x3a\x66\xc5\x20\xac\x45\x87\x5a\x4f\x19\x9e\x88\xaf\xf5\x54\x75\x1a\xc0\x25\xaf\xa0\xdf\x66\x90\xf0\x3a\x4a\x3c\x4a\xc2\x32\x73\xd0\x35\x50\x71\xc4\x35\x90\xd3\x72\x91\xa7\xe6\x71\x18\x57\x31\x9d\xa2\x04\xdc\x6d\xf0\xea\xf5\xaa\x2d\xf0\x9e\xe3\x9b\x40\xc1\x49\x41\xcb\x6d\xa4\xe9\x64\xc1\x53\xd2\xe7\xc1\x25\x90\x6e\x00\x22\x8f\x83\x0b\x94\xfe\x38\xfa\xd8\x0d\x9c\x59\x70\xbb\x12\x30\xa7\xc1\x2d\x8e\x4b\x49\x11\xff\x46\x5b\xa1\x79\xd4\x15\x1c\x01\x52\x3c\xc3\x02\x8f\xa2\xf4\x94\x37\x0b\xa5\x36\xfd\x8e\xf4\xa2\x29\x76\xcb\x8f\x87\x5c\xe4\xfc\xad\xa3\xbf\x64\x61\x1f\x63\xd0\x3e\x01\xda\xcd\x59\x01\x0d\x30\xf0\xda\x6f\x0f\x0a\xa7\x81\xf8\xda\x05\x8d\x1a\x5b\x5b\x81\xfa\x21\x8c\xbb\xd5\x48\x87\x49\xa2\x7d\x76\xa3\xa3\x56\xd2\x21\x1c\x7c\x66\x4b\x4c\xc3\x40\xa6\x83\x7d\xaa\x3e\xf1\xd3\x2d\x70\xa6\x75\xe2\xa7\x82\x86\x3d\x5b\x79\x41\xfd\xbd\x2f\x5b\xed\xc7\x2f\x82\x30\x95\xea\x68\xb5\xf2\x3a\x90\x82\xc7\x86\x5c\xd4\xd0\x0b\x0f\x6e\xeb\x5f\x8f\x51\xcb\x81\xde\xc4\x49\x02\x26\x0e\x49\xb2\xf4\x12\x0f\xbd\x07\xd7\xb4\xfa\x76\x92\xf5\x83\x49\x92\xa2\x6d\xc6\x76\x07\x14\x4e\xb3\x28\x9e\x2c\x41\x57\xc6\xe1\x54\x1d\xc2\x04\x9f\x45\xd9\x58\xf2\x64\x86\x9f\x77\xd9\x84\xe5\xad\x43\x67\xac\x7e\x4b\xe5\xba\x3b\xb3\x2b\xbe\xd7\x0a\x87\x5a\x88\xa4\x2f\x7c\xed\x53\x38\x41\x65\x5c\xb0\xd2\x9e\x68\x7c\x8a\xda\x93\xd9\xc4\xe7\x76\xaf\xb9\x83\xe1\x33\x2f\x19\x20\x8e\xda\x7b\x78\x71\xf8\x8c\x06\x20\x29\x44\x51\xbc\x4a\x50\x0c\xc1\x7c\xb9\x56\xd7\x88\x7f\xa3\xed\x15\xb6\x56\x49\xa9\x36\xb4\x45\x54\x2b\xcb\x8c\x65\xb7\x1b\xfc\xa6\xbc\xec\x90\x1a\x16\xf5\xe2\x56\x2e\xfa\x8b\xfe\x70\x97\x8a\x03\xce\xe7\x42\xaf\xd1\xa2\xbf\xb9\x99\x68\x95\x7a\xbf\x2d\x1a\x77\x6e\x49\x89\x54\xe8\x68\x21\xe2\x10\xab\x46\x17\xc3\x4f\xf8\xd7\xef\x1d\xa3\xca\xe0\x2c\xd2\x79\x79\x5a\x7d\xbf\xc9\x35\x99\xbc\x9d\xfa\xa1\xa7\xd8\x39\x91\xab\x1a\x39\x6e\x97\xa6\x9d\x4b\x2c\x86\xcf\xba\x53\xd6\x01\x1c\xcb\x52\xd1\x57\xc0\x2f\x34\x34\x88\x26\xfb\x60\x35\x58\xcd\x9a\xb5\x1d\xd0\xec\xc2\x34\x92\x67\x70\x74\x56\xed\x55\x0e\x3a\xd8\xda\x6a\x4a\x5c\x6f\x49\xaf\x33\x7a\xd3\x58\xeb\x6e\x83\x57\xc3\x06\xdc\x30\xc6\xae\xc5\x34\x5b\x24\x11\x1e\xb8\x17\x25\xdf\x78\x71\x64\x26\x65\x07\xea\x1e\x02\x16\x1c\xa2\x39\xaa\x26\xca\xea\x18\x86\x58\xaf\x4d\x9e\xa9\xd3\xbf\x75\x4a\x02\x2c\xf7\x56\x80\x0d\x9f\x82\x73\x30\xc3\x3a\xb6\x99\x28\x47\x8d\x63\x72\x4a\xaa\x0c\xe0\x9e\x76\x3e\xd9\x0d\x24\x1b\x13\xa0\x71\xd8\x2a\x27\xf6\x2c\x32\x53\x93\x6d\x3e\xa9\xd6\x10\xdb\x1e\x55\x67\x19\xa7\x8b\x4a\xd5\xea\xd4\x73\x96\xc5\x84\x69\xed\x36\x9b\x39\x23\x71\xda\x54\xa0\xba\x35\x01\x96\xc6\x1d\xaa\x47\x27\x61\x1b\x6a\xa3\x5d\x65\x18\xea\xc2\x71\x36\x09\x1f\x29\xe6\x98\xc5\x18\xdb\x04\xf9\xdb\xe8\x2c\x02\xa7\x2e\x52\x77\xdf\x60\x70\x12\xf0\xff\xd1\x6c\x29\xea\x6f\x23\x13\x61\x15\x92\xe7\x8b\xb1\x05\x31\xb7\xa4\x2c\x16\x63\x9b\xe8\xdf\x90\x5c\xe9\xb6\x3f\xbf\x4e\x41\x2a\x87\x01\xd8\x30\xd6\xba\x98\x6a\xfe\x33\x35\xa7\x3c\x4c\x43\x82\x19\x8b\x6a\x37\x3e\x52\x09\x54\x05\x95\x10\x6d\x0c\xb3\xda\x52\xd3\xb5\xfb\xc6\x79\x87\xbb\x9f\x6d\x5c\xc3\x5b\x15\x0d\x3f\x0d\x37\x47\x40\x09\xb2\xa2\x75\x84\x9f\x3b\x44\x9d\x19\x97\xe4\x6f\xbf\x80\xe1\xb2\xa3\x36\xde\x51\xa2\xcf\xdb\xa2\xb0\x27\x77\x3c\x7f\x66\x47\x3d\xaa\x9b\xbc\x24\x7c\xf1\x77\x37\x1b\x6d\x74\xa1\x6a\x16\xb8\xcd\xf2\x5e\x55\xc6\xe8\xd6\xb6\xc1\x5f\x8d\x14\x13\x8d\xae\x45\x89\x5e\x3a\xd6\xc9\x61\xdf\x39\x30\x98\xad\x60\x14\x0e\xf9\x61\xa9\xba\x84\xd0\xc6\x41\xde\xeb\x4a\x35\xb9\xc9\xc3\x9d\x3b\xe3\xc3\x11\x6d\x71\x03\xb5\x27\xee\x02\x9c\xab\x2e\xc0\x26\x38\xde\xb9\xe3\xa7\x11\x7d\xd6\xf4\x3f\x36\x55\x39\xd4\xf9\x51\x6c\x7e\xf1\x0f\xed\x74\x2b\x1a\x3a\xf2\x48\x29\x59\x27\xd4\x72\x7a\xcf\x8e\x4f\x3b\x4e\x8d\x31\x90\x57\xe5\x38\x0e\xd7\x8b\x5f\x77\x20\xa2\x47\x0f\x3d\x2f\x9d\x65\x68\xb8\xcf\xe2\xa2\x90\x41\x31\x66\x89\xe3\x2e\x86\x31\xb4\x6f\xef\xb2\x55\x1a\xf4\xd5\x13\x4c\xea\xd1\xbb\xd2\x10\x72\x8d\x79\x8b\x8c\x34\x57\xd8\x5d\xae\xae\x43\xf6\x65\x14\xed\xb3\x28\x76\xff\x72\x15\x33\xc3\x00\x7f\x85\x94\x54\xde\x83\xf8\x7e\xcb\xe6\xd7\x8a\x15\xcc\x95\xf9\xdd\xc7\xa9\x67\x7a\x21\x5f\x4b\x49\x6a\xfe\x8a\x37\x8e\xc6\xc3\x63\x1e\x06\xd9\x14\x8b\xd8\x83\xd9\x15\x93\xd8\xa2\x72\x1b\x88\x44\x88\xc9\x16\xfa\xc9\x3d\xd7\x77\x16\xde\xbb\x90\x61\x47\x51\x4c\x4b\x53\x38\x51\x06\x72\x01\x65\x0c\x4b\x3c\x11\xd5\x3d\x86\xfa\x57\x9b\xba\xaf\xec\xef\x95\xef\xb4\xdd\xc9\xc2\x66\xa1\x27\x2a\x83\xd0\xec\x13\x0f\x61\x36\xa3\x32\x16\x3d\x71\xcf\xdf\xaa\xe6\x48\x97\xd5\xde\x9e\xf2\x2b\xe0\x16\xdf\xae\x73\x33\xc1\xfa\xc1\xba\xc6\xda\xf9\x52\x38\x4b\xef\xb8\x3d\xdf\x99\x75\x2c\xf3\x9c\xf9\xe7\x29\xd8\x2c\x0f\xbc\xb2\xc6\x26\x6d\x5a\xe5\x8d\xce\xb6\x78\x6d\x4a\xe0\x38\x44\x8e\x40\xc5\x19\xc0\xe6\xd7\x1a\x67\x6f\x18\x46\x55\xb0\xdc\xc3\x6e\x3d\xf7\xac\xcd\x97\xdb\xe0\x58\x41\xf6\x18\xc9\x35\x9b\x99\x0b\xef\x1e\xb0\xca\x0a\x66\xca\x4d\x8f\xa4\x1b\x73\xf7\xb0\x7f\xf2\x8d\xb7\x1a\x81\x2f\x01\xc7\x97\x82\x73\xaf\x49\x38\xbe\x44\x13\x4b\x22\x8e\x2f\xfb\x65\x93\xc9\x38\x3e\xb0\x1a\x09\x39\x3d\x80\x5a\x33\x29\xc7\x9f\x96\x63\x4b\xcc\xe9\x90\x9a\x63\xe3\x66\x47\x55\xb7\x9e\x39\x3a\x5d\x8c\xe3\x6e\xf9\x39\x1b\x8d\xc0\x59\x72\x71\x7a\x98\xba\x5c\xe8\xa8\x59\x89\xf7\x25\x76\x5e\x73\x21\x60\xdf\x59\xeb\x98\x91\x63\x4b\x07\x9d\x88\x2d\xdc\x02\xed\x23\x0c\x81\xb9\xbf\x4e\xda\x27\x43\xe5\xee\x76\x11\x79\x82\x25\x9b\x87\x6e\xee\x83\x55\xa1\x7b\xc9\xe9\x66\x86\xa3\x20\x0d\xf1\x0c\x99\x85\x94\x66\x06\xdd\x5a\x75\x3e\x6b\x58\x30\x6a\x06\x43\x26\xd4\x51\x77\xda\x56\x1e\x94\xcb\x5c\xa6\x65\xf8\x97\x29\xc5\x37\x2b\x79\x0d\xa9\x82\x05\x41\xe5\xe7\x37\x41\xa7\x52\x1a\xf1\x28\xa8\x92\x55\x03\xa3\x0c\x3e\x87\x46\xe5\x55\x59\x59\xea\x7e\xd7\x34\x33\xae\x6f\x44\xd1\xdf\x95\xf2\xcc\xd6\x2f\x17\x69\x64\x93\xb4\x8a\x5d\x63\x7f\xd8\x28\x44\xed\x16\xbb\x38\xc1\xd0\x72\x1b\x93\xc9\xd8\x64\x92\xb8\x2c\x68\x32\x91\xbb\xc5\x22\xbe\x8a\xe9\xce\x71\x24\xca\xa7\xb2\xfc\x29\xb8\x79\x49\x61\x5a\xb8\x43\x41\xf3\x6b\x58\xf1\xd0\x62\x3d\xf1\xed\x41\xe1\x3e\xf2\x0d\x3a\x80\x81\xd7\x9a\xc5\x50\x9b\xd5\xab\x5b\x4a\x59\x7a\xc1\xb9\xd7\x1c\xbf\x8e\xe4\xe9\x5e\x62\x52\x3e\x6e\xaf\xe5\xe0\x1d\xd4\x12\xdb\x56\xeb\x78\x83\x46\xca\xe0\x57\xae\x84\xa7\xd9\x42\xc7\xf5\xbd\x5e\x11\x87\xf6\x63\x61\xab\x4d\xde\xdd\x94\x77\x38\x5b\xe1\x43\x15\x2d\x5b\x65\xfe\xa9\xf0\xd6\xae\x5e\xa7\xb6\x83\x5e\xb5\xfa\xfe\x8b\x56\xab\x24\x58\xb3\x78\xb5\xaa\xe9\xfe\x53\xc4\xfa\x0f\x5b\xc4\xda\xc6\x11\x76\x66\x68\x16\xb3\x3e\xe6\x35\xa7\x2d\xa9\x68\xd6\xda\xc3\x9b\x2a\x75\x7d\xfc\xa7\x29\x75\xdd\xbe\x5c\x5d\x82\xcb\x55\x18\xda\x31\x49\x9b\x11\x6f\x6b\x96\xbe\xee\x8d\xad\xb7\x04\x76\x2b\x3b\xfe\x29\x4a\x61\xe3\xb3\x61\x94\xda\xbd\xa4\x53\x96\xfd\x48\xa2\x6c\x16\xc0\x12\x46\xff\xb1\x50\x52\x41\x1c\x7b\x02\x9b\xf0\x82\x86\x9f\x00\x24\xc3\x5f\xf5\xa6\x19\x37\x63\x34\xcf\xcf\xce\x5d\x6c\xc1\x1b\x17\xc2\x53\x80\x86\x84\x7f\x78\xb3\x2e\x57\xcd\x3f\xd9\xbb\xc1\x04\x64\x9f\x3d\xa5\x7e\x7c\xb5\x39\xcf\x8e\x12\x0d\xda\xb7\x7a\xf7\x6d\x8f\xd5\xaf\xf6\x8e\xda\x99\xba\x5d\x89\x54\x84\x72\xe8\x11\x7e\x6d\xd8\x71\xe9\x70\xfe\xd5\x9a\x7d\xd1\xcf\xf0\xf6\x30\x8b\xb4\xba\xd5\x43\xe4\x9f\x9d\x65\xb2\x24\xba\x67\xae\x11\xa4\xa6\x37\x5f\x06\xb7\x76\xfb\x42\xd0\x1b\x8a\xd0\xae\xf2\x85\xa0\x33\x55\x30\xb8\x9c\x0a\xd9\x5d\x6b\xa9\x7a\xe8\x08\x39\xcf\x92\x74\xd2\xcb\x2f\x89\x4b\x2f\xa7\xd6\xba\x5d\x42\xbe\x19\xad\x5d\xd1\xa1\x99\x8c\xb6\x32\xfe\x5c\xef\x94\x39\xad\x70\xef\x1c\xf8\xba\x80\x97\xfc\x26\xe3\x46\xbf\x8c\x6a\x0e\x7d\x52\xd2\x19\x46\x56\x87\x61\x36\x9b\xc3\x13\x39\x74\x4b\x51\x97\x23\xd6\x7a\x69\xd3\xb2\x46\x85\xb7\x50\x6b\xb9\x5a\x4d\x2d\x75\x34\x4f\x26\xbf\x68\x86\x06\xb4\x0d\xb0\x7f\x6b\x7d\xa3\x31\x55\xe1\x3c\xd0\x4f\xdd\x81\x55\x30\x8b\xc5\x0e\x16\x83\x90\x5b\x08\x39\x3f\xe9\x8f\x22\xb3\x3e\x35\x71\x9f\xc1\x71\x81\xa0\x71\x36\xed\x24\xea\x76\x88\x55\x9f\xb3\x6e\x96\x8d\x98\xbd\x03\x14\x17\x78\xbc\x41\xfe\xb6\x7f\x63\xa3\xeb\xd9\x86\xfb\x60\x74\x13\xdb\x3b\xce\x6c\xe2\x0f\xdb\x32\x48\xd7\x99\x82\x76\x92\x8b\x6c\xd1\x4a\xaf\x8b\x37\x07\xfc\x01\xde\xf9\xf8\x51\xb2\xcc\xc7\x8f\xa8\x07\x50\x82\xb1\xa3\x41\xaa\xd9\xf3\x79\xe6\xc6\x97\x30\x59\xc1\x2f\xa1\x47\xd8\x5f\x64\xf9\x38\x8e\x22\x9a\xfe\x45\x5a\xd9\xf5\x27\x0f\xfa\xd2\x96\x7d\xee\xc0\x20\x6d\x8b\xf7\x26\x70\x38\x51\xb9\xba\xff\xf7\xef\xdc\xa1\xde\x5e\xf5\x98\x04\x5d\x05\x3a\xe2\xe3\x07\x8c\x26\xcd\x2f\xb4\x1a\x78\x79\xbf\x6c\xe0\xfc\xb6\x41\xf3\xeb\x06\x6c\xb4\x46\x56\x46\x57\x32\x36\x13\x2a\x56\x59\x1c\x8a\x80\xaa\x21\x61\xa2\xcb\x9f\xe3\x77\xd7\xf9\x1a\x2d\x54\xf6\xe7\x6c\xf8\xb3\x36\xee\x39\x6f\xc3\x9f\x8b\x60\xcd\xdd\xf0\x27\x4a\x6c\x36\x7f\xa3\xb5\x82\xc9\x5a\xc0\xad\x9d\xc7\xd1\x96\xc9\x61\xe4\x72\xc8\x9b\x2d\x5f\x54\xd6\x2c\x32\x6b\xc9\x91\xb5\xd5\x4c\x83\x14\x2d\x39\x21\x78\xad\x6f\xec\x58\x30\xe9\xe7\x3f\x6d\x22\x30\x65\x05\xa3\x7f\x15\x17\x5d\x34\x0e\x3f\xc9\x9f\x27\x1d\x4e\xa8\x1a\x20\x74\x9b\x4d\xfe\x8a\x45\xfa\x59\x42\x8a\xa1\x3e\xc0\xdd\x1d\x66\x6c\x10\xb3\xfb\xb9\xc4\x21\x66\xc1\x07\x20\x95\x54\xb0\xbc\x41\x34\xac\xcf\x65\x99\xbd\x5e\xe4\x63\x42\x3e\x8c\x61\x65\x94\x2d\x3a\xe4\x9c\x1f\x25\x08\x9b\x03\xae\x4a\xac\x73\x89\xf0\xd6\xff\x05\x00\x00\xff\xff\xb9\xcb\x90\xf1\x2b\xb3\x00\x00")

func itsyouonlineRamlBytes() ([]byte, error) {
	return bindataRead(
		_itsyouonlineRaml,
		"itsyouonline.raml",
	)
}

func itsyouonlineRaml() (*asset, error) {
	bytes, err := itsyouonlineRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "itsyouonline.raml", size: 45867, mode: os.FileMode(420), modTime: time.Unix(1470663734, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _securityschemesOauth_2_0Raml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x92\x31\x8f\x13\x41\x0c\x85\xfb\xfd\x15\xd6\x36\x34\x90\x1c\xa7\xab\xb6\x3b\x84\x44\x09\x05\xa9\x10\x8a\xe6\x66\x4c\x32\x62\x63\x0f\xb6\x37\x21\x88\x1f\x8f\x27\x9b\x8d\x92\x28\x91\xa8\x2e\x4d\x56\x6f\xe6\x3d\x7f\xf6\x38\xa1\x46\xc9\xc5\x32\x53\x07\x7f\x1b\x80\xcf\xcf\x83\xad\x1f\x21\x2b\x04\x28\xc2\xc6\x91\x7b\xb0\x75\x30\xe8\xd1\x14\xf0\xb7\xa1\x50\xe8\x21\x94\xa2\x20\xf8\x6b\x40\x35\x08\xee\x61\xc9\x7f\x42\xcd\x01\x63\x77\xe6\x6d\x30\xf4\xbc\x84\x16\x72\xaf\x90\xc9\x03\x07\x45\x79\xe3\xc9\x31\xf2\x40\x06\xbb\xec\xb6\xc1\x60\x85\x66\x99\x56\x5e\x06\xb3\x40\x09\xaa\x3b\x96\x34\x6b\x6c\x5f\xb0\x1b\x89\xe0\x71\xf6\xd0\x8c\xb0\x2f\x98\x3e\xec\x3b\x8f\x5e\x63\x48\x28\x5a\x3f\x01\x9e\xcf\x11\x46\xa9\x16\xbf\xea\x6e\xfc\x2d\x14\x53\xa5\x54\xa4\xe4\x54\xdb\xd0\xe7\x34\x95\xa9\x70\xa8\xea\xc7\x3f\x91\x66\xf0\x91\x81\xd8\x2a\xb8\x2b\xce\xb9\x46\x39\x60\x9f\xb2\x5c\x81\x76\xf4\x2c\x0f\x9e\x16\x7c\x26\xb2\x07\x35\xa9\x3d\x95\x20\x61\x83\x3e\xb4\x99\x5b\x0e\x27\x5f\x26\xe5\x48\x7e\x6e\x7e\x6d\xf0\x8b\xa9\xb5\xc7\x89\xfa\x15\x41\x2d\x4c\x8a\x47\xc4\xa7\x87\xa7\x9b\x64\x0b\x9a\x5e\x1e\x53\xa3\xe3\x2b\x1e\x2c\x17\x0b\xb1\x90\xdc\xc1\xda\xac\x68\x37\x9f\x67\xd3\x3d\x0f\x33\xa6\x3e\x13\xce\xb7\xef\xe7\x5c\xef\xce\x4f\x39\xcd\x34\x90\xaf\xb5\x8f\xff\xb2\x9e\xcd\xef\xba\xf4\x27\x09\x64\xda\xc1\xb7\x4b\x79\x19\x39\xe1\x5b\x88\x7d\x46\xb2\x65\x14\x4c\xfe\x9f\x83\xaf\xe9\x77\x4f\xd0\xc8\x65\xea\xfc\x1d\xb4\x75\x69\xbb\x90\x36\x99\xda\x0b\x29\xd3\x0f\x3e\x29\x2c\xab\x40\xd3\xf6\xf1\x8e\x50\x6e\x1f\x6d\x70\xf3\x72\xef\xec\x7e\x60\x64\x32\x09\xd1\x3b\x11\x7f\xa1\xd3\x9d\xc8\x9b\x12\x68\x7f\x05\x37\xa9\x37\xaf\x5e\xd4\x98\xc4\xbb\xf1\xa3\x7c\x47\xf5\xcd\xb6\x1c\xb3\x27\x58\xdb\xfc\x0b\x00\x00\xff\xff\xbf\x96\x88\x8b\x48\x04\x00\x00")

func securityschemesOauth_2_0RamlBytes() ([]byte, error) {
	return bindataRead(
		_securityschemesOauth_2_0Raml,
		"securitySchemes/oauth_2_0.raml",
	)
}

func securityschemesOauth_2_0Raml() (*asset, error) {
	bytes, err := securityschemesOauth_2_0RamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "securitySchemes/oauth_2_0.raml", size: 1096, mode: os.FileMode(420), modTime: time.Unix(1465910933, 0)}
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
	"itsyouonline.raml": itsyouonlineRaml,
	"securitySchemes/oauth_2_0.raml": securityschemesOauth_2_0Raml,
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
	"itsyouonline.raml": &bintree{itsyouonlineRaml, map[string]*bintree{}},
	"securitySchemes": &bintree{nil, map[string]*bintree{
		"oauth_2_0.raml": &bintree{securityschemesOauth_2_0Raml, map[string]*bintree{}},
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

