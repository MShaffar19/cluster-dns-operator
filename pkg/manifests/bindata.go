// Code generated by go-bindata.
// sources:
// manifests/00-cluster-role.yaml
// manifests/00-custom-resource-definition.yaml
// manifests/00-dns-namespace.yaml
// manifests/00-operator-namespace.yaml
// manifests/00-role.yaml
// manifests/01-cluster-role-binding.yaml
// manifests/01-role-binding.yaml
// manifests/01-service-account.yaml
// manifests/02-deployment.yaml
// manifests/image-references
// assets/dns/cluster-role-binding.yaml
// assets/dns/cluster-role.yaml
// assets/dns/configmap.yaml
// assets/dns/daemonset.yaml
// assets/dns/service-account.yaml
// assets/dns/service.yaml
// DO NOT EDIT!

package manifests

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

var _manifests00ClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\xbf\x6e\xe3\x30\x0c\xc6\x77\x3d\x05\x91\x9b\xed\x43\xb6\x83\xd6\x1b\xba\x77\xe8\x4e\x4b\x74\x4c\x44\x16\x05\x92\x4e\x81\x3e\x7d\x11\xc7\x09\x8a\xba\x7f\x26\x51\xe4\xc7\xdf\x27\x0a\xfc\x03\xff\xcb\x62\x4e\x0a\x2a\x85\x60\x14\x05\x9f\x08\xa4\x91\xa2\x8b\x02\xbb\x51\x19\xfb\x70\xe6\x9a\xe3\x5d\xfb\x2c\x85\x02\x36\x7e\x21\x35\x96\x1a\x41\x07\x4c\x3d\x2e\x3e\x89\xf2\x1b\x3a\x4b\xed\xcf\xff\xac\x67\xf9\x7b\x39\x0e\xe4\x78\x0c\x33\x39\x66\x74\x8c\x01\xa0\xe2\x4c\x11\xd2\x8d\xd5\xe5\x6a\xdd\xdd\x2e\x4a\xa3\x6a\x13\x8f\xde\x7d\x55\xde\x7a\xad\x61\xa2\x08\xbf\x68\x75\x29\x64\x31\x74\x80\x8d\x9f\x54\x96\x66\x57\xef\x0e\x72\xb5\xfe\xd1\xda\xb3\x04\x00\x25\x93\x45\x13\x6d\x8a\x0d\x97\xab\x91\x05\x80\x0b\xe9\xb0\x55\x0a\x9b\xaf\xc1\x2b\x7a\x9a\xf6\x6c\x6c\xcd\xf6\xbc\x8c\x34\x4b\x35\xf2\x4f\xb4\xa4\x84\x4e\x6b\x78\xa2\x1b\x37\x53\x21\xa7\x3d\xf8\x70\x58\x8f\x6f\xff\x79\x6f\x6a\xa4\x17\x4e\xeb\x00\x8f\x0b\xa6\x24\x4b\xf5\x5b\x2e\x49\x1d\xf9\x34\x63\xb3\x8f\x43\x5f\xb7\x60\x97\x18\xb8\x66\xae\xa7\x9f\x9f\xff\x1e\x00\x00\xff\xff\xad\xff\xc1\x5f\x4a\x02\x00\x00")

func manifests00ClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests00ClusterRoleYaml,
		"manifests/00-cluster-role.yaml",
	)
}

func manifests00ClusterRoleYaml() (*asset, error) {
	bytes, err := manifests00ClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/00-cluster-role.yaml", size: 586, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests00CustomResourceDefinitionYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xb1\x4e\x03\x31\x0c\x86\xf7\x3c\x85\x9f\xe0\xd0\x6d\x28\x2b\xdd\x40\x1d\x40\x62\x37\x89\xdb\x5a\xcd\xd9\x56\xec\x9c\x78\x7c\x74\x57\x90\x80\x8e\xf6\xf7\xeb\xff\x2c\xa3\xf1\x3b\x75\x67\x95\x0c\x68\x4c\x9f\x41\xb2\x4d\x3e\x5d\x1f\x7d\x62\x7d\x58\xe7\x0f\x0a\x9c\xd3\x95\xa5\x66\x78\x1a\x1e\xba\xbc\x92\xeb\xe8\x85\x0e\x74\x62\xe1\x60\x95\xb4\x50\x60\xc5\xc0\x9c\x00\x04\x17\xca\x50\xda\xf0\xa0\x5e\xc5\xc9\xa7\x2a\x3e\xa9\x91\xf8\x85\x4f\x31\xb1\x26\x37\x2a\x5b\xf6\xdc\x75\x58\x86\x3b\x7e\x6b\xf1\x2d\x02\xf0\xed\xbe\x15\x1e\x8e\x6f\xfb\xb2\xb1\xc7\xf3\x3f\xf0\xc2\x1e\x3b\xb4\x36\x3a\xb6\xbf\x47\xec\xc0\x59\xce\xa3\x61\xff\x8d\x12\x80\x17\x35\xca\x70\xdc\x9c\x86\x85\x6a\x02\x58\x7f\xfe\xb2\xce\xd8\xec\x82\x73\xfa\x0a\x00\x00\xff\xff\x3a\xfe\x93\xd9\x2d\x01\x00\x00")

