// Package embeddedrules Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// embeddedrules/rules.php
package embeddedrules

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _embeddedrulesRulesPhp = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\xcd\xb1\x6e\x03\x21\x10\x04\xd0\x9e\xaf\x98\x48\x14\xf6\x35\xd7\x3b\x91\x48\x7a\x77\xf9\x01\x73\x77\x6b\xdd\x46\x1c\x20\x58\x64\xf8\xfb\xc8\x38\x72\xdc\x46\x91\x3b\xb4\xc3\xcc\x7b\x33\x71\x8d\x4a\x8d\xc3\xa0\x30\xe0\xdd\x07\xf6\x39\xd2\x2c\x1c\x3c\x3e\x8e\xc7\x7e\x74\xec\x85\x12\x16\xce\x76\x72\xa4\x30\x8c\x0f\x05\xbb\x11\x84\x92\xb7\xa9\x7d\xf2\x16\x1d\x9f\x5b\x0f\x36\xdb\x26\xc2\x1c\x8a\x5b\x90\x28\x3a\x3b\x13\x64\xbd\xff\xc5\x85\x65\xc5\x57\xc9\x02\x3d\x07\xbf\xf4\x8e\xb4\x48\x98\x42\x70\xf7\xdb\xa8\xfa\x0b\x06\x92\x0a\xe1\x80\xb3\x75\x99\x5e\xff\xec\x5f\x12\x0b\xc1\x66\x9c\x76\xd7\xfd\x7d\x5f\x3d\xfd\xa2\x2f\x4f\x56\x75\x85\x39\x40\xb7\x1b\x19\x4b\x22\xe8\x7a\x93\x2a\x0c\x74\xc5\x35\xfd\xaf\x60\x7e\x84\x51\x71\xce\x24\x3b\x5d\xf7\x8f\xe3\xdf\x01\x00\x00\xff\xff\xff\xd4\x33\x95\xec\x01\x00\x00")

func embeddedrulesRulesPhpBytes() ([]byte, error) {
	return bindataRead(
		_embeddedrulesRulesPhp,
		"embeddedrules/rules.php",
	)
}

func embeddedrulesRulesPhp() (*asset, error) {
	bytes, err := embeddedrulesRulesPhpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "embeddedrules/rules.php", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"embeddedrules/rules.php": embeddedrulesRulesPhp,
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
	"embeddedrules": &bintree{nil, map[string]*bintree{
		"rules.php": &bintree{embeddedrulesRulesPhp, map[string]*bintree{}},
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
