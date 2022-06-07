// Code generated by vfsgen; DO NOT EDIT.

//go:build !gopherjsdev
// +build !gopherjsdev

package gopherjspkg

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// FS is a virtual filesystem that contains core GopherJS packages.
var FS = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2022, 6, 7, 12, 38, 17, 875426700, time.UTC),
		},
		"/js": &vfsgen۰DirInfo{
			name:    "js",
			modTime: time.Date(2022, 4, 19, 14, 12, 57, 881777300, time.UTC),
		},
		"/js/js.go": &vfsgen۰CompressedFileInfo{
			name:             "js.go",
			modTime:          time.Date(2022, 4, 19, 14, 12, 57, 881777300, time.UTC),
			uncompressedSize: 11230,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x5a\x5f\x93\xdb\x36\x92\x7f\x26\x3f\x45\x87\x95\x2a\x8b\xb6\x22\x6d\x12\x97\x6b\x6b\x72\xf3\xe0\x24\xb7\x3e\xe7\x62\x6f\x6a\x27\xbe\x3c\xb8\x5c\x2e\x88\x6c\x4a\xf0\x50\x00\x17\x00\x25\x6b\x67\xe6\xbb\x5f\x35\x1a\x20\x29\x91\x1a\x7b\xd6\x79\x58\xbf\x58\x43\x34\xba\x7f\xfd\x07\xfd\x07\xe4\x72\x09\xbf\x89\xe2\x5a\xac\x11\x3e\x58\x68\x8c\xde\xc9\x12\x2d\x54\xad\x2a\x9c\xd4\xca\x42\xa5\x0d\x48\xe5\xd0\x88\xc2\x49\xb5\x86\xbd\x74\x1b\x50\xc2\xc9\x1d\xc2\x2f\x62\x27\xae\x0a\x23\x1b\x07\xcf\x7f\x7b\x69\x17\xf0\x93\xa8\x6b\x0b\x4e\x83\xdb\xa0\xc5\x01\x17\x61\x10\x9c\x41\xe1\xb0\x04\xdb\x60\x21\x45\x5d\x1f\x60\x75\x80\x17\xba\xd9\xa0\xf9\xe5\x0a\x84\x2a\xc1\x19\xa1\x6c\xed\x89\x4a\x69\xb0\x70\xf5\x21\x30\x93\x06\x0a\x6d\x0c\xda\x46\xab\x92\x60\x0c\x44\xdb\x83\x72\xe2\xe3\x22\x5d\x2e\xd3\xe5\x12\xde\x58\x84\x57\xe2\x1a\xff\x30\xa2\x69\xd0\xd0\x7e\xfc\xd8\x68\x8b\xb0\x45\xb7\xd1\xa5\x87\xd7\xef\x5e\x74\x1b\xfe\xd6\xd6\xf5\xf9\x4d\xcf\x5f\xff\x0c\x95\xc4\x7a\xbc\xff\x8f\x0d\x2a\x68\x84\xb5\x04\x6b\x27\xea\x16\x6d\x87\x7e\x4e\xd8\xa1\xd2\x75\xad\xf7\xb4\xec\x0e\x0d\x42\xa1\xd5\x0e\x8d\xed\xec\xd2\xa0\xa9\xb4\xd9\x62\x79\x11\x54\x80\x5b\x78\xa1\x99\xf6\xf8\xdf\xed\x50\xed\xc1\xfa\x2d\xfc\x34\xe0\xb9\x12\xc5\x35\x81\xf4\x5e\xab\x44\x81\x37\x77\x70\x1b\xf8\x7e\x33\xf5\xef\xa1\xcf\x87\x14\x81\xef\x4a\xeb\x1a\x46\xff\x6e\xe1\x47\xad\x6b\x14\x6a\xf4\x7c\x9a\x7e\x40\x11\xf8\x92\x0e\x6b\x34\xd6\x87\x47\x55\x6b\xe1\xac\xdf\xff\xba\xdd\xae\xd0\x8c\xe5\x79\x92\x67\x4f\x3f\xc9\xd7\x3a\x43\xfe\x18\xed\xbf\x3a\xf3\x7c\x9a\x7e\xcc\xf7\xed\x3b\xa9\xdc\x5f\xc7\xfb\x5f\x2a\xf7\xd7\xe7\xc6\x88\xc3\xc9\xf3\x69\xfa\x33\x7c\xbf\x7d\x36\xc5\xf7\xdb\x67\x23\xc6\xe7\xe8\xcf\xf0\xfd\xfe\xbb\x39\xff\x38\xe2\xfb\xfd\x77\xe7\xf8\xc2\xe7\xe0\x6d\x27\x14\xbb\x85\x37\x72\xca\x10\xe7\xe8\xcf\xf1\x3d\x55\x8c\xf9\x8e\x0d\x71\x8e\xfe\x1c\x5f\x36\x44\xdb\xa9\xc8\x7c\xc7\x86\xb8\x3d\xa2\xba\x9f\xaf\x8f\xc8\xef\xbf\x3b\xc1\xfb\x37\x7e\x7a\xc2\xf8\x1c\xfd\x59\xbe\x27\x91\x1e\xf8\x3e\x7b\x7a\x8e\xef\xd9\x93\x11\xf9\x8a\xba\x06\xed\x36\x68\xc0\xd6\xb2\x40\x1b\xf7\x8f\x63\x77\x10\x0f\x5d\x96\xb9\x87\x2f\xed\xb7\x13\xe7\x0a\x91\x25\x1d\xa5\xbb\x73\xcf\xc7\x7c\xfb\x0a\x73\x62\x87\xf0\x7c\x94\x1f\x5a\x55\xcc\x16\x8b\xc5\x00\x75\x0e\x8f\x3f\xd8\xc5\xdf\x57\x1f\xb0\x70\x1d\x5f\x27\xb7\xb8\xf8\x5d\x6e\xf1\x64\xff\xcf\xc2\x4d\xa1\x39\x43\x3f\xc6\xfb\xcd\xf4\x2a\x48\x65\x9d\x50\x05\xea\x0a\x5e\xeb\xb2\xcf\xeb\x03\x68\xf7\xf2\xdd\x8a\xc6\xce\x29\x4b\xb5\x85\xb3\xd3\x7c\x07\x6c\x3c\xfd\x5b\xce\x69\xd3\x0e\xbc\x0d\xa5\xe8\x79\x59\x4a\xb2\x23\x95\xeb\xb9\xef\x05\x44\x90\x42\x65\xcc\x09\xa9\x28\x2d\x8a\x21\x4e\x5f\x25\xe7\xa0\x15\x15\xef\x8d\x2f\x77\x0e\x95\x03\x5d\x71\x31\xa4\x65\xd8\xcb\xba\x86\x15\xfa\xba\x89\xe5\x71\x49\xf5\xb9\x7e\x47\xbe\xa7\x92\x26\x16\x69\xd3\x35\x28\x29\x61\x0a\x72\xa4\x05\x11\x41\xa0\x09\xd8\xc6\x8d\x89\xf6\xd4\x83\xd6\x44\x3a\xdb\x55\xf5\x3f\xa1\x2d\x19\x37\x22\xf0\x1c\x94\xac\xa1\xd1\xde\xb2\x44\xd9\x23\xc6\x7f\xb6\xa2\x3e\x56\xf7\x91\x85\x4c\xb5\x75\x9d\x2d\x22\x5d\x21\x14\x28\xed\xc8\x3e\x2d\x59\x47\x90\xa6\x5b\xd1\xc0\x35\x1e\x16\xa9\x3f\x10\x81\x92\x5d\x71\x13\x94\x84\xc7\xe1\xf1\x9d\xb7\xd3\x0b\x74\x60\xd0\xb5\x46\x59\x6f\x79\x26\x7a\xe4\xbb\xbc\x06\x8d\x3b\x70\x2f\x47\x4b\x6b\xb9\x43\xc5\xec\xe9\x84\xc0\x4c\x47\x5e\x39\xb1\x99\x5d\xe3\x21\x94\xc0\xbc\x13\x72\x13\x98\x83\x5e\x04\x1b\x07\xca\x3c\xc8\xbf\x42\x07\xd4\x16\xad\x83\x7c\xdf\x1b\x05\xc3\xfd\xbb\x60\xae\x8e\xc0\xcc\x03\xcf\xa3\xd3\x7c\xd3\x03\x0a\xd4\x81\x2c\xe2\xfa\x19\x6b\x74\x08\x06\xb7\x7a\x87\x5f\x64\x1a\xe6\x74\x64\x9d\x81\xf4\x7e\x35\x4a\xfe\x15\xd5\xda\x6d\xa6\x9d\x92\xd5\x7e\x31\xeb\x20\xcc\x43\xa3\xe8\xf8\x7c\x48\xe5\x26\x10\x30\xc7\x59\x4e\xcb\x13\x1e\xe9\x96\x59\xfe\x4b\x55\xe2\xc7\x23\xf1\xf2\x91\xdb\x00\xd6\xb8\x0d\x27\x54\x28\x4e\xd5\x13\xa2\xfc\xe6\x99\x24\x49\xf7\x05\x41\x20\x1b\x04\x01\x4b\xb5\xe8\x1e\x2c\x32\x6e\x66\xa9\x9f\xe1\xed\x40\x7d\xe2\x70\x3a\xfa\x50\xf0\xf9\x1f\x9a\x9c\xb3\xc0\xa9\xab\x95\xd8\xe2\x04\x16\x62\x32\xa3\xb5\x2e\xf6\x84\x59\x5b\x18\xd5\x92\xb3\x86\xe9\x18\xf0\xce\xc5\x62\xd1\xbb\x65\xa7\xaf\x71\x84\x90\x32\x15\xd6\xd5\x02\x7e\xdf\x48\xcb\x19\xb3\x12\xb2\x06\x59\x81\xf4\xc9\x84\x72\x84\xe8\x4a\xe0\xa4\xcb\x88\xf1\xec\x81\x40\x07\xbb\x06\x20\x5f\xe3\x1e\x0a\x9f\x2a\x29\x1b\x29\xdc\x77\xb5\x85\x33\xbb\xb4\x5c\xaa\x63\xbe\x9d\x04\x7d\x8c\x18\x66\x85\x56\x9c\xc2\xb4\xc9\x27\xf0\xbf\xc6\xfd\x43\xc1\xc7\x2d\x03\xe4\x34\x83\x4c\x9c\xb9\xe3\xe3\xe5\x07\x12\x51\x14\xda\xf8\xf1\xf2\xb8\x20\x9d\x8e\x6d\x13\x50\x49\xc8\x2c\x67\x36\x63\x54\x61\x35\x1c\x09\x9e\x25\x3e\x85\x28\x8c\x1c\x5f\x80\x89\x05\xcd\xf2\xc8\x6a\x8c\xab\xa3\x88\x81\xe8\x3e\x09\x8b\x12\xcd\xe7\x62\x82\x59\x23\x8c\xc5\x97\xca\xe5\x93\xd1\xe9\xce\x26\x2e\x5e\xeb\x50\x3d\x7b\xfa\x39\xb8\x9e\x3d\xfd\xf3\x90\x3d\x7b\xca\xd8\x9e\x3d\x9d\x46\xe7\xd7\x19\xdf\x1b\xf9\x59\x00\xdb\x3f\x13\x21\xcb\x9c\xe5\x91\xeb\x18\x63\x47\xc1\x20\xfd\x60\xf0\x49\x8c\x71\x48\x78\x20\x48\xcf\x7c\x0a\xa6\x5f\x98\xe5\x1d\xdf\x31\xcc\x48\xd1\xb9\x9a\x0f\xf9\xe7\xb8\x3b\xa6\x83\x05\x5c\x21\x82\x13\xab\x9a\x6a\x03\xc4\x6e\xb1\xd0\x5b\x5f\x62\xa8\x31\x2c\xd1\x09\x59\xdb\x69\x57\x33\x1f\x76\x77\xd7\x09\x4f\x3a\xbd\xa3\x0c\x8e\x57\x56\x54\x93\x50\xa9\x63\x53\xde\x37\x8d\x33\x73\xd8\x6f\x64\xb1\xf1\x6d\xdd\x0a\x07\x6a\xec\xa4\x80\xd6\xf3\x58\xfc\xc6\xcd\xe2\x02\x5e\x6b\xe7\x71\xa8\x12\x4b\x0f\xbd\x69\x57\xb5\x2c\xa8\x11\x9c\x0a\x03\xbf\x3b\x84\x41\xe3\xcc\x54\x1c\x44\x12\xc6\xfc\xdf\xc6\x68\x03\xa8\x0a\xd1\xd8\xb6\xf6\xd9\x7c\xe0\x5f\xa4\x55\x4b\xc9\x5b\x5b\xe4\xee\xb8\x35\x0a\x4b\x82\xa4\x41\xc0\x0b\x0d\x8d\x50\xb2\xf0\x6d\xf1\x56\x1c\x48\x1f\x83\x85\xde\xa1\xc1\x72\x4e\x05\xd4\xa7\x2c\x05\x8f\x59\x8e\xdb\x08\x07\x1b\xed\x6f\xcd\x36\x38\x92\x14\x8b\x05\xf7\xb4\xbc\x25\x4c\x17\x37\x69\x12\xb4\x4c\x87\xc0\x87\xb6\xde\xa2\xb5\xe4\xe8\x30\x58\x0c\x74\x2a\xcf\x4b\x62\x13\xa2\x31\x01\x62\xce\x8c\x07\x49\x32\x4d\x82\x09\xb3\x53\x26\x17\x90\xc1\x13\xfa\xe9\x3b\xdd\x2c\xc8\xcf\xf2\x2e\x8d\xa6\x31\xc1\x8b\xe2\xfa\x08\xaa\xf5\x4f\xba\xe6\xf2\x0b\x11\x7b\xfe\x53\x88\x3b\x68\x5e\xde\x18\xd8\x8b\x5a\xaf\x44\xed\xfb\x1c\x7b\x3c\x81\xac\x79\x25\x84\xef\x2c\xdb\x4b\x55\xea\x7d\xe6\x23\x70\x65\xf4\xde\xc6\x3b\xb8\xec\xc5\xaf\x7f\xff\xf1\xf9\xaf\xbc\x42\xa3\xea\xe2\x83\xcd\x17\xe9\x4e\x98\xc8\x3d\xba\x8d\x04\xbe\xd2\x65\x5b\x63\x10\xd8\xcf\x00\x41\xff\x6c\xeb\x97\x33\xd8\x09\x23\xfd\xf1\xb5\xe8\x68\xfa\x0a\x7c\x17\xf0\x3f\x52\xb9\x0b\x1e\x24\x88\x1d\xd3\xfb\xab\x59\xe3\xb8\x6f\x7b\xf4\xc1\x2e\x58\x0a\x6b\xce\x6b\x96\x74\xef\xff\x7c\x2d\xb6\x98\xcd\xa9\x8b\xc8\x1f\xc5\x7b\xe2\xd7\xda\x21\xc7\x67\xc7\x81\x7a\x2a\x3f\xb6\x96\x58\x49\x8e\x7a\x30\xad\xa2\xd9\xde\x86\x33\x6c\xdb\xc6\xcb\xfe\x49\x6f\xb7\x5a\xfd\x72\xd5\xa3\xb2\x30\xdb\x38\xd7\xd8\x8b\xe5\x52\xe9\x12\x3f\xd8\x85\x36\xeb\xa5\x68\xe4\x32\xac\x2f\x36\x6e\x5b\xe7\x0b\xaf\xdc\x2f\x57\x91\x93\xf5\x6d\x91\x9f\x5a\xeb\xc3\x9c\xd8\xad\x5a\xca\x00\xbd\xd5\x25\x0f\x84\x1e\x58\x9c\x08\x65\xd5\x4f\xa8\xba\x75\x4d\xeb\xfb\xc1\x38\x4c\x6f\x8c\x6e\xd7\x1b\x36\xd9\xaa\x55\x65\x8d\x26\xc0\x97\xdb\x86\x1b\x6f\xdb\x69\x00\x33\xf2\x24\x7e\x14\xb4\x34\x87\x3d\xae\x28\x81\x02\x3d\xb3\xab\x56\xd6\x65\xf0\x6e\x30\xd1\xd0\xbb\x6f\x54\x34\x54\xef\xe0\x41\x18\xb3\xaf\xb3\x36\x52\x65\xcc\xa8\xdf\x35\xe4\xf5\x33\xae\xda\xf5\x1a\x0d\xac\x69\x4e\x28\xf4\xb6\x91\xf5\xe9\xc5\x00\x4d\x49\x65\xa0\xfb\x21\xa3\x43\xe5\xbc\x32\xe1\x8c\x44\x16\xb3\x1c\x6e\x06\xe5\x44\x89\x3a\x74\x8b\x47\x83\x4f\x58\x1a\x5f\x15\x70\x50\x18\x6c\x0c\x5a\x6f\x29\xf9\x39\x59\xf9\x58\x14\x0f\x2c\x13\xfd\x6a\x77\x54\x95\xac\xc3\xa1\xe4\x77\x0f\xaa\x80\xbd\x11\x8d\x1d\xb6\xc7\x74\xde\xd8\xb2\xa2\x28\xd0\xc6\x17\x2b\xf1\x25\x83\xae\x4e\x6c\x43\x4d\x78\xc6\xa7\x54\x98\x75\xeb\xfd\x9c\xd1\xe8\xba\xd7\xa6\x8c\xc5\x2f\x8a\x9b\x55\x8a\x6f\xc3\x7c\xeb\x1e\x00\xfa\xd1\x84\x37\xc2\xdb\x77\x5d\x99\xf9\x84\x2e\x7c\xf0\x79\xc0\xc9\xbe\xde\x06\x01\xd9\xfc\xd4\x28\x95\xca\x63\x26\xfa\x5f\x3c\xd8\x23\x7f\x5c\xd3\x83\x90\x17\x78\x0e\x1b\xdf\xe1\xb0\x02\xb4\x75\x58\x03\xdf\xbe\xeb\xf3\xa0\xac\x40\xc3\xe5\xa5\xbf\x7f\xb9\xbd\xe5\xdf\x7d\xbc\xdd\xa4\xc9\xd0\xfc\xc9\x5d\x9a\x08\xb8\xb8\x8c\xf8\x7d\xfe\x60\xae\x59\x1e\xb4\x21\x58\xd9\x1c\x74\x9e\x26\x96\x48\x49\xb9\x59\x94\x38\x07\xd1\x4d\xd8\x79\x9a\xf8\x37\x65\x44\xf4\x97\x1f\x40\xc2\x7f\x0d\x16\x7f\x00\xf9\xe4\x89\x17\x6f\xdf\xca\x77\x70\x09\xa2\x1b\x93\xfb\x14\x4d\x70\x02\x3a\x3b\x08\x8d\xf8\x4a\xaa\x9f\xbd\xc6\x11\xcb\x87\x7b\x23\xac\x8f\xa1\x86\xb2\x46\xe5\xab\x6f\xcc\x95\x58\x76\x57\x5e\xba\xa2\x80\x7e\x63\xfd\x52\x2d\x0b\xe9\xe8\xc8\x39\x34\x3e\x70\x2c\xff\x1c\xbc\x2a\x0b\xef\xc1\x42\x59\x9e\x7c\x05\xd6\x07\x56\x00\x7b\x4f\xf8\xef\xc8\x40\xa7\x87\x25\x4f\x13\x7d\xd6\x11\x34\xd1\x11\x01\x27\xf4\xf7\xef\xe3\xc9\x7d\xcf\xca\xbf\x7f\x9f\xcd\x61\x97\xa7\x49\xc4\x7c\x71\x09\x3b\x66\x31\x98\x2e\xb3\x3c\xd6\x6c\x4f\x94\x4d\xb8\x2b\x2c\x4d\x38\x6d\xeb\x3d\x1f\x96\xa3\xe3\xd2\x84\xa2\x6d\xcb\x6c\x9b\xeb\xf5\xa0\xda\xc2\x57\x97\x90\x65\x70\x03\xcb\xa5\x9f\x78\xa3\x0f\xd2\x24\x49\x0a\xad\x9c\x54\x2d\xa6\x09\xf9\x3b\x68\x15\xb8\x28\x2a\x53\x3d\x9b\x39\x9f\xcf\x38\x00\x77\x01\x3f\xb0\x66\x32\x7d\x04\xf1\x23\x9b\x48\xfe\x0b\xe3\x45\x38\x19\xc9\x4b\x89\x88\x8d\x6e\x06\xb2\xf2\x79\x54\xc5\x1d\x9a\x2c\x9f\x83\x33\x2d\xc6\x43\x20\x9a\xa6\x3e\x10\x03\xbe\xb9\x20\xd5\xef\x8e\xe2\x55\x1f\xa5\xb2\xfe\x35\xea\x17\xc6\xac\x2f\xae\x83\xb0\x9d\x53\x88\x52\x37\x8d\x06\x41\xf2\x05\xf0\x6c\x70\xcd\x2a\xf2\x18\xa6\x3e\x43\xce\x03\xe7\x32\x04\xb8\x25\x7e\x7d\x90\xfb\x3f\xbb\x9a\x5d\xe2\x0e\x6b\x6a\xcf\x16\x5b\xfd\x2f\x59\xd7\xc2\x97\x6f\x54\xdf\xbc\xb9\x5a\x96\xba\xb0\xcb\x3f\x70\xb5\xec\xb5\x58\xfe\x03\x2b\x34\xa8\x0a\x5c\xb2\xe9\xdf\xb3\x53\xec\x92\xff\x5f\x72\xce\xf9\x2d\x74\x7c\x39\xc9\x8a\xea\x29\xad\xbe\xc1\xed\x0a\x4b\x2a\x26\xdd\xf9\x0c\x27\x8b\x8f\xe7\xff\x71\x86\xe7\xb4\x1f\x26\x05\x7e\xa5\x1e\xec\x11\x55\x09\x9a\x71\xab\xbe\xc1\xad\xc5\x1d\xb5\x22\x51\xf1\xfd\x06\x55\xc7\x65\xee\x5b\x0b\xa1\xa8\x0b\xd0\xc6\x09\xe5\xf8\x8e\x1a\x9c\x4e\x39\x52\x7d\x07\xe4\xcb\x1f\xdf\xf0\x44\x36\xe1\xde\xcd\x06\x87\x96\xa0\x15\xa0\x28\x36\x81\xf5\x51\x65\xe9\xbc\x7f\x4f\x12\x90\xfd\xf9\x3f\x93\x0e\x06\x47\x97\x28\x06\x1b\x26\x8e\x76\x9a\x26\x21\x86\x02\xc3\x7b\xf2\x48\x9a\x1c\x7b\x86\xc8\xfd\x31\x1b\xde\x2a\x97\x68\xbd\x97\xb5\x81\x57\xb9\x3f\x67\xf7\x94\x88\x63\x7e\x59\x8c\x3a\xc2\x32\x07\x7f\xfb\xdc\xb3\xf3\xa7\xe6\x14\xc2\xb9\xa4\xf6\x8a\x04\x67\xde\xf6\xd9\xc5\xd0\x04\xf3\x94\xce\x5f\x9a\xd0\x3a\x5f\x6f\x16\x7e\x88\x00\x01\x16\x95\x95\xd4\x49\xfb\x89\x8a\xf5\x59\xa4\x4c\xf7\x07\x42\xa9\xd5\x23\x07\x7b\xe1\x9d\x1e\xe2\x00\x84\x3a\xc4\xa1\xd9\x52\xe7\xe9\x1b\x82\xf0\x60\xce\x5b\xad\x86\x3d\xed\x06\xab\xbb\x0b\x50\x20\xf4\x82\x5f\xbe\xad\x0e\xb0\x11\xaa\xf4\x92\xdc\xa1\x21\xa3\x0e\x3c\x14\x67\x12\xda\x35\x1c\x4a\x92\xa4\xb9\x5e\x4f\xd2\x1e\xe7\x53\xe2\x4a\xc3\xed\x05\xa5\x55\xce\xbb\xee\xd0\xbc\xfd\xcb\x3b\x2a\xef\x8f\x1e\x3f\xe2\x4c\x48\x14\x97\x90\x3d\xce\x7c\x6a\x4d\x93\x51\x82\xaf\x51\xcd\xdc\xa1\x19\x24\xf6\xc8\x49\x32\xa7\x45\xe0\xe4\x55\xb8\xe4\x95\x27\xdf\x5e\xbc\xf3\xcf\x56\x06\xc5\x35\xfd\xba\x8b\xfc\x9b\xeb\xf5\xef\xac\x2b\xa9\xf1\x04\xb2\x05\x8d\x87\x04\xe3\x09\xed\x4d\x93\x91\x9f\xbf\x26\xaf\x44\xcf\xf6\xae\x65\x46\xf3\x2e\xad\xa6\x09\xf5\xc9\x21\x21\xc4\x26\x79\x58\xdf\x86\xd1\xe8\xdf\xcc\xf6\x65\x92\x6a\x92\x9d\xb4\x69\x57\xfa\x7e\x20\x8a\xaf\x4e\x1b\xa3\xc8\xbe\xaf\x74\x1c\xde\x85\x56\x85\x70\xd9\x1c\xb6\x96\x73\x3e\xf5\xd5\x15\x85\x03\xe5\x1c\xd1\xbd\xe6\x0a\x6f\x77\x7c\xc2\x29\xbb\x74\x56\x19\xbd\x8d\x97\xfd\x73\xbf\x17\x6b\x8b\xf1\xbd\x60\x77\xc4\xf9\xa6\x9b\xaf\x8b\x37\x62\xc7\xb9\x6c\xe1\xb5\xc1\x49\x65\x88\x25\x69\x82\x63\x45\x82\xe4\x4b\x08\x13\x21\xff\x4d\x15\xff\xd3\x3a\xe2\x89\xa9\x48\x63\x46\x7c\xc4\x79\x04\xa7\x97\x71\xf7\x1f\xd2\x58\x9c\x84\xde\x74\xd5\x1f\x45\xe2\x67\x34\x1c\x0f\xe9\x38\x4e\xd3\xf6\x03\x7a\x8f\xd1\xf0\x70\x52\x5d\xf2\xd3\xe6\x64\x98\x1f\xbb\x36\x25\xb9\xeb\x4f\x15\x59\x35\x78\x70\x1c\x33\x27\x2e\x63\xba\x09\x8f\x25\x95\x2f\x18\xbc\x3c\xf0\x18\x31\xff\xaa\x1a\x5e\x41\x60\x99\xe5\xf1\xde\x9f\x0d\x37\xf0\x90\x77\xd1\xa9\x8f\xaa\x7b\x7d\x94\x64\x6b\x74\xd1\x45\xa7\x3e\x49\x76\xc5\x20\x2f\x04\x9f\x14\xba\x39\xbc\xac\xfe\x81\xff\x6c\xa5\xc1\x72\xc2\x1d\xd9\xd7\x3b\x51\x87\xce\xb8\x3a\xe7\x9a\x6a\xe0\x9a\x9c\x85\x7d\x2a\x02\xa8\x55\x2c\x8e\x77\x7e\xda\x9d\x9e\xb5\x77\x57\x92\x64\xb6\x57\xf5\xc3\xae\x1f\xf5\x82\xb2\xeb\xdd\x58\xd9\xa8\x1b\x8b\xff\xb0\xfb\xb7\xc4\x27\xe7\x2c\x74\x75\xd6\x42\x73\x58\xef\x86\xd8\xef\x72\x3e\x80\x7d\x73\xdc\xb7\x03\x69\xf7\x26\xcd\x27\xed\x1f\xdb\xaa\x3a\xd7\x24\x0f\x09\x7c\x0e\x15\xb0\x3a\xb8\xf0\x4d\x4c\xe8\xb7\x8e\xf9\xcc\x56\xf0\xf6\x1d\xd1\x1c\xc5\x06\x7f\x43\x33\xee\xb1\x56\x34\x51\x55\x95\x45\x47\x8b\xcc\x95\x15\xe6\xa7\x59\xce\xaf\x60\xd2\x84\x5f\x4b\x9f\x52\x85\x97\xd5\x1d\x55\x1c\x5c\x07\x24\x22\x14\x26\xff\xd7\xca\x63\xec\x7a\x26\x4f\x47\x73\xb5\x17\x16\xff\x7f\xc2\x5c\xe3\x1d\xc1\x2b\xee\xf0\xad\xbf\xb4\xf2\xdf\x3f\x50\xf9\x5c\xc0\x4b\x7f\xd9\xd5\x5d\xc7\xf8\xaf\x23\xec\x46\x1b\xb7\xf1\x1f\x09\x6a\x33\x9e\x36\x2c\xcc\x56\x58\x69\x33\x7c\x79\x91\x87\x6b\xe7\x57\x67\x3e\x86\xe1\xab\xdc\x23\x0c\xfd\x17\x49\x0f\x44\x11\x3e\x7f\x3a\x0f\xe2\xea\xf8\x4b\xaa\x94\x3d\x2c\x95\x74\x9c\x3e\x96\x4b\x78\xbe\xd3\xb2\x84\x12\x45\x09\x85\x2e\x11\xb0\x96\x5b\xa9\x04\xbf\xfa\x4d\xbc\x93\xfd\xfd\xf0\xcd\x5d\x9a\xbc\xa7\xf2\x97\xde\xa5\xff\x1f\x00\x00\xff\xff\x9c\xe8\xac\x82\xde\x2b\x00\x00"),
		},
		"/js/js_test.go": &vfsgen۰CompressedFileInfo{
			name:             "js_test.go",
			modTime:          time.Date(2022, 4, 19, 14, 12, 57, 881777300, time.UTC),
			uncompressedSize: 371,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x7c\xcf\x31\x6f\xfa\x30\x10\x05\xf0\xd9\xf7\x29\x4e\x5e\x48\xfe\xff\xca\xde\x10\x04\x31\x31\xa0\xae\x2d\x3b\x5c\xdc\x4b\xe2\xd4\x24\x91\xcf\x61\x28\xe2\xbb\x57\x41\x55\xa2\x2e\xdd\xfc\x9e\x7e\xf2\xd3\x59\x5b\xf7\x45\x39\xfa\xf0\x81\xad\x80\xb5\xf8\x7f\x0e\x30\x90\xfb\xa4\x9a\xb1\x95\x73\x62\x49\x00\xfe\x3a\xf4\x31\x61\x06\x4a\x4f\x85\xef\x6a\x0d\xa0\x74\xed\x53\x33\x96\xc6\xf5\x57\x5b\xf7\x43\xc3\xb1\x95\xe5\xd1\x8a\x86\x1c\xa0\x1a\x3b\x87\x27\x96\xf4\xda\x25\x8e\x1d\x05\xff\xc5\x07\x1f\xdd\x18\x28\xbe\x71\xc5\x91\x3b\xc7\x59\xc2\x7f\x3f\x1f\x9b\x53\x8e\x77\x50\xd6\xe2\x3b\x33\x36\x29\x0d\x52\x58\xfb\xe7\x92\x17\x19\x59\xec\x76\xbd\x31\xa0\x5a\x31\xc7\xd0\x97\x14\xcc\x81\x42\xc8\x34\xdf\x28\xe8\x17\xbc\x80\xba\x51\xc4\x27\xdd\xae\x37\x84\x7b\xbc\x3f\x76\xbf\xcb\x72\x2a\x57\xb4\x2a\x16\x36\x91\x39\x98\x09\xcc\x78\x77\xc9\x41\x9d\x71\x8f\xcb\xe2\x91\x53\xa6\x67\xae\x73\xf3\xbc\xb9\x22\xc7\x59\x0e\x0f\xf8\x0e\x00\x00\xff\xff\x8a\xd5\xbd\x72\x73\x01\x00\x00"),
		},
		"/nosync": &vfsgen۰DirInfo{
			name:    "nosync",
			modTime: time.Date(2021, 8, 22, 11, 13, 56, 543823400, time.UTC),
		},
		"/nosync/map.go": &vfsgen۰CompressedFileInfo{
			name:             "map.go",
			modTime:          time.Date(2021, 8, 22, 11, 13, 56, 543823400, time.UTC),
			uncompressedSize: 1958,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x55\x4d\x8f\xdb\x46\x0c\x3d\x5b\xbf\x82\x3d\x55\x2e\x14\xe7\x9e\x62\x0f\x05\x7a\x29\xd0\x34\x40\xdb\x5b\x90\x03\x2d\x71\xac\x81\xe7\x43\x1d\x52\xeb\x2a\x8b\xfd\xef\x05\x39\xb2\x57\xde\x24\x45\x0f\xbd\xd9\x23\x0e\xf9\xf8\xde\x23\x67\xc2\xfe\x8c\x27\x82\x94\x79\x49\x7d\xd3\xbc\x7d\x0b\xef\x71\x02\xcf\x80\xd0\xe7\xd4\xcf\xa5\x50\x12\x88\x38\xc1\xc5\xcb\x08\x18\x73\x11\xff\x99\x86\x37\x7d\x4e\x2c\x98\xe4\x8d\xf8\x48\x10\x32\x0e\xdc\x01\x4b\x2e\xc4\x1d\x60\x1a\x60\xa0\x40\x42\x7c\xd0\x9c\xbf\x88\xa6\x64\x74\x04\x2e\x17\x88\x73\x10\x3f\x05\x82\x53\x2e\x79\x16\x9f\x88\x41\x32\xf4\x18\x02\xa0\x02\xf8\x9e\x21\x92\x8c\x79\xe0\x0d\x8a\xb0\x68\x2e\x4d\xf7\xe7\x48\xf0\x99\x4a\xbe\x62\x7d\xc4\xe0\x07\x2b\x4a\x71\x92\x5b\xd8\x4f\xf6\x3d\xce\x2c\x90\xb2\xc0\x91\xa0\xcf\x93\xa7\x01\xd0\x09\x15\x70\xbe\xb0\xc0\xcc\x74\x68\x64\x99\xc8\x82\x59\xca\xdc\x0b\x3c\x35\xbb\xa8\x4d\x7f\xf4\x49\xa8\x38\xec\xe9\xe9\xf9\xd3\xe6\x77\xf3\x6c\x54\xfd\x9a\x71\x80\x42\x32\x97\xc4\x20\x23\x29\x90\x99\x2a\x0b\x03\xf8\x64\x67\xca\x9d\x36\x8d\x70\xa6\xa5\x83\x5c\x20\xf9\x00\xde\x41\xca\x9a\xa3\x5e\xf1\x0c\x53\x21\xa6\x24\x87\x6b\x83\xf9\x0c\x85\x78\x0e\x02\x3e\x0d\xbe\x47\x21\x86\xcb\x48\x32\x52\x59\x2f\x5d\x90\xc1\xe5\x39\x6d\x4b\x1d\x1a\x37\xa7\x1e\xda\x08\x3f\xbc\xc7\x69\x6f\x10\xdb\x33\x2d\xb0\x41\xbf\x87\x76\xad\xfa\x72\xd6\x69\xbd\x63\xce\x61\xaf\xcd\xdb\x67\x3b\x7a\x80\x78\x88\x1f\xcf\xb4\x7c\x6a\x76\xb5\x53\xb8\x7d\x5c\x59\xf8\x43\xdb\x05\x26\xd9\x72\x70\xeb\xf8\x35\x20\x8b\x6e\x8d\x8a\x2f\x40\x58\x6d\xef\xb4\x24\x3c\x3c\x18\x4f\x4f\xcd\x6e\x67\x7f\x21\xe2\x99\xda\x7f\xd1\x64\xdf\xec\x9e\x9b\xdd\x15\x2d\x3c\xd4\xf4\x1b\xa5\x3e\x94\x8a\x74\x2b\x18\xfd\xed\x59\x7c\x3a\x6d\x50\xeb\xb1\x11\xe6\xee\x24\xf9\xa0\xc4\x5f\x3c\x53\x07\x5e\x56\xa3\x9b\xe5\xb6\xe9\x4e\xfe\x91\x56\x82\x6e\x3a\xea\x68\xd0\x70\xd3\x92\x41\x8a\x76\xed\x36\x64\xa9\x90\x35\xac\x03\x87\x81\xed\x73\x75\xd1\xd7\xf4\x5c\x1b\xf9\x26\x89\x2d\xf6\x32\x63\xb8\x97\x77\x85\x71\x93\xd8\xbb\x17\x21\xe1\xdd\x8b\xcc\x3f\xea\x7f\x65\xfd\x5e\x6d\x05\x6d\x04\xff\xcf\xf2\xbc\x2a\x63\xdd\xaf\x9a\xfd\x6c\x0b\xe4\xba\x47\xfe\x8b\xb7\xea\x8d\x2f\xed\xfe\x55\x57\xd5\xc2\x86\xaa\x96\x68\xe3\x21\x76\x9a\x76\xbf\x02\xf8\x1d\xd3\x89\x6c\x2b\x31\x38\x60\xfa\x6b\xa6\x24\x1e\x43\x58\x0c\x02\x61\x3f\x9a\x53\xd4\x05\x15\xd9\x6a\x98\xbb\x79\xd4\xf5\xe7\xc0\xdd\x7c\x62\x2d\x76\x50\x2c\x39\x4b\x9e\x6a\x6b\x5e\xa8\xa0\xf8\x9c\xae\xdb\xab\x56\x1f\x32\xb1\x6d\xaf\x44\x3d\x31\x63\xf1\x61\x81\x3e\x97\x42\x3c\xe5\x34\xe8\xda\xc4\xa4\x27\x89\x3d\x8b\xd6\xe6\x84\x13\x8f\x59\x20\x57\x8b\xd9\x3a\xd5\x84\x7d\x4e\x1a\xc0\xef\x20\x65\xc3\x7d\xf1\x21\xe8\x56\x7c\xf4\xec\x85\x06\x88\x3a\x1d\x32\x62\x82\x9c\x7a\xea\xe0\x38\xcb\xbd\x4f\x8d\xf8\xb4\xe8\x65\x4d\xa8\x2b\xbd\xae\xba\x5c\x56\x99\x86\xbb\x7d\xdd\xad\x4d\x44\x5c\xa0\x90\x0b\xd4\x8b\xdd\x8f\x38\x4d\x3a\x74\x75\xdc\x50\xae\x09\x5d\xc9\xd1\x02\xa6\xec\x93\xc0\x30\x17\x8d\xd2\xfa\x2f\x52\xdc\xd3\xa3\x99\x8f\x04\x1f\xda\xdf\xf6\xf5\x81\xd2\xe0\x34\xc7\x23\x15\xed\x9f\x02\x45\x6d\x79\xbb\x8b\x49\x47\xd4\x6f\x14\xb1\xca\x36\x75\xf5\x5d\xb0\x97\xcf\xde\xb6\x4d\x26\x73\xc1\x6b\xbf\x19\x86\xd6\x81\x9e\x7e\x73\x1a\x6f\x13\xa7\xdd\x9e\x3b\x78\xd4\x69\xab\xea\xab\x23\xd5\x8a\xde\xc1\x77\xae\xd5\x6f\x16\xb8\xdb\x1d\x0b\xe1\xb9\xd9\xa9\x37\xf5\xad\xf9\x27\x00\x00\xff\xff\xe8\x19\x65\x16\xa6\x07\x00\x00"),
		},
		"/nosync/mutex.go": &vfsgen۰CompressedFileInfo{
			name:             "mutex.go",
			modTime:          time.Date(2021, 8, 22, 11, 13, 56, 543823400, time.UTC),
			uncompressedSize: 2073,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x54\xcb\x6e\xdb\x30\x10\x3c\x4b\x5f\xb1\xc9\xc9\x4e\x62\xa5\xbd\xb6\xf5\xa1\x68\x81\x22\x40\x7a\x09\x50\xe4\x4c\x53\x2b\x99\xb0\x44\x1a\x24\x55\xd5\x4d\xf2\xef\xc5\xf2\x21\xcb\x92\xec\xc4\x2d\xaa\x93\xb0\xe4\xce\xce\xec\x0c\xb8\x65\x7c\xc3\x4a\x04\xa9\xcc\x4e\xf2\x34\xbd\xbd\x85\xef\x8d\xc5\x5f\x20\x0c\x30\xc8\x9b\xba\xde\x41\xbb\x16\x7c\x4d\x05\xa9\xe4\x62\x55\x29\xbe\x11\xb2\xcc\x52\xbb\xdb\x62\xb8\x6c\xac\x6e\xb8\x85\xa7\x34\xa1\x53\xcc\x61\xa5\x54\x95\xbe\x38\xb8\x7b\xc5\x37\x40\x65\x03\x75\x06\x77\xd6\x23\xeb\x46\x2e\xac\xa8\x11\x50\x6b\xa5\x41\x14\x50\xbb\x83\x4a\x23\xcb\x77\xe0\x61\xb2\xb4\x68\x24\x87\x59\x0d\x57\x6e\xce\xdc\x81\xcd\xe6\x34\x88\x3a\xb2\x30\xed\x29\x4d\x92\x2d\x93\x82\xcf\x2e\xbd\x8e\x0f\x50\x77\x22\x0e\x10\x2f\xe7\x69\xf2\x92\x26\x5d\xe7\x12\xac\x6e\x30\x30\xfd\x21\xa9\x0a\x8d\x7c\x2b\x5b\xa9\xec\x51\xa6\x1e\xac\xe3\x7a\x71\x8a\xac\x9f\x08\xaa\x08\x7f\x98\x7b\xfe\x63\xb6\x05\xab\x4c\xa4\xfb\xf0\x78\x96\x53\xf1\xfa\xde\xab\x56\x0b\x8b\xf7\x1e\x9a\x3e\x67\x5a\x42\xeb\xa2\xe2\x17\xd5\x48\x8b\x1a\x84\xb4\x13\x4e\x42\xa1\x34\x10\x00\x0d\x38\xb1\x27\xdd\x8e\x4d\x70\xbd\x54\x10\xb2\x84\x1e\x4c\xd8\xa1\x6e\xe1\x2a\x90\x1d\x18\xae\xdb\x6c\xc8\xee\x62\x09\xef\xe0\xf9\x99\x8e\xfa\x72\xce\x4e\xc4\xa0\xff\x54\x2e\x74\x7b\x9e\xf8\x7d\x4a\x0e\xfa\xa6\xd4\x0e\x43\xf3\xba\xaa\x57\xa2\x33\x92\x75\x10\xa0\x91\xa1\xc1\x94\xff\x69\xe8\xc3\xd0\xd1\x7f\xb5\x6d\x90\x88\xeb\xeb\xa8\xae\xb3\x2d\x57\x48\x5a\x8c\x90\x65\x85\x41\x35\x67\x55\xf5\x11\x84\x05\x77\x48\x16\xb1\xa2\x40\x6e\x41\xd9\x35\x6a\x30\xa2\x6e\x2a\xcb\x24\xaa\xc6\x38\x65\xa8\xcd\xd9\x4e\xc7\x6d\x4e\xae\x61\x60\xf5\x44\xb4\x97\x14\xed\xbf\xb2\x7c\x80\xb4\x58\x84\x95\x3c\x32\x61\xbf\x69\xd5\x6c\xdf\xfa\x66\xec\x1b\xf6\xaf\x06\x1f\xbd\x0b\x9f\xf3\x1c\x58\x9e\x1b\xc8\xb1\xb2\xec\x26\x20\xd6\x6c\x07\x2b\x04\x89\x25\xb3\xe2\x27\xde\x80\x55\x60\xd7\x7d\xcc\xbb\xc2\x15\x22\x60\xe9\x9c\xe8\xae\x13\xaa\x53\x6e\xe2\x02\xdb\x12\xae\xba\xee\x39\x5d\x98\xb9\x89\x44\xc5\xed\xb1\x2d\xb3\x08\x76\xbd\xf4\x6c\xdc\x72\x7b\xf5\x4f\x87\x3b\xf5\x1b\x8d\x43\x7b\xdc\xc2\x7d\xbf\x53\x2f\xf3\xab\x92\x08\x39\x72\x8d\x35\x4a\x6b\x06\x62\x42\xc3\x11\xae\xd4\x3b\x8b\x1c\x89\xf8\xe2\xfd\xbc\x67\x4a\x10\x4a\x49\x9a\x44\x8d\xe1\xfa\x8d\x5a\x1d\x99\x40\xbf\x5d\x9a\x7a\x82\x2f\x96\x53\x8a\xc7\x13\x22\x7c\x54\xfc\x27\x00\x00\xff\xff\xec\x95\x29\x83\x19\x08\x00\x00"),
		},
		"/nosync/once.go": &vfsgen۰CompressedFileInfo{
			name:             "once.go",
			modTime:          time.Date(2021, 8, 22, 11, 13, 56, 543823400, time.UTC),
			uncompressedSize: 1072,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x54\x53\xcb\x92\xda\x40\x0c\x3c\xdb\x5f\xd1\xb5\x27\x9c\xa2\xe0\xbe\xa9\x1c\x52\xc5\x65\x4f\x39\xe4\x0b\xc4\x58\x03\xca\x0e\x1a\x32\x0f\x58\x67\x8b\x7f\x4f\x69\x6c\x08\xb9\xd9\x92\xba\xd5\x6a\x69\xce\xe4\xde\xe9\xc0\xd0\x98\x27\x75\x7d\xbf\xdd\xe2\x87\x3a\x86\x64\x90\x22\xee\x7f\xb1\x2b\x28\x47\x2a\xb8\x4a\x08\x38\x73\xf2\x31\x9d\xc0\x1f\xe4\x4a\x98\x10\x95\x41\xae\x48\xd4\x4d\x5f\xa6\x33\xcf\xe0\x5c\x52\x75\x05\x9f\x7d\x37\x46\xd1\x03\xf6\x31\x06\xfb\x56\xc6\xfc\x7d\x6b\x8d\x76\x11\x8e\x42\xc8\x28\x47\x86\xaf\xda\x78\xe0\x21\x1e\xa4\x23\xa2\x86\xc9\xbe\x77\xd1\xd4\xec\xd9\x98\xac\x9e\x47\xf8\x98\x0c\x64\x24\x5e\x52\x2e\x28\x72\xe2\x25\x2a\x19\xa2\xb9\x90\x09\x89\xbe\x09\xda\xe0\x4d\x11\xcb\x91\x13\xae\x31\x8d\x79\x8d\x83\x5c\x58\x0d\xde\x5d\x28\x21\x5a\xad\x15\x5a\x44\x7c\xfb\xdf\xec\xe2\xca\x0f\xd6\x79\xe9\x79\xaa\xa1\xc8\x39\x70\xeb\x95\xd7\xb3\xbc\xa6\xbc\x29\xb0\xaa\xd9\x23\xd1\x4b\x7c\x67\xf8\xb5\xb1\xf1\x85\xd5\x28\x3d\x8e\x94\x41\x18\xc5\x7b\x4e\xac\x05\x17\x0a\x95\x21\x0a\x26\x77\x6c\x20\x47\xcd\x48\xe0\x3b\x94\xaf\xcf\x53\x3c\xaf\x25\xf1\xef\x2a\x69\x31\xa1\x61\x1f\xd6\x95\x08\xfe\x60\x57\x0b\x6f\xfa\xed\x76\xb1\xb8\xf9\x51\x58\xc7\x05\x22\x2a\x45\x28\xc8\x1f\x9a\x31\xb6\xdb\x53\xcd\x05\x7b\x46\xaa\xfa\xb4\x5a\x33\x0e\x3f\xc5\xfa\x36\x05\x92\xa1\x12\x68\x14\xb7\x86\x14\x9c\x68\x32\x8c\xb2\xe3\x9c\x29\x4d\xd6\xbe\x66\x06\xfd\x13\x14\xa4\x70\xa2\x60\x19\x47\xe7\x52\x13\xdf\xd7\x46\xe9\x50\x4f\xac\x25\x5b\x8e\xfe\x1b\x61\xcf\x8b\x85\x23\xf6\x13\x76\xf1\xb5\xed\xc9\x45\xf5\x72\xd8\x3c\x56\x53\xd5\xad\x06\x7c\x62\x89\xdb\x54\x2b\x2f\x81\x95\x4e\x3c\xe0\x36\x2c\x06\xbc\x99\xf5\x8e\x6a\xe6\x6c\x66\xcc\xf4\xf3\x46\xdb\x10\xf3\x55\x93\x8a\xdb\x3c\x23\x5a\x24\xaf\xdb\x89\x46\xcd\x32\x72\xca\x56\x5e\x22\x8e\x74\x61\x24\x2e\x35\x29\x8f\x5f\xe1\x6b\x1b\x6b\x3e\xe4\xd8\xae\x75\x4e\x1a\xd7\x55\xca\x31\xd6\xf9\x38\xec\x7c\x7d\x6b\x62\xda\xb1\x8a\xf8\x62\x2b\x1d\x60\xd3\x60\x9e\x67\xb0\x37\x63\x07\xb8\x69\x8f\xe5\xb3\xef\xba\x85\xac\xbb\x3d\x12\x46\x64\x99\xa6\x71\xf5\x32\xbf\xdc\xd7\xfb\x6b\xe2\xb1\x75\x15\x85\x7f\x19\x1a\xec\x8e\xf9\x86\x92\x2a\xf7\xdd\xc8\x9e\x13\xee\x06\xf6\xdd\x53\x81\xa7\x90\x79\x89\x28\x3f\x10\xb7\xd5\xd0\x77\x7e\x35\xf4\xb7\xfe\x6f\x00\x00\x00\xff\xff\xf9\x72\xbe\xa9\x30\x04\x00\x00"),
		},
		"/nosync/pool.go": &vfsgen۰CompressedFileInfo{
			name:             "pool.go",
			modTime:          time.Date(2021, 8, 22, 11, 13, 56, 543823400, time.UTC),
			uncompressedSize: 2130,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x64\x55\x3f\x93\xdb\xc6\x0f\xad\x4f\x9f\x02\xbf\xea\x77\xca\xe8\x74\x49\xeb\x99\x2b\x32\x29\x1c\x37\x89\x8b\x74\x1e\x17\x10\x09\x8a\x88\x97\x0b\x06\xc0\x4a\xa2\x3d\xf7\xdd\x33\x58\xfe\x39\x39\xee\x44\xee\xf2\xe1\xe1\xbd\x07\x68\xc4\xe6\x0b\x9e\x09\xb2\xd8\x94\x9b\xdd\xee\xf9\x19\x7e\x85\x8f\x22\x09\xd8\x00\xc1\xc8\x41\x3a\x70\x1a\x46\x51\xd4\x09\xe4\xf4\x37\x35\x6e\xe0\x3d\x3a\x0c\x38\xc1\x89\x80\x73\xcb\x17\x6e\x0b\xa6\x34\x81\xe1\x85\x5a\xc0\xdc\x06\x94\x92\x2b\xd3\x85\xda\xe3\xee\xf9\xb9\x62\xe7\x09\xd8\x69\x00\x73\x51\x6a\x81\x33\x78\x4f\x73\xc1\x05\x4d\x69\x90\x0a\x51\x5c\x06\x74\x6e\x2a\x2c\x3a\x60\x9e\xc0\x79\x20\xb8\xb2\xf7\x52\x3c\xf0\xb2\x38\x77\xdc\xa0\xb3\xe4\x23\x7c\xe8\xde\xd0\x7a\x49\xad\xd5\x47\xc9\x69\x02\xa5\x8e\x94\x72\x43\x70\xed\x29\x8a\xb2\x41\x8f\xe3\x48\xd9\x0e\x71\x2b\xc0\x2a\xb1\x81\xcf\xbd\x07\x8f\x96\x30\x25\x69\xd0\xef\xd8\x6f\xca\x18\x76\x04\x9d\x28\x14\x23\x38\x4d\x30\x94\xe4\x3c\x26\x82\xb3\xa8\x14\xe7\x4c\x06\xc6\xf1\x16\x33\x49\xb1\x34\xad\x18\x81\xf0\x7f\x83\xb1\xe8\x28\x46\x81\xe5\x02\x0d\x36\x3d\xc1\x56\x0f\x4e\xc5\xa1\xe4\x62\xa1\x90\xd3\x60\xb5\x54\x42\x27\x05\xa5\x62\x74\x98\xc5\x4d\x4c\x17\xce\x67\x18\x95\xcc\x8a\x46\xab\xb5\xe3\x33\xea\x29\x4c\x6d\x24\x25\x6a\x5c\xf4\x08\x7f\x85\x5f\x6c\x07\xe0\xb0\xed\x0b\x59\xfc\x20\xb4\x09\x5c\x02\xec\x54\x38\xb5\x40\x5d\xc7\x0d\x53\xf6\xd0\x44\x09\xdb\xa7\xb9\x51\x25\x82\xc4\xe6\x76\x84\xdf\xe5\x4a\x17\xd2\x0a\xc4\x16\x06\x80\x15\x76\x3c\xa5\x59\x10\x4c\x29\xf0\xee\x3e\xd9\xac\x07\x1c\x47\x95\x51\x19\x9d\xaa\x70\xd2\x01\x6e\x92\xba\xc0\x80\x39\x68\x23\x9c\x55\xca\xf8\x7d\xf0\xaa\x0e\x81\x63\x9c\x28\x7b\x24\xad\xc7\x88\x10\x0e\x92\xcf\x11\x38\x18\xc5\x29\x3b\xd7\xbc\x54\x99\xda\xb0\xa6\x91\xdc\x14\x55\xca\x1e\x41\xa5\x91\x72\x4b\xb9\x86\xa7\x49\xd1\xaa\xcd\x34\x96\x41\x38\xce\x7c\x46\x95\x0b\xb7\x14\x23\x70\xc5\xd0\x28\xca\xa8\xf3\xd7\xcd\x25\x96\x0c\x72\x21\xed\x09\x6b\xd4\xb1\x51\x31\x8b\x16\xa6\x15\xf8\xae\x73\xba\xe1\x10\xf1\x90\x0e\xce\x22\xed\x8f\xdd\x2f\x83\xd0\x0d\xbe\x32\x39\xc0\xb5\xe7\xa6\x87\x01\x39\x3b\x72\x36\xc0\x00\x6b\xa7\x8c\xc3\x3c\x14\x4f\xc6\x5f\xa9\x9d\x47\xe9\x3f\x53\x5a\x7c\x2c\x0e\xa7\xd2\x75\xa4\x16\xee\xd3\x72\xcd\x1a\x4c\x64\x50\x72\x4b\x1a\x70\x49\xb0\x85\xc7\x3a\x13\x95\xfa\x5d\x7e\x51\x09\xb0\x71\xbe\x50\x9a\x60\x54\xce\xce\xf9\xbc\xaf\x4a\x5b\xaf\x9c\xbf\x58\x9d\xa5\x40\xf9\xa7\x30\x59\x43\xd9\xd7\x96\xff\x9c\xdb\x11\xef\x49\xa1\xc7\xdc\x1e\x00\xdf\x32\xb1\xf5\x14\xf6\x19\x8c\xa8\x3e\xab\x61\xbd\xa8\x3f\x25\x8e\xf9\x9f\x37\x0d\xb0\x2d\x73\x1e\xc7\x6b\xd0\x42\xbe\x1a\xb6\xaa\xdf\x01\x8c\x63\xb2\x6b\xc5\xc5\x12\x68\x85\xe6\x74\x6e\xc6\x5d\x29\x25\xe0\xca\xb7\x6e\xaf\x20\x8c\xca\x72\x84\x0f\x35\xca\x43\xe8\xb3\x4d\x40\x78\xde\xe3\x85\xc0\x4a\xd3\x6f\x6b\x8f\xc3\xc5\xa1\x1e\xf7\xc4\x0a\x72\xcd\xdf\xa5\xbd\xf6\xef\xd3\xb8\x2c\x21\x73\x2d\x8d\xc3\xb7\xdd\xc3\xac\xfe\xa7\xcf\x9c\x9d\xb4\xc3\x86\xbe\xbd\xee\x1e\xfe\xa0\x2b\x00\x74\x25\x37\x8f\x7b\xb8\x3f\x79\xad\x8b\xf8\x3d\x39\x18\xa5\x5a\x18\x33\xa0\x9e\xd8\xb7\x59\x80\x4e\x65\xd8\xd6\xdd\x61\x59\x9b\x75\xac\xd7\x93\x75\xdd\x1c\xaa\x67\x4a\x5e\x34\xd7\x0b\x2e\xf5\xc3\x08\x11\xe9\x71\x2d\x15\xfb\xb7\xe9\x25\xb6\x92\x0b\xf0\x39\x07\xe3\xb8\x37\x46\x2b\x01\xe1\x4a\xb1\x45\x3c\x4c\xa3\x61\xf4\xba\xd4\xe0\xb7\x0a\x63\x61\x5e\x49\xed\xac\xb9\x59\x19\xa8\x6e\x6c\xa5\x34\x0f\xcb\x89\xfc\x4a\x94\xe1\x82\xa9\x50\x98\x6e\x31\xa0\x2e\xf0\xb1\xf8\xfa\x7f\x11\xd5\x96\xf3\x99\xee\x3c\xc2\xef\x69\x0b\xd6\x87\xae\x72\xbd\xd6\x52\x35\x5e\x57\x36\x5a\x6e\x43\xe6\x99\xe8\x78\x0c\x69\xeb\x7a\xca\x4f\x99\xd3\xa1\x7e\xb4\x28\xb0\x16\x52\xb2\x92\x6a\xf0\x42\x88\xba\x47\xe3\xb3\xe3\x2e\x0c\x81\xc7\x11\x7e\x0a\xf1\xf6\xf1\xe9\xf7\xf6\x84\x9f\xdc\x41\xa2\xfc\x38\x1e\xab\xb1\x7b\x78\x79\x81\x9f\xe3\x7d\x1c\xcc\xd5\xff\xf7\x52\xe9\xc4\xbb\x87\x85\x5e\x3d\x78\xdc\xef\x1e\x1e\x5e\x77\xdb\xcb\xcc\x69\x17\xcf\x37\x78\xf7\x02\x0b\xde\xa7\x7b\xec\xa7\x5f\x3e\xef\x1e\x96\x07\x78\xbb\xf2\xee\x87\x3b\x0b\xe0\x6d\x89\x4f\xd5\xb5\x6d\x0d\x6e\xab\xe1\x61\xe4\x0f\xed\x7d\x2c\xfe\x78\xbb\x6f\x6f\xbf\xf4\x77\x8b\xa6\xd6\x16\x66\xec\x4a\xf4\x8d\x4a\xfd\xff\x6c\x57\x12\x07\xb8\xed\x77\xaf\xbb\x7f\x03\x00\x00\xff\xff\x07\xba\x3e\x57\x52\x08\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/js"].(os.FileInfo),
		fs["/nosync"].(os.FileInfo),
	}
	fs["/js"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/js/js.go"].(os.FileInfo),
		fs["/js/js_test.go"].(os.FileInfo),
	}
	fs["/nosync"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/nosync/map.go"].(os.FileInfo),
		fs["/nosync/mutex.go"].(os.FileInfo),
		fs["/nosync/once.go"].(os.FileInfo),
		fs["/nosync/pool.go"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