func manifests00CustomResourceDefinitionYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests00CustomResourceDefinitionYaml,
		"manifests/00-custom-resource-definition.yaml",
	)
}

func manifests00CustomResourceDefinitionYaml() (*asset, error) {
	bytes, err := manifests00CustomResourceDefinitionYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/00-custom-resource-definition.yaml", size: 301, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests00DnsNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\xb1\xae\x42\x31\x08\x06\xe0\x9d\xa7\x20\x67\xef\xbd\x71\xe5\x21\x1c\xdd\x49\xfb\x1b\x89\xa7\xd0\x14\xf4\xf9\x8d\x93\xfb\xf7\x34\x1f\xc2\x57\x9d\xc8\xa5\x1d\xa4\xcb\x6e\xd8\x69\xe1\xc2\xef\x0b\x4d\x94\x0e\x2d\x15\x62\x76\x9d\x10\x8e\x05\xcf\x87\xdd\xab\xf5\xf3\x95\x85\xdd\x86\x27\x31\xab\x7b\x94\x96\x85\xe7\x17\xf3\x0f\xfe\x59\xfc\x7b\x0c\xb4\xc4\x89\x5e\xb1\x85\x8f\x83\x3e\x01\x00\x00\xff\xff\xb5\x9f\xce\xf1\x79\x00\x00\x00")

func manifests00DnsNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests00DnsNamespaceYaml,
		"manifests/00-dns-namespace.yaml",
	)
}

func manifests00DnsNamespaceYaml() (*asset, error) {
	bytes, err := manifests00DnsNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/00-dns-namespace.yaml", size: 121, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests00OperatorNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xb1\x11\x82\x31\x08\x05\xe0\x3e\x53\xb0\x40\x0a\xdb\x0c\x61\x69\xff\xee\xcf\xf3\xe4\x34\xc0\x01\x3a\xbf\xdf\x5b\x6d\x2f\xb9\xe3\xb0\x02\x17\x07\x42\x1f\xcc\x52\xb7\x25\xbf\xdb\x38\x6c\x6c\x34\xd6\x10\x11\x31\x1c\x2e\xf1\xa0\xd5\x4b\x9f\x3d\xaf\xcf\xb7\x9a\x39\xb7\xd5\xf4\x60\xa2\x3d\xc7\x3f\x00\x00\xff\xff\x93\x5a\xd8\xf8\x52\x00\x00\x00")

func manifests00OperatorNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests00OperatorNamespaceYaml,
		"manifests/00-operator-namespace.yaml",
	)
}

func manifests00OperatorNamespaceYaml() (*asset, error) {
	bytes, err := manifests00OperatorNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/00-operator-namespace.yaml", size: 82, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests00RoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xb1\x6a\x80\x40\x0c\x86\xf7\x7b\x8a\x60\xb7\xc2\x59\xdc\xca\xbd\x40\xf7\x0e\xdd\xe3\x19\x31\xa8\x97\x23\xc9\x39\xf4\xe9\x8b\xd2\x82\x05\xb7\xf0\xc3\x97\xef\x7b\x81\x4f\xd9\x08\x66\x51\xf0\x85\x40\x2a\x29\xba\x28\x70\x39\xef\x62\x0b\xcf\x1e\xf3\xd6\xcc\x49\xe3\x54\x0c\x0a\xee\x64\x15\x33\x85\x95\xcb\x94\x2e\x3c\x60\xe5\x2f\x52\x63\x29\x09\x74\xc4\xdc\x63\xf3\x45\x94\xbf\xd1\x59\x4a\xbf\xbe\x5b\xcf\xf2\x76\x0c\x23\x39\x0e\x61\x27\xc7\x09\x1d\x53\x80\xeb\x5d\x82\x9b\x20\xfe\x25\xa4\x47\xff\x2f\x72\x15\xa4\xe7\xc4\xa0\x6d\x23\x4b\x21\x02\x56\xfe\x50\x69\xd5\x4e\x53\x84\xee\xb5\x0b\x00\x4a\x26\x4d\x33\xfd\x1b\x0f\xd2\xf1\x36\xfc\x04\x00\x00\xff\xff\xf9\xc2\x56\x68\x16\x01\x00\x00")

