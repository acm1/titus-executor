// Code generated by go-bindata. DO NOT EDIT.
// sources:
// default.json (17.888kB)
// fuse-container.json (18.199kB)
// allow-perf-syscalls.json (18.087kB)

package seccomp

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _defaultJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5b\x6f\x5b\xb9\x11\x7e\xdf\x5f\x21\xe8\x39\x28\x7c\xab\xd7\xce\x9b\xaa\xb8\x8d\xb1\x76\x9c\xda\x29\x76\x17\x45\x40\xd0\x3c\x73\x8e\x58\xf1\x66\x5e\x64\x0b\x41\xfe\x7b\xc1\x73\x24\x59\xf2\x0c\x4f\xbc\x2b\xa7\x51\x01\x3f\x24\xb6\xbf\x8f\x97\xe1\x90\x9c\x19\x0e\x79\xbe\xfc\x34\x18\x0c\x86\xdc\x8b\xc9\x25\x77\xc3\xb7\x83\x7f\xe7\xbf\x07\x83\xc1\x97\xc5\xcf\x25\x2b\x23\x88\x98\x3c\x0c\xdf\x0e\x86\x37\xe3\xcb\x8f\x6c\x74\x3d\x7e\xcf\x7e\x3b\x39\x66\xc7\x47\xc3\x37\x6b\x85\x43\xba\x1d\xad\x95\x0f\x6b\x6d\x2e\x4a\x6c\x54\x5f\xaf\xfb\x84\x3c\x3c\x18\x3e\x72\x9f\x17\xbf\x7e\x7d\xf3\x87\x04\x1c\xe5\xff\xb7\x90\x70\x74\x7d\xb9\xbd\x10\x97\xe7\x1f\x6f\xb6\x90\x21\x57\x2f\xab\xa9\x6b\xfc\xc3\x4b\x28\xeb\xb1\xa9\xef\x29\xea\xcb\xc8\x79\x76\xb1\xa5\x46\xcf\x2e\xfa\x05\xcd\x1d\xbc\x98\x56\x97\x8d\x7d\x6f\x81\xb7\x97\xf6\xe6\xf0\x74\xef\xb7\x3f\x2d\x67\xae\x4d\xc9\x90\x7f\x7e\x6e\x1b\x1d\x56\x50\xf3\xa4\xe2\x48\x44\x69\xcd\x63\xdf\xe3\x4f\xec\xec\xfa\xfa\xc3\xd5\x70\xa3\xd4\x99\xf7\xc6\x5e\x43\x1c\xbe\x1d\xec\x77\x44\x98\x07\xc1\x95\x0a\x25\x53\x85\x9b\x1d\x5d\x5c\x5c\xfd\xba\x31\x22\x69\x84\x4a\x55\x3b\x92\x2f\x9b\x23\x11\xdc\xe1\xf1\x65\x62\x3c\xfa\xc8\x6e\x7e\xbf\x61\xa3\x77\x97\xe7\x1f\x86\x1b\xf4\xe7\xc7\xbf\xbe\xae\xf7\x62\xb8\xa6\x94\x75\xeb\xea\xa7\xf3\x28\x94\x35\xf0\x14\xac\xb9\xb1\x51\xd6\x73\x26\x8d\x8c\x88\x0c\xc2\x9a\x5a\x36\x18\xd7\x36\x19\xa2\xb8\x75\x60\x30\xea\xa4\x98\x3e\x45\x95\xb5\xd3\xe4\x58\x25\xac\x9d\x4a\x24\x15\xd9\xbc\xb6\x33\x60\x24\x93\x75\xc0\xa2\x65\x13\x6e\x2a\x05\x8c\xa3\x02\x59\x2e\x16\x3d\xa0\x8e\x1c\xf8\x9a\xc1\x0c\x4c\x64\x94\xec\x77\xc9\x46\x2e\xa2\x7a\x8a\x07\x88\x95\xd5\x5c\x9a\xdc\x33\x41\x4e\x6c\x88\x05\xca\x04\x04\xce\x83\xb2\x48\xc7\x89\x1c\x69\x87\x1e\x20\xd8\x84\x09\xf7\xf0\xec\x7d\xf9\x9c\xe5\x5b\x58\x58\x5c\x08\x70\x48\xae\x0e\x3d\xa2\xe0\x80\x06\xcc\xab\xff\x44\xa9\xe1\x01\xe1\x8a\x7b\xfd\x14\xbc\x95\xa6\x42\x98\x47\xeb\x49\x70\xd7\x00\x12\x2b\x6f\x34\x02\x9d\x54\xd2\x63\x50\x5b\xd4\x91\x98\xd8\x7b\xb4\x28\x5a\xf0\x10\xcd\x81\x50\x56\x4c\xd9\x62\x6c\xbd\xe4\x31\xd2\x53\x47\x37\x10\xb3\xe1\xeb\xe1\xd8\x37\xaa\x97\xbb\x5e\x90\xa5\xba\x86\x1b\x1b\x14\x80\xfb\x06\x5d\x16\x20\x50\x1d\x07\x60\x9e\x9b\x06\x53\xd6\x18\x10\x78\x62\xac\x9b\xb3\x5a\xaa\x52\x2d\x0f\x78\x6b\x57\x09\xc9\x5c\x25\x87\x66\xa7\x4a\xee\xf0\x29\x06\xce\x2a\xc5\xda\x56\x51\x5f\xeb\xdc\x7e\x81\xc4\x76\x61\x45\x30\xab\xd0\x62\xea\x48\x77\xcf\xb1\xa1\x5d\xa3\x90\xe4\x1d\x57\xae\x95\x19\xb2\xbb\x6c\xd5\xea\x12\x8c\xbb\x79\x00\x31\xc3\x6a\x68\x51\xac\x75\x78\x20\xc4\x79\x90\x91\x35\xde\xe2\x09\xa9\x3b\x3b\x80\x9b\x59\x11\x48\x9c\x9a\x57\x33\x19\x88\xa5\xb6\x22\x18\xc5\x29\x65\x05\x31\x9b\x2b\x2f\xa7\x39\x36\x1d\x35\x69\x0f\x6a\xd2\x20\x74\x28\x31\x10\xd2\x52\xd4\x05\x53\xd1\xe1\x54\x2b\x06\xaf\xa9\x16\x24\xc6\x5a\xf1\xc8\xc3\xdc\x08\x44\x34\x10\x1f\x78\x8c\x78\x44\x4a\x86\x12\x63\xb1\x87\xae\x2d\xa1\x2a\x0f\xd9\x0b\xd3\x8d\x84\x52\xbf\x21\x12\x43\xcd\x20\x31\xaa\x0c\x17\x89\x1a\x19\xc7\x05\x4c\x95\xa7\x54\x13\x7d\x32\xe4\xfa\x58\x12\x44\x43\x29\x62\x4f\xd5\x82\x05\x5b\x58\xa7\x8c\x13\x8b\xbd\x81\x28\x5c\xa2\xd0\x7b\xb4\xce\x1a\x88\x15\x98\x88\xc6\xbb\xc4\x71\xb7\x0d\x44\x68\x24\xd5\x50\x86\xf1\x1a\xcc\x44\xa2\xcb\xa7\x42\x79\xba\xf9\x52\xeb\xad\x25\xa0\x06\xd0\x11\x64\x1d\x99\x55\x87\x16\x50\x03\xd1\x01\x78\x2a\xa0\xca\x14\x2d\x96\x6b\x3c\xb2\x43\x19\xa6\x0b\x17\x60\x2f\xad\x97\x71\x4e\x50\x9e\x9b\xca\xa2\x60\xa5\xf3\xd2\xb4\x40\x1d\x41\x0e\xdb\x43\xa0\xa7\xa2\x23\xe8\x3a\x4a\x6a\x6c\x83\x1b\x88\xcc\xdb\xdb\x14\x22\xcb\xbb\x9d\xaa\x97\x02\xc7\xae\xb5\x81\x18\x48\x09\x82\x15\xd3\x82\xe2\x33\x65\x71\x28\x98\x45\x88\x13\x0f\xbc\x62\xdc\x03\x27\xe8\x48\xf6\x94\xa7\xde\xd6\x15\xa7\xb4\x4d\x6b\xa7\xa4\x1a\xd2\x0a\xc9\x85\x07\xe0\x55\xc5\xee\x79\x14\x93\x52\x01\xea\x20\xb4\xce\xa1\x68\x60\x49\x7a\x5d\x68\xd8\x32\xc1\x8d\x00\x64\xd9\xa5\x25\x42\x08\x69\x59\x05\x21\x7a\x8b\xf4\x20\x6d\x8e\xe2\x5a\xff\x8d\xb6\x95\xb4\xcc\x3d\x8f\x2c\x58\x2d\x69\xf3\x62\x67\x44\x0c\xbd\x60\x88\x38\xba\x43\xb1\xbb\xcf\x78\xba\x25\x56\xa7\xb4\x2c\x79\x69\x1a\x06\x26\xe2\x6d\xbe\x62\x3d\x34\x32\xf4\x15\xa0\x7b\x75\xc8\xde\x4f\xa5\x42\xfa\x55\xa4\xa7\x56\x05\x4f\xad\x4a\xcb\x49\x49\x83\xcf\xb4\xd2\x4c\xb1\xd9\xcf\xfb\x10\x9f\x2c\x8b\xbe\x58\x15\x19\xa6\x54\x00\xc0\xbd\xf6\xf8\x64\xba\x42\xc9\x51\x2b\xca\x51\x2b\xda\x51\xeb\x2e\x0c\x43\x30\xe8\x5b\xee\xbd\xc4\x93\xa7\x41\xd7\x55\x21\xd8\xd6\xd2\x08\xeb\x31\x3c\x25\xe2\xb2\x16\xc4\x62\xea\xa9\xc1\xd1\x5a\x0b\x12\x65\xa9\x70\xa7\x05\xd1\xfc\xb7\x28\xc7\xab\x48\x6b\x8e\x16\x60\xc6\x70\x03\xce\xdb\x48\x1c\x75\xf4\x5d\xde\x6b\x01\x22\x35\x11\xfa\x8e\x75\x36\x85\x20\xa8\x34\x85\xbe\x6b\x77\x75\xe5\x41\x80\xc4\x41\xfc\x13\xba\x60\x01\x96\xa5\x02\xe0\x03\xf7\x3a\x57\xae\x9e\x0c\xb5\x2b\xb4\x07\x4a\x59\xa1\x21\x8c\x9f\x0e\x0d\x61\x81\x74\x68\xbc\x98\x11\x68\x20\x24\xa5\xe2\x3e\x9d\x0c\x39\xe5\x1d\x4c\x4d\x6f\x32\x84\xcc\xc5\x43\xb2\x81\xfb\x45\xe0\x8a\xf6\xac\x81\xfb\x00\x8a\x58\x00\xd4\x3c\x66\x8c\x4e\x5e\x11\x27\x24\xc7\x13\xde\x7f\x4e\x56\x75\x45\xae\x91\x8e\x69\x27\x30\xc8\xc6\x70\x34\x66\x27\x1d\xd1\x9c\x03\xdc\xb1\xc5\x0a\x73\x45\xb0\xb0\x5c\x9c\x27\xa6\xdf\xe5\xa8\x81\x2a\x0b\xbc\x42\xf3\xdf\xa1\x58\xba\x2e\x30\x22\x5a\xe9\xe6\x01\x5d\x05\x2d\xf1\x92\xa0\xf7\x5e\x52\xe7\x82\x0e\xc7\x52\x75\x30\x12\x2b\x0b\x4b\x61\x7c\x52\x20\xa8\xad\xb4\xc4\xf1\x1a\x21\x35\xe4\x01\xef\x9a\x8c\xd5\x1e\x47\xae\x19\xd7\x3a\xa0\x34\xe4\x12\x2f\x68\xa7\xa5\xa9\x5a\x9a\xbb\x2e\x93\xe3\x78\x83\x73\x5a\x3d\x2e\xcb\x03\x15\x6e\x76\x28\x35\xee\x0e\x27\xf4\x1d\x22\xf7\x91\x2d\x12\xf9\x88\xd6\x84\x6b\xf1\x01\xee\x10\x16\xf3\x76\x59\x24\x4c\x49\xce\x81\xa9\xa4\xc1\x2a\xe8\x48\x6f\x85\xe6\x01\xcf\x64\xcb\xde\x25\x48\x20\x4d\x6d\x69\xda\x43\x4c\xbe\xd0\x6b\x48\xc1\x11\x86\xba\x23\x5b\x5b\x4d\x65\x8c\x9e\xd0\xa5\x49\x8d\x2c\x36\x7d\xe2\x05\x31\x81\x2a\xfb\x2f\x5e\xd7\x39\x24\x46\x7e\xea\xb1\x00\x31\xc3\x2b\xd2\x71\xcf\xd1\x42\x5c\xb1\x6c\x79\xfa\x62\x9a\xa3\x03\x38\x55\x4a\x22\x55\xad\x4a\xb5\xbf\x24\x85\xa3\x92\xae\x84\xf7\x6d\x53\x32\x47\xa5\x33\x6c\x19\xc9\x42\x05\xdd\x75\x65\xc3\xb7\x54\x53\x70\xfc\x2b\xb2\x47\x35\xe1\x9b\xc3\x99\x4b\xc0\xe9\xc0\x00\x42\x58\x8d\x5c\x17\xed\x9c\x02\x68\xf2\xd6\x43\x13\xde\x39\x80\xb6\x44\xb3\xba\x5d\x65\x7d\x4c\x49\x83\xc4\xb2\xce\x58\xb6\x26\x25\x9c\x6e\x85\xb2\x67\x2d\x4e\xc3\x11\xaf\x73\x88\x35\x75\x9a\x5f\xe2\xf8\xbc\xd0\x32\xc4\x69\x75\x89\x93\x35\xe8\x1e\x4a\xed\xd3\x79\x95\x50\xce\xab\x84\x52\x5e\x25\xd0\xc9\x93\x50\x4e\x7c\x04\x88\x9e\xca\x32\x2d\x71\xb2\x73\x3a\x23\x12\xca\x19\x91\x50\xca\x88\x84\x72\x46\xa4\xa5\x4a\x55\x8a\x35\xc8\x1c\x4a\xe8\xcf\xa1\x04\x32\x55\x12\x8a\xf9\x90\xd0\x9f\x0f\x69\x69\x59\x31\x5e\x55\x9e\xb8\x2b\x0b\x64\xf2\x23\x14\x92\x1f\xa5\x93\x5d\x98\x68\xec\x3b\xc3\x84\xdc\xe4\x13\x5d\x51\x45\xa9\x9d\x3f\x49\xb1\x22\x0e\xd3\xd9\x61\xaa\x18\x22\xc7\x21\x77\x17\x7a\xe2\x5b\x89\x25\x8e\xf7\x71\xd9\x87\x16\x3d\x64\x9e\x08\x42\xda\x16\xa5\x82\x81\x8e\x71\x1c\x47\x04\xc1\x29\x29\xb0\xd5\x21\x02\x7d\xfa\x94\x4c\x27\xad\x4b\x39\xeb\x8c\x63\x4f\x37\xd7\x54\x28\xb8\x80\x09\x41\x88\xf3\x4f\xc6\x7a\xee\xd5\x32\x4d\x48\x39\x0f\x94\xef\x8f\xf8\x12\x3d\x36\x54\xa2\x85\xba\x8b\x6c\xcd\x50\x21\x09\xd0\x71\x15\x28\x28\x71\x0d\x44\x3b\x03\xef\x13\x9a\xf1\x15\x5f\xee\xb4\x78\x01\xda\xd1\xa1\xaf\x6e\xe8\xad\x5b\x4c\x6b\x2c\xd9\x3e\xb9\x1e\xe9\x72\xeb\x7d\xb2\x3d\xd2\x74\x7d\x34\xad\x91\x9c\xab\xc2\xbd\x48\xf9\x5a\x24\x15\x53\xd0\x89\xda\xab\x89\x0a\xea\xe9\x7c\x41\x87\xe2\x85\x9d\x28\x1d\xb4\xa0\x21\x2e\x5b\x56\x44\x21\xca\x48\xa4\x72\x66\xd4\xa5\xd7\x4c\xd3\x66\x20\x47\xd0\xa8\xdd\x0c\x62\x7b\x9d\x51\xe2\x8a\xa1\x3d\x2b\x92\xe0\xec\x45\x9f\x72\x94\x5f\x22\x69\x69\x7e\x01\x6f\x40\xe5\x16\x8e\xfe\x72\xb2\xd6\xed\x73\x5e\x19\x65\xd3\x0c\x21\xb0\x99\x66\x85\x33\xfa\x8a\x2f\x1c\x97\xa3\xe7\xe2\x65\xdf\xad\x70\xdf\x60\x49\x37\x87\xdd\xe9\xa4\x82\x87\xe1\xdb\xc1\xde\x1b\x44\x59\xb7\xea\x22\xff\x3b\xfb\xe7\x10\x97\x99\x71\x95\x20\x57\xdf\x60\xbe\xae\x0d\xe4\x39\xea\x03\x1f\xac\xe1\x2a\x47\x5a\xff\xbf\x2a\x38\x79\x55\xc1\xfe\xe1\xfe\xde\xcf\x07\xaf\x7a\xc8\x7a\x38\x79\xdd\x12\x83\xa3\x83\xd3\xa3\xd3\xe3\x9f\x0f\x4e\xff\xba\x8b\xba\x28\xbb\x03\xee\xc5\x84\x90\xa4\x15\xc6\x89\xe3\x23\x05\x5b\x3d\x4a\x7d\x12\x87\xe2\xe3\xcb\x3d\x77\xc2\x9a\x08\x0f\xf1\xc7\x0f\x18\xbf\x45\x5c\xa0\x1b\xef\x9f\xff\xb8\x12\xb8\xd7\xac\xef\x11\x53\xe6\x9f\x11\xb0\xf7\xe9\xf1\xd6\x03\x9f\x3a\x2b\xf1\xfb\x51\xc1\xc5\x04\x6a\x95\x02\xba\xa5\x6e\xcf\xa1\x2a\xec\x80\xde\x35\xbe\x02\xc8\xf8\xc3\xc6\x2b\xf9\x3f\xa3\x77\x31\x61\xdd\xb5\xc3\x4e\x8f\x91\x40\x4f\x8e\xb7\x1a\xb9\xb6\x95\xac\xe7\x4c\x55\x3b\xb0\xab\xc2\xe1\xe9\x1e\x31\xc4\x0c\x3f\x6c\x67\x5b\x0e\x4f\xf7\x98\x13\x92\x69\x2d\x2d\xa3\xae\x5b\x36\x4b\x90\xb1\x77\x5b\xc4\x27\x93\x0f\x07\x4c\x9a\x10\xfd\xff\x48\x63\xbd\x5f\x04\xbc\x1b\x8d\xd9\xf5\xd9\xe8\x1d\xbb\x39\x1b\x5d\x8f\xdf\x6f\xa5\xa6\xf6\x25\xfc\xed\x7c\xed\xa9\xfc\x0e\x0c\xf0\xf5\x93\x87\xd7\x4f\x1e\x76\xe1\x93\x87\x17\x0f\x12\x2f\x47\x37\xbf\x9c\xbd\xeb\x8f\x15\x0f\xf6\xf7\x8f\xf6\x8e\xf7\x0e\x4e\x9e\x77\x8e\x82\x87\xef\x6b\x75\x51\xc0\xf0\xdd\xf7\x6d\xb7\x49\x7f\xe0\x3c\xee\xff\x88\x79\x14\x56\x6b\x30\x31\x77\x93\xa7\x61\xd0\xde\xf8\x41\x04\x3f\xb0\xbe\x02\x2f\x4d\x33\xa8\xad\x1f\xb4\xca\x19\xc8\x30\xa8\x64\x5d\x83\x87\xcd\xfd\xd1\xb3\x1a\x5e\x72\xe2\x76\xcb\xd3\xbf\xfc\x7a\xd9\xc2\x6d\xfd\xed\xea\xea\xd3\x56\xa3\xf1\x70\x6b\xed\xce\x78\xe1\xf1\xfb\xeb\x6d\x07\x24\x26\x7e\x87\x06\x74\x79\xf5\xee\x5f\x17\x67\x5b\x0d\xa8\xbb\x92\x60\xda\x56\x09\xdf\x40\xe7\x18\xa2\x40\xd5\xeb\xdc\x8e\xa8\xe3\xe3\x68\x3c\xde\x6e\x7a\xb9\x10\x3b\x33\xb9\x1f\x3f\x5d\x8f\xc6\xdb\x4d\xee\x54\xe0\x37\x11\xdd\x53\xb9\x06\x88\xaf\xa7\x96\x09\xed\xc2\xf3\xd7\x5d\xcc\x87\x6f\xa1\xdf\xeb\xd1\xaf\xe7\x57\x5b\xa9\x57\x5a\x47\xbc\x70\x77\xe0\xf5\xae\x8c\xf1\xd3\xf9\xe5\x76\x2b\x28\xf4\x7c\xba\x10\xca\x1f\x64\x2e\x2f\xf5\x76\x45\x0d\x9f\x7e\x67\xe3\xab\x0f\x7f\x3f\xff\xc7\x56\xca\x98\x4d\xb8\x69\x92\xdb\x95\x51\x7d\x38\xdf\xd2\x3c\x34\x10\x99\x06\xed\xac\x92\x02\xbf\xc6\xa6\xbe\x4d\x0e\x1b\x35\x76\x43\x0f\x17\x57\xdb\xcd\xea\xe2\xb8\x46\x0c\xe6\xa7\xf6\x8f\xaf\x3f\xfd\x37\x00\x00\xff\xff\xb5\x98\x78\xcb\xe0\x45\x00\x00")

