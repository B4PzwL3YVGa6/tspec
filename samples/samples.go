// Code generated by go-bindata.
// sources:
// source/array_type.json
// source/basic_types.json
// source/map_type.json
// source/normal_struct.json
// source/struct_with_anonymous_field.json
// source/struct_with_circular_reference.json
// source/struct_with_inheritance.json
// source/struct_with_no_export_field.json
// DO NOT EDIT!

package samples

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

var _sourceArray_typeJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\xe2\x54\x72\x2c\x2a\x4a\xac\x0c\xa9\x2c\x48\x55\xb2\x52\xa8\xe6\xe2\xe4\x54\x2a\x81\xb0\x95\x12\x41\x12\x4a\x3a\x20\xa1\xcc\x92\xd4\xdc\x62\xa8\x3c\x42\x41\x66\x5e\x49\x6a\x7a\x6a\x11\x58\x09\xa7\x52\x5a\x7e\x51\x6e\x62\x09\x54\x5c\x89\x8b\x93\xb3\x96\x8b\xb3\x96\xab\x96\x0b\x10\x00\x00\xff\xff\xa4\x08\x73\xcd\x64\x00\x00\x00")

func sourceArray_typeJsonBytes() ([]byte, error) {
	return bindataRead(
		_sourceArray_typeJson,
		"source/array_type.json",
	)
}