func manifests00RoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests00RoleYaml,
		"manifests/00-role.yaml",
	)
}

func manifests00RoleYaml() (*asset, error) {
	bytes, err := manifests00RoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/00-role.yaml", size: 278, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests01ClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\xb1\x4e\xc3\x40\x0c\x86\xf7\x7b\x0a\xab\xcc\x09\xea\x86\x6e\x03\x76\x86\x22\xb1\x3b\x17\x97\x98\x26\x76\x64\x3b\x95\xe0\xe9\xd1\x29\xe9\x44\x45\x85\x58\x6d\xcb\xff\xf7\xfd\x77\xf0\xc4\xd2\x3b\xc4\x40\xa0\x33\x19\x86\x1a\x94\x71\xf1\x20\x03\xd3\x91\x20\x14\x38\x1c\x5e\xc9\xce\x5c\x08\x1e\x4b\xd1\x45\xa2\x4d\x27\x96\x3e\xc3\xf3\x7a\x7a\xd0\x91\xea\x23\x96\xf7\x84\x33\xbf\x91\x39\xab\x64\xb0\x0e\x4b\x8b\x4b\x0c\x6a\xfc\x85\xc1\x2a\xed\xe9\xc1\x5b\xd6\xfb\xf3\xbe\xa3\xc0\x7d\x9a\x28\xb0\xc7\xc0\x9c\x00\x04\x27\xca\x97\xf4\xa6\x17\x6f\x2e\x48\x59\x67\x12\x1f\xf8\x18\xcd\xb5\x75\xf2\xa5\xfb\xa0\x12\x9e\x53\x03\x2b\xd8\xc6\xbb\xe1\xfe\xfa\x7c\x5b\xfa\x8c\x85\x32\xdc\x48\xaa\x9d\x1c\xe8\x58\x71\x7f\x34\xf0\x2f\x85\xbf\x40\x2c\x4e\xf6\x52\xaf\xab\xef\xce\x3f\x3d\x68\xca\xbe\x1a\xe3\x6a\x7c\x23\x2e\x5f\x1b\xee\xd2\x77\x00\x00\x00\xff\xff\x75\x86\xdd\x2d\x0f\x02\x00\x00")

func manifests01ClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests01ClusterRoleBindingYaml,
		"manifests/01-cluster-role-binding.yaml",
	)
}

func manifests01ClusterRoleBindingYaml() (*asset, error) {
	bytes, err := manifests01ClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/01-cluster-role-binding.yaml", size: 527, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests01RoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\xb1\x4e\xc4\x30\x0c\x86\xf7\x3c\x85\x75\xcc\x2d\xba\x0d\x65\x83\x07\x60\x38\x24\x76\x37\xf5\x51\x73\x6d\x5c\xd9\x4e\x25\x78\x7a\x14\xb5\xa8\x0c\x15\x12\xac\xf9\xff\xe4\xff\xbe\xdc\xc1\x13\xe7\xde\xc0\x07\x02\x99\x49\xd1\x45\x41\x65\x24\x70\x01\x76\x83\x17\xd2\x85\x13\xc1\x63\x4a\x52\xb2\xb7\xe1\xc6\xb9\x8f\x70\x91\x91\xea\x4d\xce\x6f\x01\x67\x7e\x25\x35\x96\x1c\x41\x3b\x4c\x2d\x16\x1f\x44\xf9\x13\x9d\x25\xb7\xb7\x07\x6b\x59\xee\x97\x73\x47\x8e\xe7\x30\x91\x63\x8f\x8e\x31\x00\x64\x9c\x28\x42\x1a\x8b\x39\x69\xd3\x67\x6b\xbe\x19\xa2\xcc\x94\x6d\xe0\xab\x37\x3f\xe2\x60\xa5\x7b\xa7\xe4\x16\x43\x03\x2b\xc8\xc6\xb7\xe1\xfd\xfa\xe6\x16\xda\x8c\x89\x22\x1c\x0e\xec\xdd\xfa\x07\x17\xba\x56\xca\xdd\xf8\x3f\xc8\x7f\x19\x2d\x46\xfa\x5c\xdb\xd5\xef\x64\x1f\xe6\x34\x45\x5b\x0d\x71\x35\x3c\x5e\xd9\x21\x8e\x0e\x4f\xe1\x2b\x00\x00\xff\xff\x66\x7f\x2a\xe8\xe7\x01\x00\x00")

func manifests01RoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests01RoleBindingYaml,
		"manifests/01-role-binding.yaml",
	)
}