func defaultJsonBytes() ([]byte, error) {
	return bindataRead(
		_defaultJson,
		"default.json",
	)
}

func defaultJson() (*asset, error) {
	bytes, err := defaultJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "default.json", size: 17888, mode: os.FileMode(0664), modTime: time.Unix(1648072800, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcf, 0x25, 0xb7, 0xde, 0x80, 0x4b, 0xd1, 0x3a, 0x9e, 0xd, 0x4a, 0x13, 0x75, 0xae, 0xef, 0xda, 0x23, 0xc1, 0x5e, 0x3f, 0xab, 0x31, 0x6f, 0x61, 0xbd, 0xf, 0x80, 0x19, 0xde, 0x38, 0xf7, 0xea}}
	return a, nil
}

var _fuseContainerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5b\x6f\x5b\xb9\x11\x7e\xcf\xaf\x10\xf4\x1c\x14\xbe\xd5\x6b\xe7\x4d\x75\xdc\xc6\x58\x3b\x4e\xed\x14\xbb\x8b\x22\x20\x68\x9e\x39\x47\xac\x78\x33\x2f\xb2\x85\x20\xff\xbd\xe0\x39\x92\x2c\x69\x86\x27\xce\xca\x69\x54\x20\x0f\xbb\x49\xbe\x8f\x97\xe1\x0c\x39\x33\x1c\x1e\x7d\x7e\x35\x18\x0c\x86\xdc\x8b\xf1\x15\x77\xc3\x37\x83\x7f\xe7\x7f\x0f\x06\x83\xcf\xf3\x3f\x17\xac\x8c\x20\x62\xf2\x30\x7c\x33\x18\xde\x9e\x5d\x7d\x60\xa3\x9b\xb3\x77\xec\xf7\x93\x63\x76\x7c\x34\x7c\xbd\xd2\x38\xa4\xbb\xd1\x4a\xfb\xb0\x32\xe6\xbc\xc5\x5a\xf7\xd5\xbe\x1b\xe4\xe1\xc1\xf0\x89\xfb\x34\xff\xeb\x97\xd7\xdf\x24\xe0\x28\xff\x7f\x0b\x09\x47\x37\x57\xdb\x0b\x71\x75\xf1\xe1\x76\x0b\x19\x72\xf7\xb2\x9a\xba\xc1\xdf\xbf\x84\xb2\x9e\x86\xfa\x9e\xa2\xbe\x8c\x9c\xe7\x97\x5b\x6a\xf4\xfc\xb2\x5f\xd0\x3c\xc1\x8b\x69\x75\x31\xd8\xf7\x16\x78\x7b\x69\x6f\x0f\x4f\xf7\x7e\xff\xd3\x72\xe6\xde\x94\x0c\xf9\xcf\x4f\xed\xa0\xc3\x0a\x6a\x9e\x54\x1c\x89\x28\xad\x79\x9a\xfb\xec\x23\x3b\xbf\xb9\x79\x7f\x3d\x5c\x6b\x75\xee\xbd\xb1\x37\x10\x87\x6f\x06\xfb\x1d\x11\x66\x41\x70\xa5\x42\xc9\x55\xe1\x61\x47\x97\x97\xd7\xbf\xad\xad\xc8\x70\x4d\x2d\x43\x28\x6b\x60\x53\xc7\x75\x10\xd6\xd4\xb2\xc1\xb8\xb6\xc9\x44\x0c\x5b\x07\x06\xa3\x4e\x8a\xc9\x26\x4a\x0e\xa0\xed\x14\x18\xc9\xa4\x1e\xf4\x00\xc1\x26\x8c\xb9\x87\x67\xef\x87\xe7\xa8\x4d\x1a\xa1\x52\xd5\x6a\xee\xf3\x86\xe6\xb8\xc3\xfa\xcc\xc4\xd9\xe8\x03\xbb\xfd\xe3\x96\x8d\xde\x5e\x5d\xbc\x1f\xae\xd1\x9f\x9e\xfe\xf5\xe5\x19\xc6\xb9\x73\xf5\xe6\x1a\x69\x7b\x71\x63\xa3\xac\x67\x4c\x1a\x49\x58\xe7\xbb\x19\x53\x59\x3b\x49\x8e\x55\xc2\xda\x89\x44\x52\x7d\xa3\xa9\xb3\x0e\x58\xb4\x6c\xcc\x4d\xa5\x80\x71\xd4\x20\xcb\xc5\xa2\x07\x34\x91\x03\x5f\x33\x98\x82\x89\x8c\x92\xfd\x3e\xd9\xc8\x45\x54\x9b\x78\x80\x58\x59\xcd\xa5\xc9\x33\x13\xe4\xd8\x86\x58\xa0\x4c\x40\xe0\x2c\x28\x8b\x74\xfc\xc3\xb7\x6f\x61\x63\x71\x21\xc0\x21\xb9\x3a\xf4\x88\x82\x03\x5a\x30\xaf\xfe\x13\xa5\x86\x47\x84\x2b\xee\xf5\x26\x78\x27\x4d\x85\x30\x8f\xf6\x93\xe0\xae\x01\x24\x56\x3e\x68\x04\x3a\xae\xa4\xc7\xa0\xb6\x68\x22\x31\xb6\x0f\x68\x53\xb4\xe0\x21\xb2\x81\x50\x56\x4c\xd8\x7c\x6d\xbd\xe4\x31\xd2\x53\x47\x37\x10\x73\xbc\xe8\xe1\xd8\x57\xba\x97\xa7\x9e\x93\xa5\xbe\x86\x1b\x1b\x14\x80\xfb\x0a\x5d\x16\x20\x50\x13\x07\x60\x9e\x9b\x06\x53\xd6\x18\x10\xd8\x30\xd6\xcd\x58\x2d\x55\xa9\x97\x07\x7c\xb4\xab\x84\x64\xae\x92\x43\xd6\xa9\x92\x3b\xdc\xc4\xc0\x59\xa5\x58\x3b\x2a\x9a\x6b\x95\xdb\x2f\x90\xd8\x2f\x2c\x09\x66\x15\xda\x4c\x1d\xe9\x1e\x38\x76\xb4\x2b\x14\x92\xbc\xe3\xca\xbd\x32\x43\x4e\x97\xbd\x5a\x5d\x82\xf1\x34\x8f\x20\xa6\x58\x0d\x2d\x8a\xb5\x0e\x8f\x84\x38\x8f\x32\xb2\xc6\x5b\x6c\x90\xba\xf3\x03\x78\x98\x25\x81\xc4\xa9\x79\x35\x95\x81\xd8\x6a\x4b\x82\x51\x9c\x52\x56\x10\xd6\x5c\x46\x39\xcd\xb1\xeb\xa8\x49\x7f\x50\x93\x0e\xa1\x43\x89\x85\x90\x9e\xa2\x2e\xb8\x8a\x0e\xa7\x46\x31\x78\x4f\xb5\x20\xb1\xd6\x8a\x47\x1e\x66\x46\x20\xa2\x81\xf8\xc8\x63\xc4\x2b\x52\x32\x94\x18\x8b\x23\x74\x6d\x09\x55\x79\xc8\x51\x98\x1e\x24\x94\xe6\x0d\x91\x58\x6a\x06\x89\x55\x65\xb8\x48\xd4\xc8\x39\xce\x61\xaa\x3d\xa5\x9a\xe8\x93\x21\xf7\xc7\x82\x20\x06\x4a\x11\x47\xaa\x16\x2c\xf8\xc2\x3a\x65\x9c\xd8\xec\x0d\x44\xe1\x12\x85\x3e\xa0\x7d\xd6\x40\xac\xc0\x44\xb4\xde\x05\x8e\xa7\x6d\x20\x42\x23\xa9\x81\x32\x8c\xf7\x60\x26\x12\xdd\x3e\x15\xda\xd3\xc3\x97\x46\x6f\x3d\x01\xb5\x80\x8e\x20\xfb\xc8\xac\x3a\xb4\x81\x1a\x88\x0e\xc0\x53\x09\x55\xa6\x68\xb1\x5c\xe3\x91\x1f\xca\x30\xdd\xb8\x00\x7b\x69\xbd\x8c\x33\x82\xf2\xdc\x54\x16\x25\x2b\x5d\x94\xa6\x05\xea\x08\x72\xd9\x1e\x02\x6d\x8a\x8e\xa0\xfb\x28\xa9\xb1\x0f\x6e\x20\x32\x6f\xef\x52\x88\x2c\x9f\x76\xaa\x5f\x0a\x1c\x87\xd6\x06\x62\x20\x25\x08\x56\x4c\x0a\x8a\xcf\x94\xc5\xa9\x60\x16\x21\x8e\x3d\xf0\x8a\x71\x0f\x9c\xa0\x23\x39\x53\x36\xbd\xad\x2b\x4e\x69\x9b\xd6\x4e\x49\x35\xa4\x17\x92\xf3\x08\xc0\xab\x8a\x3d\xf0\x28\xc6\xa5\x06\xd4\x45\x68\x95\x43\xd9\xc0\x82\xf4\xba\x30\xb0\x65\x82\x1b\x01\xc8\xb3\x4b\x4b\xa4\x10\xd2\xb2\x0a\x42\xf4\x16\xe9\x41\xda\x9c\xc5\xb5\xf1\x1b\x1d\x2b\x69\x99\x7b\x1e\x59\xf0\x5a\xd2\xe6\xcd\xce\x88\x1c\x7a\xce\x10\x79\x74\x87\xe2\x70\x9f\xf1\x74\x47\xec\x4e\x69\x59\xf2\xd2\x34\x0c\x4c\xc4\xc7\x7c\xc9\x7a\x68\x64\xe8\x6b\x40\xcf\xea\x90\xbf\x9f\x48\x85\xf4\xab\xc8\x48\xad\x0a\x91\x5a\x95\xb6\x93\x92\x06\xdf\x69\xa5\x99\x60\xb7\x9f\xcf\x21\xbe\x59\x16\x63\xb1\x2a\x32\x4c\xa9\x00\x80\x67\xed\x89\xc9\x74\x87\x52\xa0\x56\x54\xa0\x56\x74\xa0\xd6\x5d\x1a\x86\x60\xd0\x77\xdc\x7b\x89\x8d\xa7\x41\xd7\x55\x21\xd9\xd6\xd2\x08\xeb\x31\x3c\x21\xf2\xb2\x16\xc4\x62\xea\x89\xc1\xd9\x5a\x0b\x12\x6d\xa9\x74\xa7\x05\x91\xfd\x5b\x94\xe3\x5d\xa4\x35\x47\x1b\x30\x63\x78\x00\xe7\x6d\x24\xae\x3a\xfa\x3e\x9f\xb5\x00\x91\x32\x84\xbe\x67\x9d\x4f\x21\x08\xaa\x4c\xa1\xef\xdb\x53\x5d\x79\x10\x20\x71\x12\xbf\x41\x17\x3c\xc0\xa2\x55\x00\x7c\xe1\x5e\xe5\xca\xdd\x93\xa1\x4e\x85\xf6\x40\x29\x2b\x34\x84\xf3\xd3\xa1\x21\x3c\x90\x0e\x8d\x17\x53\x02\x0d\x84\xa4\x54\xde\xa7\x93\x21\x4d\xde\xc1\x94\x79\x93\x21\x64\x2e\x5e\x92\x0d\x3c\xcc\x13\x57\x74\x66\x0d\x3c\x04\x50\xc4\x06\xa0\xec\x98\x31\xba\x78\x45\xdc\x90\x1c\x4f\xf8\xfc\x39\x59\xd5\x15\xb9\x47\x3a\xa6\x35\x60\x90\x8d\xe1\x68\xcd\x4e\x3a\x62\x38\x07\x78\x62\x8b\x15\xe6\x8a\x60\x61\xbb\x38\x4f\x98\xdf\xe5\xac\x81\x6a\x0b\xbc\x42\xf6\xef\x50\x2c\x5d\x97\x18\x11\xa3\x74\x76\x40\x2f\x68\x0b\xbc\x24\xe8\x83\x97\xd4\xbd\xa0\xc3\xb1\x54\x1d\x8c\xc4\xca\xc2\x52\x18\x1f\x17\x08\xea\x28\x2d\x70\xbc\x47\x48\x0d\x79\xc0\xa7\x26\x63\xb5\xc7\x99\x6b\xc6\xb5\x0e\xa8\x0c\xb9\xc0\x0b\xda\x69\x69\xaa\x97\xe6\xae\xab\xe4\x38\xde\xe0\x9a\x56\x4f\xc8\xf2\x40\xa5\x9b\x1d\x4a\xad\xbb\xc3\x09\x7d\x87\xc8\x7d\x64\xf3\xf7\x0f\x44\x6b\x22\xb4\xf8\x00\xf7\x08\x8b\xf9\xb8\xcc\x0b\xa6\x24\xe7\xc0\x54\xd2\x60\x15\x74\xa4\xb7\x42\xf3\x80\x2d\xd9\xb2\xf7\x09\x12\x48\x53\x5b\x9a\xf6\x10\x93\x2f\xcc\x1a\x52\x70\x84\xa3\xee\xc8\xd6\x57\x53\x15\xa3\x0d\xba\x64\xd4\xc8\x62\xd3\x27\x5e\x10\x63\xa8\x72\xfc\xe2\x75\x9d\x53\x62\x14\xa7\x9e\x1a\x10\x16\x5e\x92\x8e\x7b\x8e\x36\xe2\x92\x65\x8b\xdb\x17\xd3\x1c\x5d\xc0\xa9\x56\x12\xa9\x6a\xd9\xaa\xfd\x4b\x52\x38\x2b\xe9\x5a\x78\xdf\x0e\x25\x73\x56\x3a\xc5\x9e\x91\x6c\x54\xd0\x5d\xd7\x36\x7c\x4d\x35\x85\xc0\xbf\x24\x7b\x54\x13\xbe\xba\x9c\x99\x04\x5c\x0e\x0c\x20\x84\xd5\x28\x74\xd1\xc1\x29\x80\x26\x5f\x3d\x34\x11\x9d\x03\x68\x4b\x0c\xab\xdb\x5d\xd6\xc7\x94\x34\x48\x6c\xeb\x8c\x65\x6f\x52\xc2\xe9\x51\x28\x7f\xd6\xe2\x34\x1c\xf1\x3e\x87\x58\x53\xb7\xf9\x05\x8e\xef\x0b\x2d\x43\xdc\x56\x17\x38\xd9\x83\x9e\xa1\x34\x3e\x5d\x57\x09\xe5\xba\x4a\x28\xd5\x55\x02\x5d\x3c\x09\xe5\xc2\x47\x80\xe8\xa9\x2a\xd3\x02\x27\x27\xa7\x2b\x22\xa1\x5c\x11\x09\xa5\x8a\x48\x28\x57\x44\x5a\xaa\xd4\xa5\xd8\x83\xac\xa1\x84\xfe\x1a\x4a\x20\x4b\x25\xa1\x58\x0f\x09\xfd\xf5\x90\x96\x96\x15\xe3\x55\xe5\x89\xb7\xb2\x40\x16\x3f\x42\xa1\xf8\x51\xba\xd9\x85\xb1\xc6\xb1\x33\x8c\xc9\x43\x3e\xd6\x15\xd5\x94\x3a\xf9\xe3\x14\x2b\xe2\x32\x9d\x03\xa6\x8a\x21\x72\x9c\x72\x77\xa9\x27\x7e\x95\x58\xe0\xf8\x1c\x97\x63\x68\x31\x42\x66\x43\x10\xd2\xb6\x28\x95\x0c\x74\x8c\xe3\x38\x23\x08\x4e\x49\x81\xbd\x0e\x91\xe8\xd3\xb7\x64\xba\x68\x5d\xaa\x59\x67\x1c\x47\xba\x99\xa6\x52\xc1\x39\x4c\x08\x42\xdc\x7f\x32\xd6\xf3\xae\x96\x69\x42\xca\x59\xa0\x62\x7f\xc4\x8f\xe8\xb1\xa1\x0a\x2d\xd4\x5b\x64\xeb\x86\x0a\x45\x80\x8e\xab\x40\x41\x89\x6b\x20\xda\x29\x78\x9f\x90\xc5\x97\x7c\x79\xd2\xe2\x03\x68\x47\x87\xbe\xbe\xa1\xb7\x6f\xb1\xac\xb1\x60\xfb\xe4\x7a\xa2\xcb\xa3\xf7\xc9\xf6\x44\xd3\xfd\x91\x59\x23\x69\xab\xc2\xbb\x48\xf9\x59\x24\x15\x4b\xd0\x89\x3a\xab\x89\x4a\xea\xe9\x7a\x41\x87\xe2\x8d\x9d\x28\x1d\xb4\xa0\x21\x1e\x5b\x96\x44\x21\xcb\x48\xa4\x72\xa6\xd4\xa3\xd7\x54\xd3\x6e\x20\x67\xd0\x68\xdc\x0c\x62\x7f\x9d\x51\xe2\x89\xa1\xbd\x2b\x92\xe0\xf4\x7f\xf4\x25\x92\x96\xe6\x57\xf0\x06\x54\x1e\xe1\xe8\x2f\x27\x2b\xd3\x3e\xe7\x2b\xa3\xec\x9a\x21\x04\x36\xd5\xac\x70\x47\x5f\xf2\x85\xeb\x72\xf4\x5c\xbc\xec\x77\x2b\xdc\x37\x58\xd2\xf5\x65\x77\x3a\xa9\xe0\x71\xf8\x66\xb0\xf7\x1a\x51\xd6\x2d\xa7\xc8\xff\x9d\xff\x73\x88\xdb\x4c\xb9\x4a\x90\xbb\xaf\x31\x5f\x56\x16\xf2\x1c\xf5\x81\x0f\xd6\x70\x95\x33\xad\xff\x5f\x15\x9c\xfc\x54\xc1\xfe\xe1\xfe\xde\x2f\x07\x3f\xf5\x90\xf5\x70\xf2\xf3\x48\x0c\x8e\x0e\x4e\x8f\x4e\x8f\x7f\x39\x38\xfd\xeb\x2e\xea\xa2\x1c\x0e\xb8\x17\x63\x42\x92\x56\x18\x27\x8e\x8f\x14\x6c\xf5\x51\xea\x46\x1e\x8a\xaf\x2f\x0f\xdc\x09\x6b\x22\x3c\xc6\x1f\xbf\x60\xfc\x2d\xe2\x1c\x5d\xfb\x6c\xfc\xdb\x95\xc0\xbd\x66\x7d\x1f\x31\x65\xfe\x19\x09\x7b\x9f\x1e\xef\x3c\xf0\x89\xb3\x12\x7f\x3f\x2a\xb8\x18\x43\xad\x52\x40\xaf\xd4\xed\x3d\x54\x85\x1d\xd0\xbb\xc6\x4f\x00\x19\x7f\x5c\xfb\x71\xc1\x9f\xd1\xbb\x18\xb3\xee\xd9\x61\xa7\xd7\x48\xa0\x27\xc7\x5b\xad\x5c\xdb\x4a\xd6\x33\xa6\xaa\x1d\x38\x55\xe1\xf0\x74\x8f\x58\x62\x86\x1f\xb7\xf3\x2d\x87\xa7\x7b\xcc\x09\xc9\xb4\x96\x96\x51\xcf\x2d\xeb\x2d\xc8\xdc\xbb\x6d\xe2\x93\xc9\x97\x03\x26\x4d\x88\x7e\x17\x7e\x11\xf0\x76\x74\xc6\x6e\xce\x47\x6f\xd9\xed\xf9\xe8\xe6\xec\xdd\x56\x6a\x6a\xbf\x84\xbf\x9b\xad\x7c\x2a\xbf\x03\x0b\xfc\xf9\x93\x87\x9f\x3f\x79\xd8\x85\x9f\x3c\xbc\x78\x92\x78\x35\xba\xfd\xf5\xfc\x6d\x7f\xae\x78\xb0\xbf\x7f\xb4\x77\xbc\x77\x70\xf2\xbc\x7b\x14\x3c\x7e\x5f\xaf\x8b\x12\x86\xef\x7e\x6e\xbb\x43\xfa\x03\xed\xb8\xff\x23\xec\x28\xac\xd6\x60\x62\x9e\x26\x9b\x61\xd0\xbe\xf8\x41\x04\x3f\xb0\xbe\x02\x2f\x4d\x33\xa8\xad\x1f\xb4\xca\x19\xc8\x30\xa8\x64\x5d\x83\x87\xf5\xf3\xd1\xb3\x1b\x5e\xd2\x70\xbb\x15\xe9\x5f\x7e\xbf\x6c\x11\xb6\xfe\x76\x7d\xfd\x71\xab\xd5\x78\xb8\xb3\x76\x67\xa2\xf0\xd9\xbb\x9b\x6d\x17\x24\xc6\x7e\x87\x16\x74\x75\xfd\xf6\x5f\x97\xe7\x5b\x2d\xa8\x7b\x92\x60\xda\x56\x09\xbf\x40\xe7\x1c\xa2\x40\xd5\xab\xdc\x8e\xa8\xe3\xc3\xe8\xec\x6c\x3b\xf3\x72\x21\x76\xc6\xb8\x1f\x3e\xde\x8c\xce\xb6\x33\xee\x44\xe0\x6f\x22\xba\x4f\xe5\x1a\x20\x7e\x3d\xb5\x28\x68\x17\x3e\x7f\xdd\xc5\x7a\xf8\x16\xfa\xbd\x19\xfd\x76\x71\xbd\x95\x7a\xa5\x75\xc4\x17\xee\x0e\xbc\xde\x95\x35\x7e\xbc\xb8\xda\x6e\x07\x85\x9e\x9f\x2e\x84\xf2\x0f\x32\x17\x8f\x7a\xbb\xa2\x86\x8f\x7f\xb0\xb3\xeb\xf7\x7f\xbf\xf8\xc7\x56\xca\x98\x8e\xb9\x69\x92\xdb\x95\x55\xbd\xbf\xd8\xd2\x3d\x34\x10\x99\x06\xed\xac\x92\x02\x7f\x8d\x4d\xfd\x36\x39\xac\xf5\xd8\x0d\x3d\x5c\x5e\x6f\x67\xd5\xf9\x75\x8d\x58\xcc\xab\xf6\x1f\x5f\x5e\xfd\x37\x00\x00\xff\xff\xc4\xf4\x76\x8a\x17\x47\x00\x00")

