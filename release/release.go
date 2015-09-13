package release

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/jolshevski/chester/metadata"
	"io/ioutil"
	"path/filepath"
)

// Release represents a specific Puppet
// module release on disk.
type Release struct {
	LocalPath string
	Metadata  metadata.Metadata
	File_uri  string
	File_md5  string
}

// New instantiates a new release object
// given the module's name, version and
// the path containing module tarballs.
func New(q string, v string, modulepath string) *Release {
	path, _ := filepath.Glob(modulepath + "/" + q + "-" + v + ".tar.gz")

	return &Release{LocalPath: path[0]}
}

// Tarball returns the path to the
// release's tarball on disk.
func (r *Release) Tarball() string {
	return r.LocalPath
}

// FromDisk populates the applicable
// properties given the tarball on
// disk.
func (r *Release) FromDisk() (err error) {
	// Get the metadata object from metadata.json
	m, err := metadata.FromRelease(r)
	r.Metadata = *m

	// Get the tarball's MD5 checksum
	raw, _ := ioutil.ReadFile(r.LocalPath)
	checker := md5.New()
	checker.Write(raw)
	r.File_md5 = hex.EncodeToString(checker.Sum(nil))

	return
}