func manifests01RoleBindingYaml() (*asset, error) {
	bytes, err := manifests01RoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/01-role-binding.yaml", size: 487, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests01ServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xb1\x0a\xc3\x30\x0c\x44\x77\x7f\xc5\x41\xe6\x04\xba\x7a\xeb\xd8\xb9\xd0\xdd\xd8\x0a\x11\x4d\x2c\x57\x92\xf3\xfd\x25\x21\xdd\x3a\xdf\xe3\xdd\x1b\x70\xcf\x59\x7a\x75\xcc\xa2\xf0\x85\x20\x8d\x34\xb9\x28\xd8\x8d\xd6\x79\xc2\xc3\x61\x8b\xf4\xb5\x40\xe9\xd3\x59\x09\x35\x6d\x64\x2d\x65\x82\x65\x69\x54\xc2\x80\x46\xba\xb1\x19\x4b\xb5\x29\xbc\xb9\x96\x88\x27\xe9\xce\x99\x2e\x7f\x48\x8d\x5f\xa4\x07\x11\xb1\xdf\xc2\x46\x9e\x4a\xf2\x14\x03\x4e\x5f\x44\x5e\xbb\x39\xe9\x58\xaa\x8d\xbf\x88\x6b\x3c\xcf\xe2\x91\x56\x6d\xe1\xd9\xc7\xbf\xec\x37\x00\x00\xff\xff\x83\x8f\x49\xa7\xcc\x00\x00\x00")

func manifests01ServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests01ServiceAccountYaml,
		"manifests/01-service-account.yaml",
	)
}

func manifests01ServiceAccountYaml() (*asset, error) {
	bytes, err := manifests01ServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/01-service-account.yaml", size: 204, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests02DeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x4f\x6f\xf2\x30\x0c\xc6\xef\xfd\x14\x16\xf7\xbe\x88\xf7\x98\x5b\xc5\xd8\x1f\x69\x83\x0a\xd0\x76\x9c\xbc\xd4\x85\x68\x49\x1c\x25\x6e\x25\xbe\xfd\xd4\xd1\x41\xd9\x3a\x31\x1f\xf3\x3c\x7e\xfc\xab\x5d\x0c\xe6\x99\x62\x32\xec\x15\x60\x08\x69\xda\xce\xb2\x77\xe3\x2b\x05\x37\x14\x2c\x1f\x1c\x79\xc9\x1c\x09\x56\x28\xa8\x32\x00\x8f\x8e\x14\x68\xdb\x24\xa1\x98\x57\x3e\xe5\x1c\x28\xa2\x70\xec\xc5\x14\x50\x93\x02\x0e\xe4\xd3\xde\xd4\x92\x8f\x7a\x53\x20\xdd\xc5\x45\x0a\xd6\x68\x4c\x0a\x66\x19\x40\x22\x4b\x5a\x38\x76\x0a\x80\x43\xd1\xfb\x47\x7c\x23\x9b\x8e\x0f\x57\xa6\x0b\xb9\x60\x51\xa8\xef\x1e\x40\x77\x65\x2f\x82\xae\x44\x01\x7c\x01\x76\xa5\xd9\x0b\x1a\x4f\x71\xd0\x9e\x5f\x0b\x38\x96\x71\xb8\x1b\x6e\x63\xca\xd1\xec\x8c\x1f\x5d\x8a\xea\xe0\x93\x0c\xba\x35\x3b\x87\xbe\x52\x83\xa7\xfc\x4f\x13\xcb\xc6\xda\x92\xad\xd1\x07\x05\x0f\xf5\x92\xa5\x8c\x94\xba\x53\x9e\x7d\x42\xd1\x19\x8f\x62\xd8\xdf\x45\xd4\x54\x52\x34\x5c\x6d\x48\xb3\xaf\x92\x82\xff\x03\x2b\xf9\x76\x88\x70\xfe\xf8\x97\x62\x3b\xbf\x7f\x5d\x16\x4f\x8b\x4d\x59\xcc\x17\x17\x1e\x80\x16\x6d\x43\xb7\x91\x9d\xfa\x26\x00\xd4\x86\x6c\xb5\xa6\xfa\xa7\xd2\x6b\x25\xca\x5e\x9d\x6e\xf8\xef\xf4\x63\x8d\x62\xac\xca\xc5\xba\xd8\xae\xd6\x9f\x24\x63\x10\x0a\x26\x63\x5b\x9b\xf4\xde\x44\xb1\x35\x9a\x0a\xad\xb9\xf1\xb2\xfc\xfd\xae\x1f\x01\x00\x00\xff\xff\xef\x62\xfc\xb9\x2e\x03\x00\x00")

func manifests02DeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests02DeploymentYaml,
		"manifests/02-deployment.yaml",
	)
}

