// Package s3manageriface provides an interface for the s3manager package
package s3manageriface

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// DownloaderAPI is the interface type for s3manager.Downloader.
type DownloaderAPI interface {
	Download(io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) (int64, error)
	DownloadWithContext(aws.Context, io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) (int64, error)
}

var _ DownloaderAPI = (*s3manager.Downloader)(nil)

type DownloadWithIterator interface {
	DownloadWithIterator(aws.Context, s3manager.BatchDownloadIterator, ...func(*s3manager.Downloader)) error
}

type NewDownloader interface {
	NewDownloader(client.ConfigProvider, ...func(*s3manager.Downloader)) *s3manager.Downloader
}

type NewDownloaderWithClient interface {
	NewDownloaderWithClient(s3iface.S3API, ...func(*s3manager.Downloader)) *s3manager.Downloader
}


// UploaderAPI is the interface type for s3manager.Uploader.
type UploaderAPI interface {
	Upload(*s3manager.UploadInput, ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
	UploadWithContext(aws.Context, *s3manager.UploadInput, ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
}

var _ UploaderAPI = (*s3manager.Uploader)(nil)

type UploadWithIterator interface {
	UploadWithIterator(aws.Context, s3manager.BatchUploadIterator, ...func(*s3manager.Uploader)) error
}

type NewUploader interface {
	NewUploader(client.ConfigProvider, ...func(*s3manager.Uploader)) *s3manager.Uploader
}

type NewUploaderWithClient interface {
	NewUploaderWithClient(s3iface.S3API, ...func(*s3manager.Uploader)) *s3manager.Uploader
}


type BatchDeleteAPI interface {
	Delete(aws.Context, s3manager.BatchDeleteIterator) error
}