func sourceArray_typeJson() (*asset, error) {
	bytes, err := sourceArray_typeJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "source/array_type.json", size: 100, mode: os.FileMode(420), modTime: time.Unix(1503379971, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sourceBasic_typesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x94\xcf\x8b\xea\x30\x10\xc7\xcf\xed\x5f\x21\x39\xfb\x0e\xfe\xa0\x88\xc7\xf7\x40\xf0\xfa\xd6\x9e\x97\x56\xa3\xcc\x92\x26\x25\xce\xc0\x16\xf1\x7f\x5f\x9a\xa2\xeb\x32\x93\xb2\x06\x0f\xbd\x4c\x32\x9f\x4f\x32\x7c\xd3\x4b\x9e\xa9\xbf\xd5\x19\xf6\xbb\xae\xd5\x67\xb5\x9e\x5c\xf2\x2c\x53\xd8\xb5\x5a\xad\x27\xca\xd5\x1f\x7a\x8f\x6a\x1a\x6a\x80\x26\x14\x1f\xb6\x87\x85\xd6\xbb\x56\x7b\x84\x7b\x7b\xa6\x6a\xe7\xcc\xfb\x11\xb4\x39\xdc\x4a\xdf\xcc\x7e\x4d\x57\x36\xf4\xfe\xc0\x3a\x67\x36\xa1\xa5\x5f\xb8\x4e\x07\x50\x87\x3a\x06\x3a\xa3\x07\x7b\xba\x71\x8e\xce\x37\x15\x06\x41\x87\x9a\xd3\x3b\xd4\x8c\xbe\x77\x4d\x6b\xf4\xe7\x6c\xbe\x8a\x39\x2c\x35\xb5\xf6\xdc\x71\x70\x54\x1b\x6e\xf9\x77\x07\xc6\x5c\xc5\xf2\x59\xd5\xd1\xb8\x0a\x63\xa6\x62\xc9\x44\x61\xff\x62\x9e\xa4\x59\xcc\x99\x68\x33\xd4\x65\x4d\xe2\x6d\x8a\xa5\xac\x11\x6e\x03\x16\x67\x45\x4c\x02\x16\xf5\x49\xb2\x84\x2e\xe6\xd8\xf6\x55\xc9\x10\x9f\xd6\x98\x41\x18\xd6\xd6\x4a\xa3\x02\x3b\x32\xa8\x31\x83\x30\xa7\xad\x8d\x4c\x29\x1a\xe1\x31\xc1\x4a\xe2\xf3\xec\x82\xc5\x14\xba\x04\x67\x6c\x4f\xf6\x35\x2f\xfc\x3f\x59\xfe\xc2\x07\xc2\xef\xf8\x77\xd2\x5b\x28\x33\x16\x42\xf3\xf4\x49\x0f\x15\xea\x3f\x7d\x23\x93\xec\xa0\xe1\xc7\xa5\xb4\xb8\x93\x9c\xf7\x12\xc4\xc0\x53\x5a\xe2\x49\x8e\x7c\x09\x62\xe6\x29\x2d\xf4\x24\xa7\xbe\x04\x31\xf6\x94\x94\x7b\x12\x83\xdf\x2b\x78\xf2\x29\x25\xfa\x24\x65\xbf\xe7\x8b\xf8\x16\xfd\x8b\xfe\x0d\xe5\x40\x7b\xb0\xe4\xe1\xbb\xe6\xd7\xfc\x2b\x00\x00\xff\xff\x3a\x0f\x46\xfa\xe8\x07\x00\x00")

func sourceBasic_typesJsonBytes() ([]byte, error) {
	return bindataRead(
		_sourceBasic_typesJson,
		"source/basic_types.json",
	)
}

func sourceBasic_typesJson() (*asset, error) {
	bytes, err := sourceBasic_typesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "source/basic_types.json", size: 2024, mode: os.FileMode(420), modTime: time.Unix(1504168910, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sourceMap_typeJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\xe2\x54\xf2\x4d\x2c\x08\xa9\x2c\x48\x55\xb2\x52\xa8\xe6\xe2\xe4\x54\x2a\x81\xb0\x95\xf2\x93\xb2\x52\x93\x4b\x94\x74\x40\x62\x89\x29\x29\x99\x25\x99\xf9\x79\x89\x39\x01\x45\xf9\x05\xa9\x45\x25\x99\xa9\xc5\x50\xf5\x08\x0d\x99\x79\x25\xa9\xe9\xa9\x45\x60\x1d\x9c\x4a\x69\xf9\x45\xb9\x89\x25\x50\x71\x25\x2e\x4e\xce\x5a\x2e\xce\x5a\xae\x5a\x2e\x40\x00\x00\x00\xff\xff\x14\x15\x20\xef\x72\x00\x00\x00")

func sourceMap_typeJsonBytes() ([]byte, error) {
	return bindataRead(
		_sourceMap_typeJson,
		"source/map_type.json",
	)
}

func sourceMap_typeJson() (*asset, error) {
	bytes, err := sourceMap_typeJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "source/map_type.json", size: 114, mode: os.FileMode(420), modTime: time.Unix(1503379952, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sourceNormal_structJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x96\xcf\x6a\xe3\x30\x10\x87\xcf\xf6\x53\x14\xed\x1e\xbb\x94\xa6\xc1\x94\x1e\xb7\x50\xc8\xa5\x87\x6d\x7d\x2e\xfe\x33\x2e\xb3\xc8\x92\x91\x47\xb0\xa6\xf4\xdd\x17\xc9\x9b\xac\x83\x46\x26\x51\x73\xc8\x65\xe4\xf9\xbe\x78\xf8\xc9\xd2\x47\x9e\x89\x9f\xd5\x88\xcd\xeb\x34\xc0\x28\x1e\xae\x3e\xf2\x2c\x13\x34\x0d\x20\x1e\xae\x84\xae\x7f\x43\x43\xe2\xda\xd7\x90\xa4\x2f\x2e\x1e\xf7\x0b\x83\xd1\x03\x18\xc2\x43\x7b\x26\x6a\xad\xe5\x5b\x87\x20\xdb\x7d\xe9\x3f\xd3\xad\x41\xa5\x7c\xef\x11\x56\x6b\xf9\xe4\x5b\xdc\xc2\xe7\xf5\x0c\x9a\x08\x62\xa0\x91\x0c\xaa\xf7\x3d\xa7\xd3\xa6\xaf\xc8\x0b\x26\x82\x90\x3e\x11\x04\xf4\x46\xf7\x83\x84\x3f\xb7\x9b\xfb\x98\x43\xd9\xbe\x06\x13\x3a\x5a\x6d\x6b\x19\x5a\x1e\x0f\xc0\x98\xab\xd8\x9e\xab\xea\xa4\xae\x28\x66\x2a\xb6\x81\xc8\x3f\x7f\xb7\x49\xd2\xdc\x6d\x02\xd1\xd3\x5c\xe7\x35\x89\x6f\x53\x6c\x79\x0d\xf3\x36\xa8\xe8\xb6\x88\x49\x50\x11\xbc\x73\x16\xdf\x15\x38\x76\xae\xca\x19\xe2\xd3\x5a\x33\x30\xc3\xda\x29\x6e\x54\xa8\x56\x06\xb5\x66\x60\xe6\xb4\x53\x91\x29\x45\x23\xbc\x26\xb8\xe7\xf8\x61\x76\x51\x51\x0a\x9d\x83\x07\x6c\x63\xd5\x65\x76\xf8\x2f\xab\xc2\x1d\x3e\x13\x4e\xe3\x1f\x48\x2f\xbe\x1c\xb0\x08\xfb\xb3\xff\x69\x5b\x11\xfc\x70\x8d\x81\xe4\x15\xfb\xf0\xef\xda\xb4\xb8\x5b\x3e\xef\x25\xb2\x81\xb7\x69\x89\xb7\x7c\xe4\x4b\x64\x33\x6f\xd3\x42\x6f\xf9\xd4\x97\xc8\xc6\xde\x26\xe5\xde\xb2\xc1\x77\x8a\x30\xf9\x36\x25\xfa\x96\xcb\xbe\xe3\xb3\xf8\x81\xcc\x85\xbe\x0d\xe5\x4c\x5b\x58\x72\xff\x73\x32\xf1\xec\x9a\xe5\x0b\x19\xdb\xd0\x49\xc7\xfc\x51\x43\xf4\xa0\x77\x77\x81\x37\x5a\xdc\x1d\x56\x6e\x0a\x59\x26\xbe\x1b\xe8\xdc\xca\xb7\x9b\x16\x3a\x54\x48\xa8\xd5\x78\xb3\x78\x6e\x79\x62\x1a\xa8\x08\x2e\xb1\xd7\x1e\x67\xd2\x82\xfd\xef\x78\xfa\xda\xe7\xec\x79\x86\x2c\x27\x9d\x7f\xe6\x7f\x03\x00\x00\xff\xff\x15\x3c\x3f\x3f\x52\x09\x00\x00")

func sourceNormal_structJsonBytes() ([]byte, error) {
	return bindataRead(
		_sourceNormal_structJson,
		"source/normal_struct.json",
	)
}

func sourceNormal_structJson() (*asset, error) {
	bytes, err := sourceNormal_structJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "source/normal_struct.json", size: 2386, mode: os.FileMode(420), modTime: time.Unix(1504168532, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sourceStruct_with_anonymous_fieldJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x93\xb1\x4e\xc3\x30\x10\x86\x67\xfb\x29\x2a\xc3\x18\xa9\x7b\xb7\x22\xc1\x56\x09\x29\x03\x63\xe4\x36\x0e\x18\x39\x3e\xcb\xb9\x0e\x51\x95\x77\x47\x71\xe2\x72\xe0\x58\x08\xc2\x02\x4b\x87\xbb\xeb\x77\x9f\x7e\x5f\x2e\x9c\x89\x12\xfd\xf9\x84\x4f\x1a\x5f\xf6\x16\x6c\xdf\xc2\xb9\x7b\xd0\xca\xd4\x62\xb7\xb9\x70\xc6\x04\xf6\x4e\x89\xdd\x46\xc0\xf1\x55\x9d\x50\x14\xa1\xa6\xd1\x84\x62\xf6\xcf\x61\xcc\x79\x70\xca\xa3\x56\xdd\x0c\x63\x42\xc6\xb1\x4a\x7a\x2f\xfb\x58\x7f\x5f\x33\x95\x8b\xb9\x18\xf7\x5c\xe9\x7b\xda\xd6\xa8\xda\x2b\x9a\x31\x71\xeb\x55\x33\x4e\xdf\x6c\x6b\xd5\x68\xab\x51\x83\xed\xb6\x39\xc7\xea\x23\xb4\xba\x37\x28\x02\x68\x18\x7f\x87\xe2\x93\x6f\x2b\x5d\x6a\x4b\x42\x59\xd4\x3d\x48\x17\x9b\xb2\xae\x83\x90\x34\x8f\x49\x2c\x2b\xdc\x0f\xd2\x7d\x65\xde\x05\x08\x91\x4f\x34\xa7\x35\xd1\xf4\xa7\x2a\x33\x85\xcf\x22\x03\x0f\x2a\xd9\x1b\x59\xca\x7f\xcd\xcd\x2d\xf1\x72\x67\x78\x04\x30\x55\x43\xce\x9c\x2c\x1d\x7b\x4a\xda\xe4\x51\xef\x00\xcc\x74\xdc\x24\xe5\x0e\xbd\xb6\xcf\x39\xd4\xd4\x4d\x48\x65\x28\x13\xd6\xb7\xd2\x8a\x2f\xfe\x3b\x59\x45\xda\x7f\x4c\xaa\xa4\x87\xbf\x36\x28\xf2\x8d\xfc\x9d\x9c\xf8\xc0\xdf\x02\x00\x00\xff\xff\x24\xd9\x67\xc7\xe4\x05\x00\x00")

func sourceStruct_with_anonymous_fieldJsonBytes() ([]byte, error) {
	return bindataRead(
		_sourceStruct_with_anonymous_fieldJson,
		"source/struct_with_anonymous_field.json",
	)
}

func sourceStruct_with_anonymous_fieldJson() (*asset, error) {
	bytes, err := sourceStruct_with_anonymous_fieldJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "source/struct_with_anonymous_field.json", size: 1508, mode: os.FileMode(420), modTime: time.Unix(1504168587, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sourceStruct_with_circular_referenceJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xb1\x0a\x02\x31\x0c\x86\xe7\xe4\x29\x24\x3a\x0a\xb7\xbb\xfa\x06\x3a\x38\x8a\xc6\x14\x23\x47\x5b\x72\xb9\x41\x8e\xbe\xbb\xb4\xdc\x20\x22\x1d\x02\xe1\xe7\xfb\xbf\x64\x41\xa0\xb3\xdb\xcc\x7e\x51\x7f\x1e\xd5\x78\x1e\x6f\x76\x92\x20\x26\x91\x85\x0e\x9b\x05\x01\xc8\xdf\xb9\xee\x94\xee\x2f\x61\xa7\x7d\xcb\xd4\xc7\x16\xf6\xfa\x8d\xcc\x96\xb2\x98\xab\x4c\xab\x0f\x88\x57\xf2\x6a\x3f\xa7\xbe\xc5\xff\x75\x00\xb4\x33\x09\x15\xd8\x0e\x0f\x09\x1a\xd5\x35\xc5\x69\xe8\xfd\x51\x7b\x05\xdb\x14\x2c\xf8\x09\x00\x00\xff\xff\xb2\xdf\x8b\x0c\xf7\x00\x00\x00")

func sourceStruct_with_circular_referenceJsonBytes() ([]byte, error) {
	return bindataRead(
		_sourceStruct_with_circular_referenceJson,
		"source/struct_with_circular_reference.json",
	)
}

func sourceStruct_with_circular_referenceJson() (*asset, error) {
	bytes, err := sourceStruct_with_circular_referenceJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "source/struct_with_circular_reference.json", size: 247, mode: os.FileMode(420), modTime: time.Unix(1504168568, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sourceStruct_with_inheritanceJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x96\xc1\x6b\xdb\x30\x14\xc6\xcf\xf6\x5f\x51\xbc\x1d\x3b\x4a\xd3\x10\x4a\x8f\x0b\x14\x72\xe9\xa0\x6d\xd8\x61\x8c\xa0\x38\xcf\xed\x1b\xb2\x64\xe4\x27\x98\x09\xfd\xdf\x87\xe4\x36\x53\xaa\x27\x93\x38\x39\xe4\xf2\xa4\xf7\xfd\xa4\x2f\x9f\x2c\x6d\xf3\xac\xf8\x2e\x5a\x2c\x9f\xbb\x06\xda\xe2\xee\x62\x9b\x67\x59\x41\x5d\x03\xc5\xdd\x45\xa1\xd7\x7f\xa0\xa4\xe2\xd2\xd7\x90\xa4\x2f\x06\xd3\xfd\x40\x63\x74\x03\x86\x70\xd7\x9e\x15\x6b\xad\xe5\xaa\x42\x90\x9b\x8f\xd2\x7f\x4d\x37\x06\x42\xf9\xde\x3d\x59\xad\xe5\xbd\x6f\x71\x03\x6f\x97\xbd\x50\x47\x90\x12\x6a\xc9\xa0\x7a\xf9\xd0\xa9\xb4\xa9\x05\x79\x40\x47\x10\xab\x77\x04\x91\x7a\xa9\xeb\x46\xc2\xdf\xeb\xc9\x6d\x8a\xa1\x6c\xbd\x06\x13\x33\x36\xda\xae\x65\x4c\x99\xef\x04\x53\xac\xd9\xf4\x58\x54\x25\xb5\xa0\x14\x69\x36\x8d\x40\x7e\xfe\xcd\x64\x14\xe6\x66\x12\x81\xee\xfb\x3a\x8f\x19\xb9\x9b\xd9\x94\xc7\x30\xbb\x41\x45\xd7\xb3\x14\x04\x15\xc1\x0b\x47\xf1\x5d\x11\x63\xe1\xaa\x1c\x21\xed\xd6\x10\x81\x31\x6b\xa1\x38\xab\x50\x0d\x18\x35\x44\x60\x7c\x5a\xa8\x84\x4b\xc9\x08\x0f\x01\x6e\x39\xfd\x38\xbb\xa8\x68\x8c\x3a\x27\x1e\x69\x1b\xab\xce\x73\xc2\x1f\xad\x8a\x4f\x78\xaf\x70\x98\xfe\x4e\xe9\xc9\x97\x23\x2d\xc2\xfa\xe8\x95\x6e\x04\xc1\x37\xd7\x18\x41\x9e\xb1\x8e\x97\x6b\xc7\xc5\xdd\xf2\x79\x5f\x22\x1b\x78\x3b\x2e\xf1\x96\x8f\xfc\x12\xd9\xcc\xdb\x71\xa1\xb7\x7c\xea\x97\xc8\xc6\xde\x8e\xca\xbd\x65\x83\xef\x10\x71\xf2\xed\x98\xe8\x5b\x2e\xfb\x4e\x9f\x95\x6f\xc8\x9c\xe9\xdb\xb0\xec\xd5\x02\x4a\xee\x7f\x0e\x56\x3c\xb8\x66\xf9\x44\xc6\x96\x74\xd0\x35\xbf\xd7\x90\xbc\xe8\xdd\x5b\x60\x45\xc1\xdb\x61\xe0\xa5\x90\x65\xc5\x57\x03\x95\x1b\xf9\x72\xb5\x81\x0a\x15\x12\x6a\xd5\x5e\x05\xf3\xc2\x1b\xd3\x80\x20\x38\xc7\x59\x9b\xf7\x4a\x81\xf6\xfb\xf5\x74\xda\xe7\xec\xa1\x17\xf9\xec\x74\x6f\xd9\x4f\xa4\xd7\x39\x9a\xd2\x4a\x61\x1e\xa1\x02\x03\xaa\x84\x83\x8c\x1f\xea\x4f\xfd\x0f\xe5\xfb\xcc\x95\xf9\x84\xda\xb3\x81\x95\x4b\xfe\x2b\x43\xeb\x48\xef\x79\xa1\x5e\xc1\x20\x89\xe3\x77\x1b\x76\xfa\x39\x42\xca\x1f\x6e\x5d\xbf\x1c\x6d\x7b\xea\x52\x7f\xa7\xbc\x53\x3e\xe8\xab\x36\x3c\x1a\x49\xd2\xde\xa9\x08\x5d\xc8\xdf\xf2\x7f\x01\x00\x00\xff\xff\x1c\x2d\xb4\xd4\x4c\x0b\x00\x00")

func sourceStruct_with_inheritanceJsonBytes() ([]byte, error) {
	return bindataRead(
		_sourceStruct_with_inheritanceJson,
		"source/struct_with_inheritance.json",
	)
}

func sourceStruct_with_inheritanceJson() (*asset, error) {
	bytes, err := sourceStruct_with_inheritanceJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "source/struct_with_inheritance.json", size: 2892, mode: os.FileMode(420), modTime: time.Unix(1504168549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sourceStruct_with_no_export_fieldJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\xe2\x54\x0a\x2e\x29\x2a\x4d\x2e\x09\xcf\x2c\xc9\xf0\xcb\x77\xad\x28\xc8\x2f\x2a\x71\xcb\x4c\xcd\x49\x51\xb2\x52\xa8\xe6\xe2\xe4\x54\x2a\xa9\x2c\x48\x55\xb2\x52\x50\xca\x4f\xca\x4a\x4d\x2e\x51\xd2\x01\x8b\x65\x96\xe4\x80\x05\x71\xe9\xe5\xe2\xac\xe5\xaa\xe5\x02\x04\x00\x00\xff\xff\xab\x87\xe6\x3d\x5e\x00\x00\x00")

func sourceStruct_with_no_export_fieldJsonBytes() ([]byte, error) {
	return bindataRead(
		_sourceStruct_with_no_export_fieldJson,
		"source/struct_with_no_export_field.json",
	)
}

func sourceStruct_with_no_export_fieldJson() (*asset, error) {
	bytes, err := sourceStruct_with_no_export_fieldJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "source/struct_with_no_export_field.json", size: 94, mode: os.FileMode(420), modTime: time.Unix(1504168606, 0)}
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
	"source/array_type.json": sourceArray_typeJson,
	"source/basic_types.json": sourceBasic_typesJson,
	"source/map_type.json": sourceMap_typeJson,
	"source/normal_struct.json": sourceNormal_structJson,
	"source/struct_with_anonymous_field.json": sourceStruct_with_anonymous_fieldJson,
	"source/struct_with_circular_reference.json": sourceStruct_with_circular_referenceJson,
	"source/struct_with_inheritance.json": sourceStruct_with_inheritanceJson,
	"source/struct_with_no_export_field.json": sourceStruct_with_no_export_fieldJson,
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
	"source": &bintree{nil, map[string]*bintree{
		"array_type.json": &bintree{sourceArray_typeJson, map[string]*bintree{}},
		"basic_types.json": &bintree{sourceBasic_typesJson, map[string]*bintree{}},
		"map_type.json": &bintree{sourceMap_typeJson, map[string]*bintree{}},
		"normal_struct.json": &bintree{sourceNormal_structJson, map[string]*bintree{}},
		"struct_with_anonymous_field.json": &bintree{sourceStruct_with_anonymous_fieldJson, map[string]*bintree{}},
		"struct_with_circular_reference.json": &bintree{sourceStruct_with_circular_referenceJson, map[string]*bintree{}},
		"struct_with_inheritance.json": &bintree{sourceStruct_with_inheritanceJson, map[string]*bintree{}},
		"struct_with_no_export_field.json": &bintree{sourceStruct_with_no_export_fieldJson, map[string]*bintree{}},
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