func manifests02DeploymentYaml() (*asset, error) {
	bytes, err := manifests02DeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/02-deployment.yaml", size: 814, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsImageReferences = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8d\xb1\x0e\xc2\x30\x0c\x44\xf7\x7c\x85\x95\x3d\xad\x58\x33\xb3\x30\x23\xb1\x5b\xad\x1b\xac\x36\x71\x64\x1b\xbe\x1f\xb5\x45\x4c\x4c\x77\xba\x93\xde\x5b\xb9\xcd\x19\x6e\x15\x0b\xdd\x5d\x09\x6b\xc0\xce\x0f\x52\x63\x69\x19\x78\xdf\x07\xe9\xd4\xec\xc9\x8b\x0f\x2c\xe3\xfb\x12\xac\xd3\x94\x03\x80\x63\xb1\x3d\x13\x34\xac\x94\x61\xda\x5e\xe6\xa4\x69\x6e\x96\xa4\x93\xa2\x8b\x06\x00\x80\x45\xa5\x66\x38\x2a\xc0\x69\x8c\x57\x99\x56\xd2\x43\x1c\xbf\xcf\x49\x89\x3f\xdd\x28\xca\x85\x5b\xfa\xc7\xcd\x1b\x3a\x99\xc7\xf0\x09\x00\x00\xff\xff\x37\x7c\x86\x26\xc1\x00\x00\x00")

func manifestsImageReferencesBytes() ([]byte, error) {
	return bindataRead(
		_manifestsImageReferences,
		"manifests/image-references",
	)
}

func manifestsImageReferences() (*asset, error) {
	bytes, err := manifestsImageReferencesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/image-references", size: 193, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x31\x4b\x04\x41\x0c\x85\xfb\xfc\x8a\x70\xfd\xae\xd8\x49\x3a\xb5\xb7\x38\xc1\x3e\x37\x93\xf3\xe2\xed\x66\x96\x49\x66\x41\x7f\xbd\x0c\x83\x20\xa8\x60\x97\xe2\xbd\xf7\x7d\xb9\xaa\x65\xc2\xc7\xa5\x79\x48\x3d\x96\x45\x1e\xd4\xb2\xda\x2b\xf0\xa6\x2f\x52\x5d\x8b\x11\xd6\x13\xa7\x99\x5b\x5c\x4a\xd5\x0f\x0e\x2d\x36\x5f\xef\x7c\xd6\x72\xb3\xdf\xc2\x2a\xc1\x99\x83\x09\x10\x11\x8d\x57\x21\x4c\x63\x6f\xca\xe6\x94\xcd\xc1\xdb\xe9\x4d\x52\x38\xc1\x84\x83\xf8\x2c\x75\xd7\x24\xf7\x29\x95\x66\x01\x5f\xc5\x1e\x1e\xb7\x6f\x9c\x84\xb0\x6c\x62\x7e\xd1\x73\x4c\xdf\x36\xa1\x96\x45\x8e\x72\xee\xc8\x1f\x0f\xc0\x5f\x12\xff\xd8\x6d\x2e\xf5\xa9\x87\xba\xe8\xc1\xdf\x3d\x64\x25\x1f\xaa\x3c\x54\xe9\xd7\x66\x07\x1c\xe0\x33\x00\x00\xff\xff\xdb\x7c\x21\xe3\x4d\x01\x00\x00")

func assetsDnsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleBindingYaml,
		"assets/dns/cluster-role-binding.yaml",
	)
}

func assetsDnsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role-binding.yaml", size: 333, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8d\x3d\x6b\xc4\x30\x0c\x40\x77\xff\x0a\x71\x7b\x52\xba\x15\xaf\x1d\xba\x77\xe8\xae\xb3\x05\x27\xe2\x48\x46\x92\x53\xe8\xaf\x2f\x89\xb3\xbd\xf7\xd0\xc7\xc6\x52\x33\x7c\xb6\xe1\x41\xf6\xad\x8d\x12\x76\xfe\x21\x73\x56\xc9\x60\x4f\x2c\x2b\x8e\x78\xa9\xf1\x1f\x06\xab\xac\xdb\x87\xaf\xac\x6f\xc7\x7b\xda\x29\xb0\x62\x60\x4e\x00\x82\x3b\x65\x28\xf3\xcc\x52\xc5\x73\x15\x4f\x36\x1a\x79\x4e\x0b\x60\xe7\x2f\xd3\xd1\xfd\x9c\x5d\xe0\xf1\x48\x00\x46\xae\xc3\x0a\xdd\x8d\xa4\x76\x65\x09\xbf\xcc\xc9\x0e\x2e\x34\xa5\x6b\x9d\x70\x7e\xf1\x8e\xb3\x1f\x64\xcf\x7b\xb7\xb1\xc7\x05\xbf\x18\xe5\x95\xfe\x03\x00\x00\xff\xff\x8e\xf7\xdc\x36\xd4\x00\x00\x00")

func assetsDnsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleYaml,
		"assets/dns/cluster-role.yaml",
	)
}

func assetsDnsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role.yaml", size: 212, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsConfigmapYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8c\x4f\x4b\x03\x31\x10\xc5\xef\xf9\x14\x0f\x3c\x77\x6b\x59\x56\x30\xd7\x9e\xf5\xe8\x7d\x4c\x66\x9b\xd0\x6c\x12\x66\x92\xa2\xa8\xdf\x5d\x6a\x75\xb5\xe0\x3b\xbd\x7f\xfc\x8e\x31\x7b\x8b\x7d\xc9\x73\x3c\x3c\x50\x35\x54\xe3\x13\x8b\xc6\x92\x2d\x4e\x3b\xb3\x70\x23\x4f\x8d\xac\x01\x6e\xf0\x48\x0b\x23\x2a\x94\x1b\xa8\x41\x7a\x6e\x71\x61\x03\x64\x5a\x58\x2b\x39\xb6\x28\x95\xb3\x86\x38\xb7\x8d\x4b\x5d\x1b\xcb\xc6\x67\x35\x3f\x8c\x7d\x11\x9e\x63\x62\x8b\x77\x03\x00\x83\x9d\xc6\x69\xc4\xdb\x57\x38\x8b\x45\x8a\xe8\x1a\x03\x53\x6a\x61\x8d\xc7\xfe\xcc\x92\xb9\xb1\xe2\x9b\x3e\xa4\xe2\x28\x21\xe6\x0d\x79\x2f\x03\x49\x25\xc4\x7a\x77\x31\xbf\xd8\xb3\x6a\xf1\x8a\x98\x95\x5d\x17\xbe\x5a\x7a\xd5\x26\x4c\xcb\x55\x39\x53\x4a\x2d\x48\xe9\x87\xf0\x3f\x7e\x7d\x7f\xac\xae\x4a\x59\xb8\x05\xee\x0a\x7b\xbf\x9b\xc6\xbf\xc3\xcb\x2b\x06\x6c\xb9\xb9\xad\xb0\x96\x74\x1a\x5c\xc9\xf3\x7a\x70\xe4\x02\x63\xbc\x5d\x0b\xe1\x54\xc8\x9b\x0b\xff\x33\x00\x00\xff\xff\x6c\x54\x72\x1f\xa6\x01\x00\x00")

func assetsDnsConfigmapYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsConfigmapYaml,
		"assets/dns/configmap.yaml",
	)
}

