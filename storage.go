package main

import (
	"time"
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
	Get(name string) error
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
func (s *s3Service) Get(name string) error {
	return nil
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
func (f *fsService) Get(name string) error {
	return nil
}

// @TODO
func (f *fsService) Put(name, uri string) error {
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
