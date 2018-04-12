// Code generated by go-bindata.
// sources:
// resources/en.yml
// DO NOT EDIT!

package resources

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

var _resourcesEnYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x94\x5b\x92\xea\x20\x10\x86\xdf\x5d\x45\x6f\xe0\x54\x91\xe3\x3d\xbb\x51\x69\x27\x5d\x45\xc0\x22\xa0\x33\xb3\xfa\x29\xe3\xcc\x43\x5f\x88\x3e\xa6\x3f\x7e\xe0\xa3\x6d\x8c\xfd\x0a\xc0\xe3\x85\xc6\x53\x98\x7a\x58\x01\x00\xfc\x83\x47\xca\xbe\x87\x5b\xa2\x58\xd8\x97\x53\xf4\x2b\x80\x4b\xaa\xb1\x60\x9e\x7a\x56\xfb\xc6\x9c\xe6\x0f\x00\xb1\x8e\x67\xcc\x3d\x38\x06\xa4\x61\xb9\x1c\x51\xd4\x3b\x56\x2f\x0f\x19\xff\x9f\xd7\x87\x8c\x32\x61\xcd\x88\x6b\xaa\x59\x00\x1b\x0e\xd0\x5d\x26\x6c\x19\x30\xd1\xa7\xa8\xef\x78\x1d\xef\x18\x05\xb1\x67\x04\xd2\xc7\x50\x04\x71\x60\x44\x24\x25\xe2\xc8\x2f\xaa\xb6\xe8\xb8\x49\x0c\xc6\x31\x3a\x69\x13\x83\xba\x6c\xa7\x8c\x52\x41\x9d\xa4\xad\x5a\x94\x54\x7b\xb5\x20\xa5\xd7\x82\x0c\xc7\x16\x66\x88\x36\x28\x2d\xdb\xa2\x8e\x52\x56\x2c\x5f\xb2\xfb\x9c\x90\x45\x59\x31\x6b\xa7\x55\x49\x66\xe3\x94\x28\x89\x6c\x9d\xd2\x24\x91\x9d\x33\x24\x49\x68\xef\xb4\x22\xc9\x1c\x9c\x16\x24\x99\xe3\x93\x19\x6b\x28\x74\x0b\xa4\x06\xc1\x50\xa3\xcf\xe8\x55\x8f\x4a\x5f\xa9\x4e\xaf\x79\x22\x30\xce\x8d\x14\x02\x25\xdd\xf1\xf3\x8f\x91\xe7\x25\x52\xc2\x25\xbf\xa1\x5f\x0b\x3c\xdd\xc9\xab\x1b\x16\x8c\x45\x0e\x33\xf1\x1f\x9c\x91\x69\x99\xf9\x15\x65\x44\x35\x38\x9d\xc7\xc1\x0b\x46\x39\x5b\x2c\xe4\x5d\xcc\xdf\xd3\x58\x07\x6b\x91\x46\xa6\xf9\x8e\x66\x66\x93\xb5\x52\x1b\xcf\xde\x0a\x6e\xe2\x8d\xec\x56\xa3\x2c\xe4\xb7\x97\xb4\xf7\x78\xae\xf9\x09\x00\x00\xff\xff\x89\x97\x2d\xc7\x7d\x07\x00\x00")

func resourcesEnYmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesEnYml,
		"resources/en.yml",
	)
}

func resourcesEnYml() (*asset, error) {
	bytes, err := resourcesEnYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/en.yml", size: 1917, mode: os.FileMode(420), modTime: time.Unix(1523533889, 0)}
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
	"resources/en.yml": resourcesEnYml,
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
	"resources": &bintree{nil, map[string]*bintree{
		"en.yml": &bintree{resourcesEnYml, map[string]*bintree{}},
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