func assetsDnsConfigmapYaml() (*asset, error) {
	bytes, err := assetsDnsConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/configmap.yaml", size: 422, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsDaemonsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x4d\x6b\xdb\x4c\x10\xbe\xfb\x57\x0c\x7e\xaf\xaf\xe2\x98\x90\x92\xee\xad\xc4\xd0\x16\x9a\x54\x10\xb7\x97\xd2\xc3\x64\x35\xb2\x96\xec\xee\x2c\x3b\x23\x83\xff\x7d\x91\x6d\x29\x2b\xea\x06\xa2\x93\xd8\xe7\x63\x66\x1e\xcd\xea\xc5\xc5\xc6\xc0\x06\x29\x70\x7c\x22\x5d\x60\x72\x3f\x29\x8b\xe3\x68\x00\x53\x92\xd5\x7e\xbd\x08\xa4\xd8\xa0\xa2\x59\x00\xfc\x07\x8f\x18\x08\x9c\x80\x90\x02\x2a\xe4\x3e\xaa\x0b\xb4\x00\x88\x18\x48\x12\x5a\x32\xc0\x89\xa2\x74\xae\xd5\xca\xfa\x5e\x94\x72\xd5\x44\x59\x00\x78\x7c\x26\x2f\x83\x0f\x14\x1c\x4c\xc9\xc0\x40\x90\x44\x76\x00\x85\x3c\x59\xe5\x7c\x22\x06\x54\xdb\x7d\x2b\x94\x17\xb5\x00\x4a\x21\x79\x54\x3a\xab\x8a\xa6\x87\xc7\xcf\x0c\xfe\x61\x01\x30\xb6\x70\x7c\xa7\xbc\x77\x96\x3e\x59\xcb\x7d\xd4\x61\xee\x57\x1e\x80\xe5\xa8\xe8\x22\xe5\xc9\xb4\x3a\x46\x50\x72\x00\x5c\xc0\x1d\x19\x58\x36\x6c\x5f\x28\x5f\x39\x5e\x4d\x85\x57\x9c\xdd\xce\xc5\xca\x72\xa6\x26\x8a\xd9\xdf\x5c\xad\xd7\x57\xd7\xcb\xb9\xb6\xee\xbd\xaf\xd9\x3b\x7b\x30\xf0\xb5\x7d\x64\xad\x33\x09\x45\x9d\x58\x96\x43\xc0\xe1\x23\xfe\x82\xe5\xd9\x6a\x09\xbf\x27\x18\xf3\x4e\x8e\x58\x65\x39\xb6\xcb\xff\x61\xb9\x22\xb5\xab\x33\x73\x75\xcf\x99\x5a\xe7\xa9\x94\xec\xd9\xf7\x81\x1e\x86\xa1\x8b\xc0\xc6\xe9\x06\x1b\xb7\xab\x4e\xa4\x09\x05\x08\x03\xbf\x46\xed\x0c\x94\x15\x0a\x46\x26\x6c\xbe\x47\x7f\x30\xa0\xb9\x7f\x95\x26\xce\xf3\x3a\x53\xb2\x35\x67\x35\x70\x7b\x73\x7b\x53\xb8\xfc\x9d\x31\x40\xca\xac\x6c\xd9\x1b\xf8\xb1\xa9\xdf\xef\x54\xa9\x4d\x17\xdd\xb6\xf7\x6f\xb8\x7d\x5c\x5f\x70\x0b\xa4\xd9\xd9\xcb\xbd\x95\x6e\xde\xed\x29\x92\x48\x9d\xf9\x99\x4c\x41\xef\x54\xd3\x67\xd2\xf2\x08\x20\x9d\x62\xed\x08\xbd\x76\x73\xe4\xd8\xca\xdd\xf5\xdd\xf5\xec\x58\x6c\x47\x43\x3b\x5f\xb6\xdb\xba\x00\x5c\x74\xea\xd0\x6f\xc8\xe3\xe1\x89\x2c\xc7\x46\x0c\x7c\x28\xa5\xc3\x5d\xe6\x5e\x27\xf0\xb6\xc0\xa4\xb7\x96\x44\xb6\x5d\x26\xe9\xd8\x37\x06\xd6\x05\xda\xa2\xf3\x7d\xa6\x02\x1d\xb5\x4d\x94\x71\x83\x37\xd4\x62\xef\xc7\xe5\x3d\xed\xd0\x3b\x76\xec\x74\xfe\x80\x69\x1e\xcf\x1b\x3f\xa5\x62\x76\xa5\x20\x73\x5d\x05\x2f\x74\x30\x30\xde\x81\x19\x36\x86\x3e\x81\x7f\x02\x00\x00\xff\xff\x6f\x48\xea\xb4\x2a\x05\x00\x00")

func assetsDnsDaemonsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsDaemonsetYaml,
		"assets/dns/daemonset.yaml",
	)
}

func assetsDnsDaemonsetYaml() (*asset, error) {
	bytes, err := assetsDnsDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/daemonset.yaml", size: 1322, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xca\x31\x0e\x82\x21\x0c\x06\xd0\x9d\x53\xf4\x02\x0c\xae\xdd\x3c\x83\x89\x7b\x53\x3e\x63\xa3\x14\x42\x0b\xe7\x37\x26\xff\xf6\x86\xf7\x31\x6f\x4c\x0f\xac\x63\x8a\xbb\xea\xd8\x9e\x45\xa6\x3d\xb1\xc2\x86\x33\x9d\x5b\xe9\x48\x69\x92\xc2\x85\xc8\xa5\x83\xa9\x79\x5c\x8e\x29\x0a\xa6\x31\xe1\xf1\xb6\x57\x56\xfd\xee\x48\xac\xfa\x2f\xbf\x00\x00\x00\xff\xff\x35\xeb\xbe\x6a\x5d\x00\x00\x00")

func assetsDnsServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceAccountYaml,
		"assets/dns/service-account.yaml",
	)
}

