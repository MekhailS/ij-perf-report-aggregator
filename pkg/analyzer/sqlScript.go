// Package analyzer Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// pkg/analyzer/sql/create-db.sql
// pkg/analyzer/sql/insert-report.sql
package analyzer

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

// Mode return file modify time
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

var _createDbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xb1\x6e\xe3\x30\x0c\x86\x77\x3d\x05\x47\x1b\xc8\x72\x09\x6e\x0a\x6e\xbb\xe5\xa6\x03\xba\x74\x14\x14\x89\x71\xd8\x5a\x94\x41\xd3\x6d\xf3\xf6\x85\x6a\x39\x4d\x1c\x3b\xa8\x17\x43\xd0\xf7\x51\x94\x7e\x1e\xb0\x21\xde\x1b\xe3\x05\x9d\x22\xa8\x3b\xb4\x08\xd1\xf9\x13\x31\x9a\xca\x00\xb0\x8b\x08\x14\xa0\x57\x21\x6e\x80\x93\x02\x0f\x6d\x0b\x9d\x50\x74\x72\x86\x57\x3c\x9b\x7a\x5e\x40\xb0\x4b\xa2\x5f\x3e\x05\xb8\xfd\x1e\x14\xda\x18\x98\xce\xfe\xe6\x89\x35\xff\x26\x3e\x33\x0d\x32\x8a\x53\x0c\x56\x29\xe2\x32\xa3\xde\x1e\x06\x6a\x83\x9d\x1a\x20\xd6\x8d\x31\x00\x9d\xa4\x30\x78\x5d\x6d\x28\xcb\xa3\xe9\x7f\x3d\x6a\xa2\x30\xdb\x1f\x30\xbb\x75\x26\xdf\x18\x55\xc8\xf7\xf6\x0d\xa5\xa7\xc4\x8b\x85\xc2\x20\x4e\x29\xb1\x2d\xf0\x52\xd3\xc4\xbd\x3a\xd6\x0b\x72\xcf\x18\x00\x71\xef\x76\x4c\x67\xed\xf6\x06\xe0\x98\x04\xa9\xe1\x1c\x09\x54\x25\x90\x1a\x04\x8f\x28\xc8\x1e\xfb\x4b\x48\xd5\xd3\xff\xe7\x7f\x7f\x6b\x48\x0c\x01\x5b\xd4\x9c\x7c\x2e\xe8\xf5\x7a\x26\x88\x03\x7e\xcc\x43\x4b\x5c\xa6\x04\xaa\xdb\x9d\x7a\x7f\xeb\x95\xb3\xec\xb8\xba\xd2\xa6\xc6\x66\x7c\xc9\xf7\x9e\x2f\x1b\x73\x7e\xcc\x28\xba\x97\x24\xf7\xce\x34\x08\x2b\x12\xf1\x03\x69\xbb\x2c\x75\x4e\xfd\x69\x55\xda\xe5\x77\xeb\xc4\x35\xd1\xc1\xd0\xa3\x4c\x53\xf1\xe7\x77\x7e\xcf\x14\x23\xe9\xfe\x33\x00\x00\xff\xff\xdf\x2d\xbc\x8e\xb4\x03\x00\x00")

func createDbSqlBytes() ([]byte, error) {
	return bindataRead(
		_createDbSql,
		"create-db.sql",
	)
}

func createDbSql() (*asset, error) {
	bytes, err := createDbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "create-db.sql", size: 948, mode: os.FileMode(420), modTime: time.Unix(1570797917, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _insertReportSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xcf\xca\xc2\x40\x0c\xc4\xef\x7d\x8a\x1c\x5b\xc8\xe5\xfb\x7c\x00\x29\xd2\x83\x50\x54\xfc\x77\x5d\xd6\xdd\xa0\x01\x9b\x2d\x69\xaa\xaf\x2f\x58\xed\xa9\x35\xe4\x30\x99\xe1\x17\x98\x7d\xb5\xab\xcb\x55\x05\xeb\xcd\x71\x0b\x4a\x6d\x52\x83\x9c\x23\x42\xe3\xc3\x8d\x85\x10\x5a\x4d\xb1\x0f\x86\x19\x4c\xcd\x95\x84\xd4\x1b\x45\x67\xdc\x10\x82\x05\x77\xe9\xf9\x1e\x1d\xc7\x19\x62\x88\xc3\x1f\x7e\xd5\xff\xa8\x16\x33\x48\x43\xa6\x1c\x3a\xf7\x20\xed\x38\x09\x42\xec\xd5\x1b\x27\x71\x9f\x04\x81\xa5\x33\x2f\x36\x1a\xd3\x8f\xd4\x3f\xdd\x50\xb2\xc8\xce\x65\x7d\xaa\x0e\x90\x2f\x11\xde\x3b\x12\x13\xd7\x6f\xa3\x78\x05\x00\x00\xff\xff\x27\x14\xd6\xf3\x45\x01\x00\x00")

func insertReportSqlBytes() ([]byte, error) {
	return bindataRead(
		_insertReportSql,
		"insert-report.sql",
	)
}

func insertReportSql() (*asset, error) {
	bytes, err := insertReportSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "insert-report.sql", size: 325, mode: os.FileMode(420), modTime: time.Unix(1570092302, 0)}
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
	"create-db.sql":     createDbSql,
	"insert-report.sql": insertReportSql,
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
	"create-db.sql":     &bintree{createDbSql, map[string]*bintree{}},
	"insert-report.sql": &bintree{insertReportSql, map[string]*bintree{}},
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
