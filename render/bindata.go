// Code generated by go-bindata.
// sources:
// render/html/block.tmpl
// render/html/bold.tmpl
// render/html/code-block.tmpl
// render/html/code-inline.tmpl
// render/html/italic.tmpl
// render/html/list.tmpl
// render/html/page.tmpl
// render/html/paragraph.tmpl
// render/html/preformatted.tmpl
// render/html/reference.tmpl
// render/html/section.tmpl
// render/html/sequence.tmpl
// render/html/string.tmpl
// render/html/target.tmpl
// render/html/toc.tmpl
// DO NOT EDIT!

package render

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

var _renderHtmlBlockTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\xc9\x2c\x53\x48\xce\x49\x2c\x2e\xb6\x55\xaa\xae\xd6\x73\x06\xb1\x6a\x6b\x95\xec\x40\xec\xfc\xbc\x92\xd4\xbc\x12\x85\x1a\x85\xa2\xd4\xbc\x94\xd4\xa2\xda\x5a\x1b\xfd\x94\xcc\x32\x3b\x2e\x40\x00\x00\x00\xff\xff\xa6\x01\x03\x65\x34\x00\x00\x00")

func renderHtmlBlockTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlBlockTmpl,
		"render/html/block.tmpl",
	)
}

func renderHtmlBlockTmpl() (*asset, error) {
	bytes, err := renderHtmlBlockTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/block.tmpl", size: 52, mode: os.FileMode(420), modTime: time.Unix(1499727621, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlBoldTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\x2e\x29\xca\xcf\x4b\xb7\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x87\x0a\x73\x01\x02\x00\x00\xff\xff\x40\x12\x1c\xc4\x20\x00\x00\x00")

func renderHtmlBoldTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlBoldTmpl,
		"render/html/bold.tmpl",
	)
}

func renderHtmlBoldTmpl() (*asset, error) {
	bytes, err := renderHtmlBoldTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/bold.tmpl", size: 32, mode: os.FileMode(420), modTime: time.Unix(1499437131, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlCodeBlockTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\x28\x4a\xb5\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x07\x89\x71\x01\x02\x00\x00\xff\xff\x76\xaa\x4d\xfe\x1a\x00\x00\x00")

func renderHtmlCodeBlockTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlCodeBlockTmpl,
		"render/html/code-block.tmpl",
	)
}

func renderHtmlCodeBlockTmpl() (*asset, error) {
	bytes, err := renderHtmlCodeBlockTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/code-block.tmpl", size: 26, mode: os.FileMode(420), modTime: time.Unix(1499107718, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlCodeInlineTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\xce\x4f\x49\xb5\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x07\x0b\x72\x01\x02\x00\x00\xff\xff\x25\x74\x64\xb1\x1c\x00\x00\x00")

func renderHtmlCodeInlineTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlCodeInlineTmpl,
		"render/html/code-inline.tmpl",
	)
}

func renderHtmlCodeInlineTmpl() (*asset, error) {
	bytes, err := renderHtmlCodeInlineTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/code-inline.tmpl", size: 28, mode: os.FileMode(420), modTime: time.Unix(1499107718, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlItalicTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\xcd\xb5\xab\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x4f\xcd\xb5\xe3\x02\x04\x00\x00\xff\xff\x99\x57\x43\xb8\x18\x00\x00\x00")

func renderHtmlItalicTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlItalicTmpl,
		"render/html/italic.tmpl",
	)
}

func renderHtmlItalicTmpl() (*asset, error) {
	bytes, err := renderHtmlItalicTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/italic.tmpl", size: 24, mode: os.FileMode(420), modTime: time.Unix(1499171452, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlListTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcc\x31\x0a\x82\x31\x10\x44\xe1\x3e\xa7\xd8\x13\x24\x17\x08\xe9\xad\x3c\x83\xb0\xa3\x04\xd6\x08\xfb\x6b\x35\xce\xdd\x05\x1b\x63\xf7\x9a\xf7\x91\xf3\x6a\xf5\x9c\x8e\x84\x4b\xfd\x11\x83\x44\x1c\x90\xfa\xeb\xdb\xcb\xa5\x42\xe6\x65\xdd\x60\xf5\xf4\xc4\xfd\x90\x8a\x59\x8f\x39\xc8\x6a\x6f\x4b\x2c\x47\x4a\xbd\xc5\x1c\xe5\xb7\xfc\xcb\x6d\xa7\xdb\x66\x7f\x02\x00\x00\xff\xff\x26\x88\xac\xef\x83\x00\x00\x00")

func renderHtmlListTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlListTmpl,
		"render/html/list.tmpl",
	)
}

func renderHtmlListTmpl() (*asset, error) {
	bytes, err := renderHtmlListTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/list.tmpl", size: 131, mode: os.FileMode(420), modTime: time.Unix(1499801773, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlPageTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8e\x41\x4f\xc5\x20\x10\x84\xef\xfe\x8a\x95\x7b\xe1\x6a\x22\xed\x45\x3d\x6b\x62\x2f\x1e\xb1\x5d\x85\x04\x28\xb6\x53\xf3\x1a\xc2\x7f\x7f\xe1\xd1\x77\xda\xd9\xec\xce\x7c\xa3\x1f\x5f\xdf\x5f\xc6\xaf\x8f\x37\xb2\x08\x7e\x78\xd0\x6d\x10\x69\xcb\x66\xae\x82\x48\x07\x86\x21\x0b\xa4\x8e\xff\x76\xf7\xdf\x8b\x69\x89\xe0\x88\x0e\x47\x62\x41\xe7\xd6\x0b\xf0\x05\xaa\x06\x3c\xd3\x64\xcd\xba\x31\xfa\x1d\x3f\xdd\x93\x20\x75\x26\xc1\xc1\xf3\x90\xb3\x1c\xab\x90\x9f\x58\x5d\xfc\x2d\x45\xab\x76\xa8\x5c\x75\x07\xeb\xef\x65\x3e\x9a\x2f\x67\x70\x48\xde\x80\x49\x6c\x3c\xc1\x2d\x51\x22\x24\x2f\x48\x96\x72\x33\xb5\x5f\xad\x5a\xfd\x6b\x00\x00\x00\xff\xff\x50\x1e\xa4\x17\xd6\x00\x00\x00")

func renderHtmlPageTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlPageTmpl,
		"render/html/page.tmpl",
	)
}

func renderHtmlPageTmpl() (*asset, error) {
	bytes, err := renderHtmlPageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/page.tmpl", size: 214, mode: os.FileMode(420), modTime: time.Unix(1499107718, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlParagraphTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x29\xb0\xab\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xcc\x4b\x49\xad\xd0\x51\x50\x29\x4e\xcd\x2b\x49\xcd\x4b\x4e\x55\xb0\xb2\x55\xd0\xab\xad\xad\xae\xce\x4c\x83\x4a\xd6\xd6\x2a\x54\x57\xa7\xe6\xa5\x80\x04\x11\xca\x6a\x14\x8a\x52\xf3\x52\x52\x8b\x40\xa2\x60\x49\x1b\xfd\x02\x3b\x2e\x40\x00\x00\x00\xff\xff\xd3\x3e\x01\x8c\x5a\x00\x00\x00")

func renderHtmlParagraphTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlParagraphTmpl,
		"render/html/paragraph.tmpl",
	)
}

func renderHtmlParagraphTmpl() (*asset, error) {
	bytes, err := renderHtmlParagraphTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/paragraph.tmpl", size: 90, mode: os.FileMode(420), modTime: time.Unix(1498953201, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlPreformattedTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xcc\x4b\x49\xad\xd0\x51\x50\x29\x4e\xcd\x2b\x49\xcd\x4b\x4e\x55\xb0\xb2\x55\xd0\xab\xad\xad\xae\xce\x4c\x83\x4a\xd6\xd6\x72\x55\x57\xa7\xe6\xa5\x80\x04\x11\xca\x6a\x14\x8a\x52\xf3\x52\x52\x8b\x40\xa2\x60\x49\x2e\x40\x00\x00\x00\xff\xff\x53\xe4\x01\x6f\x53\x00\x00\x00")

func renderHtmlPreformattedTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlPreformattedTmpl,
		"render/html/preformatted.tmpl",
	)
}

func renderHtmlPreformattedTmpl() (*asset, error) {
	bytes, err := renderHtmlPreformattedTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/preformatted.tmpl", size: 83, mode: os.FileMode(420), modTime: time.Unix(1499115653, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlReferenceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\x54\xc8\x28\x4a\x4d\xb3\x55\xaa\xae\xd6\x0b\x49\x4c\x57\xa8\x51\x28\x2d\xca\xa9\xad\x55\xb2\xab\xae\xd6\x73\xc9\x2c\x2e\xc8\x49\xac\x54\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\xaa\xad\xb5\xd1\x4f\xb4\xe3\x02\x04\x00\x00\xff\xff\x85\xfd\x53\x5c\x33\x00\x00\x00")

func renderHtmlReferenceTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlReferenceTmpl,
		"render/html/reference.tmpl",
	)
}

func renderHtmlReferenceTmpl() (*asset, error) {
	bytes, err := renderHtmlReferenceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/reference.tmpl", size: 51, mode: os.FileMode(420), modTime: time.Unix(1498873062, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlSectionTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\xc1\x4e\xc4\x30\x0c\x44\xef\xf9\x8a\x51\xef\xcd\xfe\x40\xd9\x03\x70\x46\x48\xcb\x0f\x84\x8d\x97\x46\x4a\xdd\x28\x31\x07\x64\xfc\xef\x28\xdd\x96\x03\x9c\x6c\x8d\x67\xc6\x6f\x9a\x55\x67\x0a\x91\xea\x33\x15\x99\xe1\xcd\x70\xcd\xa1\xb5\x87\xa1\xd1\x55\xd2\xca\xe3\xfd\x3c\x9c\x1d\xa0\x3a\x22\xdd\xe0\x5f\x43\x25\x16\x8c\x66\x0e\x00\xa6\x56\x02\xff\x8d\xf1\xe7\xf2\xde\x63\xaa\xfe\x65\x5b\xcd\x30\x9d\xba\xf3\x68\x22\x8e\x7b\x85\xaa\x7f\x4b\x92\x09\xdf\xa8\xc4\x91\xea\xa6\x4f\xa7\xff\x70\x67\xe7\x54\xfd\xe3\x1a\xbf\x7e\xbd\x66\x5d\x4b\x37\xf0\x2a\xf0\x97\x92\x93\x5c\xee\x0c\x6d\x2f\xaf\x81\x3f\x08\xfe\x69\x4e\x39\x56\xe2\x9d\x5a\x55\x68\x29\x39\x08\xe1\x80\xf6\xb2\x94\x3c\xf4\x3f\x5b\x8e\x38\x9a\xb9\x63\xfe\x04\x00\x00\xff\xff\x42\xb1\xd3\x52\x2b\x01\x00\x00")

func renderHtmlSectionTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlSectionTmpl,
		"render/html/section.tmpl",
	)
}

func renderHtmlSectionTmpl() (*asset, error) {
	bytes, err := renderHtmlSectionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/section.tmpl", size: 299, mode: os.FileMode(420), modTime: time.Unix(1499119690, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlSequenceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\xd0\x73\xce\xcf\x2b\x49\xcd\x2b\x29\xae\xad\xad\xae\xd6\x53\xa8\x51\x28\x4a\xcd\x4b\x49\x2d\x02\xf1\x52\xf3\x52\x6a\x6b\xb9\x00\x01\x00\x00\xff\xff\xfc\xbe\x50\x06\x29\x00\x00\x00")

func renderHtmlSequenceTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlSequenceTmpl,
		"render/html/sequence.tmpl",
	)
}

func renderHtmlSequenceTmpl() (*asset, error) {
	bytes, err := renderHtmlSequenceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/sequence.tmpl", size: 41, mode: os.FileMode(420), modTime: time.Unix(1498830402, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlStringTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\xd0\x0b\x2e\x29\xca\xcc\x4b\x57\xa8\xad\xe5\x02\x04\x00\x00\xff\xff\xc7\x02\x8e\x75\x0e\x00\x00\x00")

func renderHtmlStringTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlStringTmpl,
		"render/html/string.tmpl",
	)
}

func renderHtmlStringTmpl() (*asset, error) {
	bytes, err := renderHtmlStringTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/string.tmpl", size: 14, mode: os.FileMode(420), modTime: time.Unix(1498870948, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlTargetTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\x54\xc8\x4b\xcc\x4d\xb5\x55\xaa\xae\xd6\x0b\x49\x4c\xf7\x4b\xcc\x4d\xad\xad\x55\xb2\xb3\xd1\x4f\xb4\xe3\x02\x04\x00\x00\xff\xff\x6d\x16\xe9\x86\x1c\x00\x00\x00")

func renderHtmlTargetTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlTargetTmpl,
		"render/html/target.tmpl",
	)
}