func assetsDnsServiceAccountYaml() (*asset, error) {
	bytes, err := assetsDnsServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service-account.yaml", size: 93, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\xbd\x4e\x03\x41\x0c\x84\xfb\x7d\x8a\x91\xa8\x0f\x09\x21\x9a\x6d\xa1\xa1\x41\x27\xf1\xd3\x3b\x7b\x43\x58\xe1\xfd\xd1\xda\x09\xe2\xed\x51\x2e\x10\xa0\x40\x94\x1e\x7f\xfe\xc6\xaf\xb9\x2e\x11\xf7\x1c\xfb\x9c\x18\xa4\xe7\x27\x0e\xcb\xad\x46\xec\x2f\x42\xa1\xcb\x22\x2e\x31\x00\x67\xb8\x93\x42\x64\x83\xd1\x21\x8e\xb1\xab\x9e\x0b\x03\x50\xa5\xd0\xba\x24\x46\xb4\xce\x6a\x2f\xf9\xd9\xa7\xa4\x3b\x73\x8e\x69\xa9\x16\x00\x95\x0d\xd5\x0e\x1e\xfc\x60\xa4\xf7\x88\x03\x60\x9d\xe9\x58\xf2\x79\x76\x3b\xe3\x2d\xab\x62\x43\xc8\xce\x5b\x11\xcf\x49\x54\xdf\x51\xa4\xca\x96\xcb\x79\x00\x8c\xca\xe4\x6d\xfc\x69\x05\x7a\x1b\xbe\xb6\x4e\xeb\x93\x5f\xf1\x71\x11\x71\x75\xb9\x0e\x2e\x63\x4b\x9f\xd7\xe8\x04\x8c\xe6\x2d\x35\x8d\x78\xbc\x99\x7f\x0b\x26\x4f\xfd\x5f\xc9\x37\x74\x12\x3d\x5c\xcf\xe1\x23\x00\x00\xff\xff\xd5\x5c\x70\x51\x6f\x01\x00\x00")

func assetsDnsServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceYaml,
		"assets/dns/service.yaml",
	)
}

func assetsDnsServiceYaml() (*asset, error) {
	bytes, err := assetsDnsServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service.yaml", size: 367, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"manifests/00-cluster-role.yaml": manifests00ClusterRoleYaml,
	"manifests/00-custom-resource-definition.yaml": manifests00CustomResourceDefinitionYaml,
	"manifests/00-dns-namespace.yaml": manifests00DnsNamespaceYaml,
	"manifests/00-operator-namespace.yaml": manifests00OperatorNamespaceYaml,
	"manifests/00-role.yaml": manifests00RoleYaml,
	"manifests/01-cluster-role-binding.yaml": manifests01ClusterRoleBindingYaml,
	"manifests/01-role-binding.yaml": manifests01RoleBindingYaml,
	"manifests/01-service-account.yaml": manifests01ServiceAccountYaml,
	"manifests/02-deployment.yaml": manifests02DeploymentYaml,
	"manifests/image-references": manifestsImageReferences,
	"assets/dns/cluster-role-binding.yaml": assetsDnsClusterRoleBindingYaml,
	"assets/dns/cluster-role.yaml": assetsDnsClusterRoleYaml,
	"assets/dns/configmap.yaml": assetsDnsConfigmapYaml,
	"assets/dns/daemonset.yaml": assetsDnsDaemonsetYaml,
	"assets/dns/service-account.yaml": assetsDnsServiceAccountYaml,
	"assets/dns/service.yaml": assetsDnsServiceYaml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"dns": &bintree{nil, map[string]*bintree{
			"cluster-role-binding.yaml": &bintree{assetsDnsClusterRoleBindingYaml, map[string]*bintree{}},
			"cluster-role.yaml": &bintree{assetsDnsClusterRoleYaml, map[string]*bintree{}},
			"configmap.yaml": &bintree{assetsDnsConfigmapYaml, map[string]*bintree{}},
			"daemonset.yaml": &bintree{assetsDnsDaemonsetYaml, map[string]*bintree{}},
			"service-account.yaml": &bintree{assetsDnsServiceAccountYaml, map[string]*bintree{}},
			"service.yaml": &bintree{assetsDnsServiceYaml, map[string]*bintree{}},
		}},
	}},
	"manifests": &bintree{nil, map[string]*bintree{
		"00-cluster-role.yaml": &bintree{manifests00ClusterRoleYaml, map[string]*bintree{}},
		"00-custom-resource-definition.yaml": &bintree{manifests00CustomResourceDefinitionYaml, map[string]*bintree{}},
		"00-dns-namespace.yaml": &bintree{manifests00DnsNamespaceYaml, map[string]*bintree{}},
		"00-operator-namespace.yaml": &bintree{manifests00OperatorNamespaceYaml, map[string]*bintree{}},
		"00-role.yaml": &bintree{manifests00RoleYaml, map[string]*bintree{}},
		"01-cluster-role-binding.yaml": &bintree{manifests01ClusterRoleBindingYaml, map[string]*bintree{}},
		"01-role-binding.yaml": &bintree{manifests01RoleBindingYaml, map[string]*bintree{}},
		"01-service-account.yaml": &bintree{manifests01ServiceAccountYaml, map[string]*bintree{}},
		"02-deployment.yaml": &bintree{manifests02DeploymentYaml, map[string]*bintree{}},
		"image-references": &bintree{manifestsImageReferences, map[string]*bintree{}},
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

