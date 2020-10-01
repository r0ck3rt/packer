package driver

import (
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type DatastoreMock struct {
	FileExistsCalled    bool
	MakeDirectoryCalled bool

	ResolvePathCalled bool
	ResolvePathReturn string

	DeleteCalled bool
	DeletePath   string
	DeleteErr    error

	UploadFileCalled  bool
	UploadFileSrc     string
	UploadFileDst     string
	UploadFileHost    string
	UploadFileSetHost bool
	UploadFileErr     error
}

func (ds *DatastoreMock) Info(params ...string) (*mo.Datastore, error) {
	return nil, nil
}

func (ds *DatastoreMock) FileExists(path string) bool {
	ds.FileExistsCalled = true
	return false
}

func (ds *DatastoreMock) Name() string {
	return "datastore-mock"
}

func (ds *DatastoreMock) Reference() types.ManagedObjectReference {
	return types.ManagedObjectReference{}
}

func (ds *DatastoreMock) ResolvePath(path string) string {
	ds.ResolvePathCalled = true
	return ds.ResolvePathReturn
}

func (ds *DatastoreMock) UploadFile(src, dst, host string, setHost bool) error {
	ds.UploadFileCalled = true
	ds.UploadFileSrc = src
	ds.UploadFileDst = dst
	ds.UploadFileHost = host
	ds.UploadFileSetHost = setHost
	return ds.UploadFileErr
}

func (ds *DatastoreMock) Delete(path string) error {
	ds.DeleteCalled = true
	ds.DeletePath = path
	return ds.DeleteErr
}

func (ds *DatastoreMock) MakeDirectory(path string) error {
	ds.MakeDirectoryCalled = true
	return nil
}