func renderHtmlTargetTmpl() (*asset, error) {
	bytes, err := renderHtmlTargetTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/target.tmpl", size: 28, mode: os.FileMode(420), modTime: time.Unix(1499649481, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _renderHtmlTocTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x31\xca\xc3\x30\x0c\x46\xf7\x9c\xe2\xc3\x07\xb0\x2f\xe0\xdf\xcb\xbf\x97\x0e\xb9\x80\xdb\x28\x89\x41\x76\x83\x70\x0a\x45\xf1\xdd\x4b\xd2\x66\xe8\x24\xf1\x10\x7a\x4f\x35\x8d\xb0\xff\x73\xe2\x41\xa8\xb4\xd6\xf9\x12\x9f\xa1\x03\xfc\xca\xfb\x50\x95\x58\x26\xfa\x39\x01\x00\xcf\x29\x1c\x0b\xe0\x23\x66\xa1\xf1\xcf\xa8\xda\xab\xa4\x1c\xe5\xd5\xc7\x09\x1b\x56\xe1\xd6\x4c\x50\xb5\x97\x35\xdf\x48\x5a\x83\xaa\xed\x53\x65\xc2\x06\xa1\x32\xec\xcc\xbb\x18\xba\xef\x2b\xd5\x4a\x79\xe1\x58\x09\xa6\x3e\xee\xb6\xe6\x85\x0d\xec\xe9\x74\x1f\xa9\x2a\x95\xe1\x60\xde\xed\x95\xde\x1d\xcd\x27\x7e\x07\x00\x00\xff\xff\x8e\xbb\xb4\xef\xd3\x00\x00\x00")

func renderHtmlTocTmplBytes() ([]byte, error) {
	return bindataRead(
		_renderHtmlTocTmpl,
		"render/html/toc.tmpl",
	)
}

func renderHtmlTocTmpl() (*asset, error) {
	bytes, err := renderHtmlTocTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "render/html/toc.tmpl", size: 211, mode: os.FileMode(420), modTime: time.Unix(1499010799, 0)}
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
	"render/html/block.tmpl": renderHtmlBlockTmpl,
	"render/html/bold.tmpl": renderHtmlBoldTmpl,
	"render/html/code-block.tmpl": renderHtmlCodeBlockTmpl,
	"render/html/code-inline.tmpl": renderHtmlCodeInlineTmpl,
	"render/html/italic.tmpl": renderHtmlItalicTmpl,
	"render/html/list.tmpl": renderHtmlListTmpl,
	"render/html/page.tmpl": renderHtmlPageTmpl,
	"render/html/paragraph.tmpl": renderHtmlParagraphTmpl,
	"render/html/preformatted.tmpl": renderHtmlPreformattedTmpl,
	"render/html/reference.tmpl": renderHtmlReferenceTmpl,
	"render/html/section.tmpl": renderHtmlSectionTmpl,
	"render/html/sequence.tmpl": renderHtmlSequenceTmpl,
	"render/html/string.tmpl": renderHtmlStringTmpl,
	"render/html/target.tmpl": renderHtmlTargetTmpl,
	"render/html/toc.tmpl": renderHtmlTocTmpl,
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
	"render": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"block.tmpl": &bintree{renderHtmlBlockTmpl, map[string]*bintree{}},
			"bold.tmpl": &bintree{renderHtmlBoldTmpl, map[string]*bintree{}},
			"code-block.tmpl": &bintree{renderHtmlCodeBlockTmpl, map[string]*bintree{}},
			"code-inline.tmpl": &bintree{renderHtmlCodeInlineTmpl, map[string]*bintree{}},
			"italic.tmpl": &bintree{renderHtmlItalicTmpl, map[string]*bintree{}},
			"list.tmpl": &bintree{renderHtmlListTmpl, map[string]*bintree{}},
			"page.tmpl": &bintree{renderHtmlPageTmpl, map[string]*bintree{}},
			"paragraph.tmpl": &bintree{renderHtmlParagraphTmpl, map[string]*bintree{}},
			"preformatted.tmpl": &bintree{renderHtmlPreformattedTmpl, map[string]*bintree{}},
			"reference.tmpl": &bintree{renderHtmlReferenceTmpl, map[string]*bintree{}},
			"section.tmpl": &bintree{renderHtmlSectionTmpl, map[string]*bintree{}},
			"sequence.tmpl": &bintree{renderHtmlSequenceTmpl, map[string]*bintree{}},
			"string.tmpl": &bintree{renderHtmlStringTmpl, map[string]*bintree{}},
			"target.tmpl": &bintree{renderHtmlTargetTmpl, map[string]*bintree{}},
			"toc.tmpl": &bintree{renderHtmlTocTmpl, map[string]*bintree{}},
		}},
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

