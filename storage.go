package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	s3URI := conf.Storage.S3StorageBucketURI

	if s3URI != "" {
		storage = newS3Service(s3URI)
	} else {
		storage = newFSService()
	}
}

var storage storageService

type storageService interface {
	Get(name string) (*os.File, error)
	Put(name, uri string) error
	Exists(name string) (bool, error)
	List() ([]string, error)
}

type storageStats struct {
	// total size in bytes
	TotalSize     uint
	MemeCount     uint
	InitializedOn time.Time
}

type storageFileStats struct {
	// size in bytes
	Size       uint
	Dimensions string
	UploadedOn time.Time
}

type s3Service struct {
	URI string
}

type fsService struct {
}

// @TODO
func (s *s3Service) Get(name string) (*os.File, error) {
	return nil, nil
}

// @TODO
func (s *s3Service) Put(name, uri string) error {
	return nil
}

// @TODO
func (s *s3Service) Exists(name string) (bool, error) {
	return false, nil
}

// @TODO
func (s *s3Service) List() ([]string, error) {
	return []string{}, nil
}

// @TODO
func (f *fsService) Get(name string) (*os.File, error) {
	fPath := "/usr/share/memes/" + name + ".jpg"
	fHandle, err := os.Open(fPath)
	if err != nil {
		log.WithFields(logrus.Fields{
			"path":         fPath,
			"errorMessage": err.Error(),
		}).Error("could not load file")
		return nil, err
	}

	log.WithFields(logrus.Fields{
		"path": fPath,
	}).Error("successfully got image")
	return fHandle, nil
}

// @TODO
func (f *fsService) Put(name, uri string) error {
	res, err := http.Get(uri)
	if err != nil {
		log.WithFields(logrus.Fields{
			"uri":          uri,
			"errorMessage": err.Error(),
		}).Errorf("could not get image at uri `%s`", uri)
		return err
	}

	defer res.Body.Close()

	fPath := "/usr/share/memes/" + name + ".jpg"
	fHandle, err := os.Create(fPath)
	if err != nil {
		log.WithFields(logrus.Fields{
			"dest":         fPath,
			"errorMessage": err.Error(),
		}).Errorf("could not create file")
		return err
	}

	defer fHandle.Close()

	_, err = io.Copy(fHandle, res.Body)
	if err != nil {
		return err
	}

	log.WithFields(logrus.Fields{
		"uri":  uri,
		"dest": fPath,
	}).Info("successfully saved image")

	return nil
}

// @TODO
func (f *fsService) Exists(name string) (bool, error) {
	return false, nil
}

// @TODO
func (f *fsService) List() ([]string, error) {
	return []string{}, nil
}

func newS3Service(uri string) *s3Service {
	return &s3Service{uri}
}

func newFSService() *fsService {
	return &fsService{}
}
