// Code generated by go-bindata.
// sources:
// scripts/alias.cmd
// scripts/alias.sh
// DO NOT EDIT!

package aliases

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

var _scriptsAliasCmd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8c\x31\xae\x02\x31\x0c\x44\xfb\x9c\x62\x9a\x34\xbf\xfa\x17\x40\xa2\x59\x3a\x1a\xe0\x02\x56\xe2\x10\x4b\xc1\x46\x71\x02\x1d\x67\x47\xab\x5d\xda\x99\xf7\xde\x91\x53\x35\x58\x29\x21\x5c\x96\x33\x8a\x68\x86\xcd\x81\x77\x95\x54\x41\x4d\xc8\x21\x8e\x44\xad\x71\x0e\xd7\xe5\xb6\x6d\x27\xeb\x87\xf8\xd1\xff\xcd\x5a\x5f\xb0\xbe\x52\x13\x14\xeb\x18\x95\x77\x95\x34\xe3\x49\xee\x58\x09\xea\xf7\xf9\x60\x1d\x1e\x76\xb6\x4f\x45\xfc\xf5\x22\xe2\x5f\xf8\x06\x00\x00\xff\xff\xec\x3e\xfb\xca\x8f\x00\x00\x00")

func scriptsAliasCmdBytes() ([]byte, error) {
	return bindataRead(
		_scriptsAliasCmd,
		"scripts/alias.cmd",
	)
}

func scriptsAliasCmd() (*asset, error) {
	bytes, err := scriptsAliasCmdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/alias.cmd", size: 143, mode: os.FileMode(511), modTime: time.Unix(1567334214, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scriptsAliasSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x51\x4d\xcb\x13\x31\x10\xbe\xe7\x57\x3c\xe6\x0d\xbc\xed\x61\x5d\xbd\x5a\x16\x44\xad\x22\x88\x37\xbd\x18\x91\xec\x66\xb2\x1b\x4c\x13\xcd\xc7\x2a\xd4\xfe\x77\xc9\xb6\x5b\x8a\xd2\x53\x66\x32\xcf\xe4\xf9\xc8\xc3\x93\xb6\xa4\xd8\xf6\xd6\xb7\xe4\x67\xf4\x2a\x4d\x8c\x3d\x40\x53\x5f\x46\x1c\x82\x26\xf6\x66\xff\xea\xd3\xbb\x4e\x1c\x97\xf3\x45\x63\x94\x4b\x74\x62\xd6\xe0\x0b\xb8\x58\x2e\x39\xba\x0e\x3c\xc7\x42\x1c\x5f\x77\xc8\x13\x79\x06\x00\x89\x32\x9a\xdf\xcc\xd8\xfa\xa2\xb1\x5e\x23\x94\x8c\x5f\x93\x1d\x26\x28\x67\x55\x82\x4d\x18\x94\x73\xa4\xd9\xd2\xbf\x0d\xb1\x13\x1b\x1a\xa6\x00\xb1\xe9\x55\x22\xaf\x0e\x04\xb1\x89\xa4\xb4\xb3\xfe\x3b\x1a\x6f\x20\x9e\x6d\xb7\xf8\x83\xa1\x64\x34\x06\xcf\xd1\x68\x3c\x3e\x7d\xdc\x56\x8e\x9f\x25\x64\x82\x72\x0e\x2a\x8e\x89\x99\x10\x6b\x01\xeb\xc1\xc5\x4b\xbe\x83\x0e\x58\x84\xa9\x38\x76\x5c\x1c\x55\x1c\xdb\x56\xca\x56\x4a\x29\x4f\xfc\x3c\x71\xae\xae\x76\x5c\x5c\x2a\xc8\x2b\x90\x57\x20\x3f\x49\xce\x99\x0e\x9e\x2a\x61\x15\x0f\xf2\xf3\xe0\x2c\x2a\x5b\x9e\xe8\xe2\x4c\x79\x8d\x1f\x2a\xa5\x55\x4d\x39\x90\xcf\x89\xed\x3f\x7e\x7e\xfd\xe1\xfd\xb7\x35\xd3\xdb\xf6\xdf\x68\x6f\x67\xf7\x12\xa6\x59\x5d\xf9\x9b\xc6\x85\xd1\xd1\x4c\xae\x3b\xff\x5e\x2c\x1e\x62\xcd\x15\xab\x21\x46\x2e\xd1\x7f\xcb\x77\xb0\xc6\xb2\xbf\x01\x00\x00\xff\xff\xdf\x13\x3f\xbc\x21\x02\x00\x00")

func scriptsAliasShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsAliasSh,
		"scripts/alias.sh",
	)
}

func scriptsAliasSh() (*asset, error) {
	bytes, err := scriptsAliasShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/alias.sh", size: 545, mode: os.FileMode(511), modTime: time.Unix(1567331319, 0)}
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
	"scripts/alias.cmd": scriptsAliasCmd,
	"scripts/alias.sh": scriptsAliasSh,
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
	"scripts": &bintree{nil, map[string]*bintree{
		"alias.cmd": &bintree{scriptsAliasCmd, map[string]*bintree{}},
		"alias.sh": &bintree{scriptsAliasSh, map[string]*bintree{}},
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