func fuseContainerJsonBytes() ([]byte, error) {
	return bindataRead(
		_fuseContainerJson,
		"fuse-container.json",
	)
}

func fuseContainerJson() (*asset, error) {
	bytes, err := fuseContainerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fuse-container.json", size: 18199, mode: os.FileMode(0664), modTime: time.Unix(1648072800, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9c, 0x5d, 0x2a, 0x60, 0xef, 0x88, 0x32, 0xe, 0xf7, 0xe3, 0x41, 0x1f, 0xa4, 0x3b, 0x81, 0xd7, 0xc9, 0xf, 0x26, 0x74, 0xb7, 0xe1, 0xb6, 0xf4, 0x86, 0xd9, 0x3f, 0x31, 0x27, 0x97, 0xe0, 0xfd}}
	return a, nil
}

var _allowPerfSyscallsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5b\x6f\x5b\xb9\x11\x7e\xdf\x5f\x21\xe8\x79\x51\xf8\x56\xaf\x9d\x37\x55\x71\xbb\xc6\xda\x71\x6a\xa7\xd8\x5d\x14\x01\x41\xf3\xcc\x39\x62\xc5\x9b\x79\x91\x2d\x04\xf9\xef\x05\xcf\x91\x64\x49\x33\x54\xbc\x7b\x9c\x46\x05\xf2\xb0\x9b\xe4\xfb\x78\x19\xce\x90\x33\xc3\xe1\xd1\xa7\x1f\x06\x83\xc1\x90\x7b\x31\xb9\xe6\x6e\xf8\x66\xf0\xef\xfc\xef\xc1\x60\xf0\x69\xf1\xe7\x92\x95\x11\x44\x4c\x1e\x86\x6f\x06\xc3\xbb\xf1\xf5\x7b\x36\xba\x1d\xff\xcc\x7e\x3b\x3b\x65\xa7\x27\xc3\x1f\xd7\x1a\x87\x74\x3f\x5a\x6b\x1f\xd6\xc6\x5c\xb4\xd8\xe8\xbe\xde\x77\x8b\x3c\x3e\x1a\x3e\x73\x1f\x17\x7f\xfd\xfc\xe3\x1f\x12\x70\x94\xff\xdf\x43\xc2\xd1\xed\x75\x7f\x21\xae\x2f\xdf\xdf\xf5\x90\x21\x77\x2f\xab\xa9\x1b\xfc\xdd\x6b\x28\xeb\x79\xa8\xaf\x29\xea\xeb\xc8\x79\x71\xd5\x53\xa3\x17\x57\xbb\x05\xcd\x13\xbc\x9a\x56\x97\x83\x7d\x6d\x81\xfb\x4b\x7b\x77\x7c\x7e\xf0\xdb\x9f\x96\x33\xf7\xa6\x64\xc8\x7f\x7e\x6c\x07\x1d\x56\x50\xf3\xa4\xe2\x48\x44\x69\xcd\xf3\xdc\xe3\x0f\xec\xe2\xf6\xf6\xdd\xcd\x70\xa3\xd5\x85\xf7\xc6\xde\x42\x1c\xbe\x19\x1c\x76\x44\x98\x07\xc1\x95\x0a\x05\x57\x25\x8d\x50\xa9\x6a\x85\xfc\xb4\x29\xa4\xe0\x0e\x8b\x9e\x89\xf1\xe8\x3d\xbb\xfb\xfd\x8e\x8d\xde\x5e\x5f\xbe\x1b\x6e\xd0\x1f\x9f\xff\xf5\x79\x5d\x25\x86\x6b\x4a\x0f\xf7\xae\xde\x36\x91\x03\x5f\x33\x98\x81\x89\xcc\x3a\x30\x2f\xb6\x0f\xd6\xce\xe8\xea\xea\xe6\xd7\x0d\xc3\xec\xdb\x5a\x85\xb2\x06\xb6\xc1\x9a\x1b\x1b\x65\x3d\x67\xd2\xc8\x88\xc8\x20\xac\xa9\x65\x83\x71\x6d\x93\x21\x9a\xb7\x3a\x44\xa8\x93\x62\xba\x8d\x2a\x6b\xa7\xc9\xb1\x4a\x58\x3b\x95\x48\x2a\x72\x78\x6d\x67\xc0\x48\x26\xeb\x80\x45\xcb\x26\xdc\x54\x0a\x18\x47\x0d\xb2\x5c\x2c\x7a\x40\x13\x6d\xdb\x7f\x8b\x7e\x48\x36\x72\x11\xd5\x36\x1e\x20\x56\x56\x73\x69\xf2\xcc\x04\x39\xb1\x21\x16\x28\x13\x10\x38\x0f\xca\x22\x1d\x27\x72\xa5\x1d\x7a\x84\x60\x13\x26\xdc\xc3\xab\x6e\xdf\xc2\xc6\xe2\x42\x80\x43\x72\x75\xe8\x09\x05\x07\xb4\x60\x5e\xfd\x27\x4a\x0d\x4f\x08\x57\xdc\xeb\x6d\xf0\x5e\x9a\x0a\x61\x1e\xed\x27\xc1\x5d\x03\x48\xac\x7c\xd0\x08\x74\x52\x49\x8f\x41\x6d\xd1\x44\x62\x62\x1f\xd1\xa6\x68\xc1\x63\x64\x03\xa1\xac\x98\xb2\xc5\xda\x76\x92\xa7\x48\x4f\x1d\xdd\x40\xcc\xfe\x7b\x07\xc7\xbe\xd0\xbd\x3c\xf5\x82\x2c\xf5\x35\xdc\xd8\xa0\x00\xdc\x17\xe8\xb2\x00\x81\x9a\x38\x00\xf3\xdc\x34\x98\xb2\xc6\x80\xc0\x86\xb1\x6e\xce\x6a\xa9\x4a\xbd\x3c\xe0\xa3\x5d\x25\x24\x73\x95\x1c\xb2\x4e\x95\xdc\xf1\x36\x06\xce\x2a\xc5\xda\x51\xd1\x5c\xeb\xdc\x61\x81\xc4\x7e\x61\x45\x30\xab\xd0\x66\xea\x48\xf7\xc8\xb1\xa3\x5d\xa3\x90\xe4\x1d\x57\xee\x95\x19\x72\xba\xec\xd5\xea\x12\x8c\xa7\x79\x02\x31\xc3\x6a\x68\x51\xac\x75\x78\x22\xc4\x79\x92\x91\x35\xde\x62\x83\xd4\x9d\x1f\xc0\xc3\xac\x08\x24\x4e\xcd\xab\x99\x0c\xc4\x56\x5b\x11\x8c\xe2\x94\xb2\x82\xb0\xe6\x2a\xca\x69\x8e\x5d\x47\x4d\xfa\x83\x9a\x74\x08\x1d\x4a\x2c\x84\xf4\x14\x75\xc1\x55\x74\x38\x35\x8a\xc1\x7b\xaa\x05\x89\xb5\x56\x3c\xf2\x30\x37\x02\x11\x0d\xc4\x27\x1e\x23\x5e\x91\x92\xa1\xc4\x58\x1c\xa1\x6b\x4b\xa8\xca\x43\x8e\xc2\xf4\x20\xa1\x34\x6f\x88\xc4\x52\x33\x48\xac\x2a\xc3\x45\xa2\x46\xce\x71\x01\x53\xed\x29\xd5\x44\x9f\x0c\xb9\x3f\x96\x04\x31\x50\x8a\x38\x52\xb5\x60\xc1\x17\xd6\x29\xe3\xc4\x66\x6f\x20\x0a\x97\x28\xf4\x11\xed\xb3\x06\x62\x05\x26\xa2\xf5\x2e\x71\x3c\x6d\x03\x11\x1a\x49\x0d\x94\x61\xbc\x07\x33\x91\xe8\xf6\xa9\xd0\x9e\x1e\xbe\x34\x7a\xeb\x09\xa8\x05\x74\x04\xd9\x47\x66\xd5\xa1\x0d\xd4\x40\x74\x00\x9e\x4a\xa8\x32\x45\x8b\xe5\x1a\x8f\xfc\x50\x86\xe9\xc6\x05\xd8\x4b\xeb\x65\x9c\x13\x94\xe7\xa6\xb2\x28\x59\xe9\xa2\x34\x2d\x50\x47\x90\xcb\xf6\x10\x68\x53\x74\x04\xdd\x47\x49\x8d\x7d\x70\x03\x91\x79\x7b\x9f\x42\x64\xf9\xb4\x53\xfd\x52\xe0\x38\xb4\x36\x10\x03\x29\x41\xb0\x62\x5a\x50\x7c\xa6\x2c\x4e\x05\xb3\x08\x71\xe2\x81\x57\x8c\x7b\xe0\x04\x1d\xc9\x99\xb2\xe9\x6d\x5d\x71\x4a\xdb\xb4\x76\x4a\xaa\x21\xbd\x90\x5c\x44\x00\x5e\x55\xec\x91\x47\x31\x29\x35\xa0\x2e\x42\xeb\x1c\xca\x06\x96\xa4\xd7\x85\x81\x2d\x13\xdc\x08\x40\x9e\x5d\x5a\x22\x85\x90\x96\x55\x10\xa2\xb7\x48\x0f\xd2\xe6\x2c\xae\x8d\xdf\xe8\x58\x49\xcb\xdc\xcb\xc8\x82\xd7\x92\x36\x6f\x76\x46\xe4\xd0\x0b\x86\xc8\xa3\x3b\x14\x87\xfb\x8c\xa7\x7b\x62\x77\x4a\xcb\x92\x97\xa6\x61\x60\x22\x3e\xe6\x2b\xd6\x43\x23\xc3\xae\x06\xf4\xac\x0e\xf9\xfb\xa9\x54\x48\xbf\x8a\x8c\xd4\xaa\x10\xa9\x55\x69\x3b\x29\x69\xf0\x9d\x56\x9a\x29\x76\xfb\xf9\x1c\xe2\x9b\x65\x31\x16\xab\x22\xc3\x94\x0a\x00\x78\xd6\x1d\x31\x99\xee\x50\x0a\xd4\x8a\x0a\xd4\x8a\x0e\xd4\xba\x4b\xc3\x10\x0c\xfa\x9e\x7b\x2f\xb1\xf1\x34\xe8\xba\x2a\x24\xdb\x5a\x1a\x61\x3d\x86\xa7\x44\x5e\xd6\x82\x58\x4c\x3d\x35\x38\x5b\x6b\x41\xa2\x2d\x95\xee\xb4\x20\xb2\x7f\x8b\x72\xbc\x8b\xb4\xe6\x68\x03\x66\x0c\x0f\xe0\xbc\x8d\xc4\x55\x47\x3f\xe4\xb3\x16\x20\x52\x86\xd0\x0f\xac\xf3\x29\x04\x41\x95\x29\xf4\x43\x7b\xaa\x2b\x0f\x02\x24\x4e\xe2\xb7\xe8\x82\x07\x58\xb6\x0a\x80\x2f\xdc\xeb\x5c\xb9\x7b\x32\xd4\xa9\xd0\x1e\x28\x65\x85\x86\x70\x7e\x3a\x34\x84\x07\xd2\xa1\xf1\x62\x46\xa0\x81\x90\x94\xca\xfb\x74\x32\xa4\xc9\x3b\x98\x32\x6f\x32\x84\xcc\xc5\x4b\xb2\x81\xc7\x45\xe2\x8a\xce\xac\x81\xc7\x00\x8a\xd8\x00\x94\x1d\x33\x46\x17\xaf\x88\x1b\x92\xe3\x09\x9f\x3f\x27\xab\xba\x22\xf7\x48\xc7\xb4\x06\x0c\xb2\x31\x1c\xad\xd9\x49\x47\x0c\xe7\x00\x4f\x6c\xb1\xc2\x5c\x11\x2c\x6c\x17\xe7\x09\xf3\xbb\x9c\x35\x50\x6d\x81\x57\xc8\xfe\x1d\x8a\xa5\xeb\x12\x23\x62\x94\xce\x0e\xe8\x45\x6b\x89\x97\x04\x7d\xf4\x92\xba\x17\x74\x38\x96\xaa\x83\x91\x58\x59\x58\x0a\xe3\x93\x02\x41\x1d\xa5\x25\x8e\xf7\x08\xa9\x21\x0f\xf8\xd4\x64\xac\xf6\x38\x73\xcd\xb8\xd6\x01\x95\x21\x97\x78\x41\x3b\x2d\x4d\xf5\xd2\xdc\x75\x95\x1c\xc7\x1b\x5c\xd3\xda\x11\xb2\x3c\x50\xe9\x66\x87\x52\xeb\xee\x70\x42\xdf\x21\x72\x1f\xd9\xe2\x3d\x02\xd1\x9a\x08\x2d\x3e\xc0\x03\xc2\x62\x3e\x2e\x8b\x82\x29\xc9\x39\x30\x95\x34\x58\x05\x1d\xe9\xad\xd0\x3c\x60\x4b\xb6\xec\x43\x82\x04\xd2\xd4\x96\xa6\x3d\xc4\xe4\x0b\xb3\x86\x14\x1c\xe1\xa8\x3b\xb2\xf5\xd5\x54\xc5\x68\x8b\x2e\x19\x35\xb2\xd8\xec\x12\x2f\x88\x09\x54\x39\x7e\xf1\xba\xce\x29\x31\x8a\x53\xcf\x0d\x08\x0b\xaf\x48\xc7\x3d\x47\x1b\x71\xc5\xb2\xe5\xed\x8b\x69\x8e\x2e\xe0\x54\x2b\x89\x54\xb5\x6a\xd5\xfe\x25\x29\x9c\x95\x74\x2d\xbc\x6f\x87\x92\x39\x2b\x9d\x61\xcf\x48\x36\x2a\xe8\xae\x6b\x1b\xbe\xa4\x9a\x42\xe0\x5f\x91\x3b\x54\x13\xbe\xb8\x9c\xb9\x04\x5c\x0e\x0c\x20\x84\xd5\x28\x74\xd1\xc1\x29\x80\x26\x5f\x3d\x34\x11\x9d\x03\x68\x4b\x0c\xab\xdb\x5d\xb6\x8b\x29\x69\x90\xd8\xd6\x19\xcb\xde\xa4\x84\xd3\xa3\x50\xfe\xac\xc5\x69\x38\xe2\x7d\x0e\xb1\xa6\x6e\xf3\x4b\x1c\xdf\x17\x5a\x86\xb8\xad\x2e\x71\xb2\x07\x3d\x43\x69\x7c\xba\xae\x12\xca\x75\x95\x50\xaa\xab\x04\xba\x78\x12\xca\x85\x8f\x00\xd1\x53\x55\xa6\x25\x4e\x4e\x4e\x57\x44\x42\xb9\x22\x12\x4a\x15\x91\x50\xae\x88\xb4\x54\xa9\x4b\xb1\x07\x59\x43\x09\xbb\x6b\x28\x81\x2c\x95\x84\x62\x3d\x24\xec\xae\x87\xb4\xb4\xac\x18\xaf\x2a\x4f\xbc\x95\x05\xb2\xf8\x11\x0a\xc5\x8f\xd2\xcd\x2e\x4c\x34\x8e\x9d\x61\x42\x1e\xf2\x89\xae\xa8\xa6\xd4\xc9\x9f\xa4\x58\x11\x97\xe9\x1c\x30\x55\x0c\x91\xe3\x94\xbb\x4b\x3d\xf1\xab\xc4\x12\xc7\xe7\xb8\x1c\x43\x8b\x11\x32\x1b\x82\x90\xb6\x45\xa9\x64\xa0\x63\x1c\xc7\x19\x41\x70\x4a\x0a\xec\x75\x88\x44\x9f\xbe\x25\xd3\x45\xeb\x52\xcd\x3a\xe3\x38\xd2\xcd\x35\x95\x0a\x2e\x60\x42\x10\xe2\xfe\x93\xb1\x1d\xef\x6a\x99\x26\xa4\x9c\x07\x2a\xf6\x47\xfc\x88\x1e\x1b\xaa\xd0\x42\xbd\x45\xb6\x6e\xa8\x50\x04\xe8\xb8\x0a\x14\x94\xb8\x06\xa2\x9d\x81\xf7\x09\x59\x7c\xc5\x97\x27\x2d\x3e\x80\x76\x74\xd8\xd5\x37\xec\xec\x5b\x2c\x6b\x2c\xd9\x5d\x72\x3d\xd3\xe5\xd1\x77\xc9\xf6\x4c\xd3\xfd\x91\x59\x23\x69\xab\xc2\xbb\x48\xf9\x59\x24\x15\x4b\xd0\x89\x3a\xab\x89\x4a\xea\xe9\x7a\x41\x87\xe2\x8d\x9d\x28\x1d\xb4\xa0\x21\x1e\x5b\x56\x44\x21\xcb\x48\xa4\x72\x66\xd4\xa3\xd7\x4c\xd3\x6e\x20\x67\xd0\x68\xdc\x0c\x62\x7f\x9d\x51\xe2\x89\xa1\xbd\x2b\x92\xe0\xec\x7f\xf4\x25\x92\x96\xe6\x17\xf0\x06\x54\x1e\xe1\xe4\x2f\x67\x6b\xd3\xbe\xe4\x2b\xa3\xec\x9a\x21\x04\x36\xd3\xac\x70\x47\x5f\xf1\x85\xeb\x72\xf4\x5c\xbc\xee\x77\x2b\xdc\x37\x58\xd2\xcd\x65\x77\x3a\xa9\xe0\x69\xf8\x66\x70\xf0\x23\xa2\xac\x5b\x4d\x91\xff\xbb\xf8\xe7\x10\xb7\x99\x71\x95\x20\x77\xdf\x60\x3e\xaf\x2d\xe4\x25\xea\x03\x1f\xac\xe1\x2a\x67\x5a\xff\xbf\x2a\x38\xfb\xae\x82\xc3\xe3\xc3\x83\x9f\x8e\xbe\xeb\x21\xeb\xe1\xec\xfb\x91\x18\x9c\x1c\x9d\x9f\x9c\x9f\xfe\x74\x74\xfe\xd7\x7d\xd4\x45\x39\x1c\x70\x2f\x26\x84\x24\xad\x30\x4e\x9c\x9e\x28\xe8\xf5\x51\xea\x56\x1e\x8a\xaf\x2f\x8f\xdc\x09\x6b\x22\x3c\xc5\x6f\xbf\x60\xfc\x2d\xe2\x02\xdd\xf8\x8c\xfb\x8f\x2b\x81\x7b\xcd\x76\x7d\xc4\x94\xf9\x17\x24\xec\xbb\xf4\x78\xef\x81\x4f\x9d\x95\xf8\xfb\x51\xc1\xc5\x04\x6a\x95\x02\x7a\xa5\x6e\xef\xa1\x2a\xec\x81\xde\x35\x7e\x02\xc8\xf8\xd3\xc6\xc7\xfe\x7f\x46\xef\x62\xc2\xba\x67\x87\xbd\x5e\x23\x81\x9e\x9d\xf6\x5a\xb9\xb6\x95\xac\xe7\x4c\x55\x7b\x70\xaa\xc2\xf1\xf9\x01\xb1\xc4\x0c\x3f\xf5\xf3\x2d\xc7\xe7\x07\xcc\x09\xc9\xb4\x96\x96\x51\xcf\x2d\x9b\x2d\xc8\xdc\xbb\x6d\xe2\x93\xc9\x97\x03\x26\x4d\x88\x7e\x1f\x7e\x11\xf0\x76\x34\x66\xb7\x17\xa3\xb7\xec\xee\x62\x74\x3b\xfe\xb9\x97\x9a\xda\x2f\xe1\xef\xe7\x6b\x9f\xca\xef\xc1\x02\xbf\xff\xe4\xe1\xfb\x4f\x1e\xf6\xe1\x27\x0f\xaf\x9e\x24\x5e\x8f\xee\x7e\xb9\x78\xbb\x3b\x57\x3c\x3a\x3c\x3c\x39\x38\x3d\x38\x3a\x7b\xd9\x3d\x0a\x9e\xbe\xae\xd7\x45\x09\xc3\x57\x3f\xb7\xdd\x21\xfd\x86\x76\x3c\xfc\x16\x76\x14\x56\x6b\x30\x31\x4f\x93\xcd\x30\x68\x5f\xfc\x20\x82\x1f\x58\x5f\x81\x97\xa6\x19\xd4\xd6\x0f\x5a\xe5\x0c\x64\x18\x54\xb2\xae\xc1\xc3\xe6\xf9\xd8\xb1\x1b\x5e\xd3\x70\xfb\x15\xe9\x5f\x7f\xbf\xf4\x08\x5b\x7f\xbb\xb9\xf9\xd0\x6b\x35\x1e\xee\xad\xdd\x9b\x28\x3c\xfe\xf9\xb6\xef\x82\xc4\xc4\xef\xd1\x82\xae\x6f\xde\xfe\xeb\xea\xa2\xd7\x82\xba\x27\x09\xa6\x6d\x95\xf0\x0b\x74\xce\x21\x0a\x54\xbd\xce\xed\x89\x3a\xde\x8f\xc6\xe3\x7e\xe6\xe5\x42\xec\x8d\x71\xdf\x7f\xb8\x1d\x8d\xfb\x19\x77\x2a\xf0\x37\x11\xdd\xa7\x72\x0d\x10\xbf\x9e\x5a\x16\xb4\x0b\x9f\xbf\xee\x63\x3d\xbc\x87\x7e\x6f\x47\xbf\x5e\xde\xf4\x52\xaf\xb4\x8e\xf8\xc2\xdd\x81\xd7\xfb\xb2\xc6\x0f\x97\xd7\xfd\x76\x50\xd8\xf1\xd3\x85\x50\xfe\x41\xe6\xf2\x51\x6f\x5f\xd4\xf0\xe1\x77\x36\xbe\x79\xf7\xf7\xcb\x7f\xf4\x52\xc6\x6c\xc2\x4d\x93\xdc\xbe\xac\xea\xdd\x65\x4f\xf7\xd0\x40\x64\x1a\xb4\xb3\x4a\x0a\xfc\x35\x36\xf5\xdb\xe4\xb0\xd1\x63\x3f\xf4\x70\x75\xd3\xcf\xaa\x8b\xeb\x1a\xb1\x98\x1f\xda\x7f\x7c\xfe\xe1\xbf\x01\x00\x00\xff\xff\xbe\xd7\xf7\x81\xa7\x46\x00\x00")

func allowPerfSyscallsJsonBytes() ([]byte, error) {
	return bindataRead(
		_allowPerfSyscallsJson,
		"allow-perf-syscalls.json",
	)
}

func allowPerfSyscallsJson() (*asset, error) {
	bytes, err := allowPerfSyscallsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "allow-perf-syscalls.json", size: 18087, mode: os.FileMode(0664), modTime: time.Unix(1648072800, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xce, 0x1f, 0xce, 0x25, 0xe6, 0x76, 0xc4, 0xe4, 0xb2, 0x86, 0xc2, 0x1e, 0x57, 0xaa, 0xba, 0x84, 0xcd, 0x70, 0x75, 0xcd, 0xa2, 0x38, 0xb4, 0xb3, 0x5a, 0xe4, 0x6, 0x4d, 0x55, 0x17, 0xe2, 0xd7}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"default.json":             defaultJson,
	"fuse-container.json":      fuseContainerJson,
	"allow-perf-syscalls.json": allowPerfSyscallsJson,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"allow-perf-syscalls.json": &bintree{allowPerfSyscallsJson, map[string]*bintree{}},
	"default.json":             &bintree{defaultJson, map[string]*bintree{}},
	"fuse-container.json":      &bintree{fuseContainerJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
