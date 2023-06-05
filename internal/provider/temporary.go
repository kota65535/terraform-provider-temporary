package provider

import (
	"errors"
	"fmt"
	"os"
	"path"
)

type Temporary struct {
	BaseDir string
}

func NewTemporary(baseDir string) *Temporary {
	return &Temporary{
		BaseDir: baseDir,
	}
}

const MarkerFileName = ".terraform-provider-temporary"

func (r *Temporary) Create() error {
	exists, err := r.check()
	if err != nil {
		return err
	}
	if exists {
		err = r.deleteDir()
		if err != nil {
			return err
		}
	}
	return r.createDir()
}

func (r *Temporary) check() (bool, error) {
	s, err := os.Stat(r.BaseDir)
	if err != nil {
		return false, nil
	}
	if s.IsDir() {
		return true, nil
	}
	return false, errors.New("file exists")
}

func (r *Temporary) createDir() error {
	err := os.MkdirAll(r.BaseDir, 0777)
	if err != nil {
		return err
	}
	_, err = os.Create(path.Join(r.BaseDir, MarkerFileName))
	if err != nil {
		return err
	}
	return nil
}

func (r *Temporary) deleteDir() error {
	_, err := os.Stat(path.Join(r.BaseDir, MarkerFileName))
	if err != nil {
		return errors.New(fmt.Sprintf("Marker file %s not found", MarkerFileName))
	}
	return os.RemoveAll(r.BaseDir)
}
