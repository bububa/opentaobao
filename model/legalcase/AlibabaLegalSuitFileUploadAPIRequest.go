package legalcase

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitFileUploadAPIRequest 诉讼文件上传接口 API请求
// alibaba.legal.suit.file.upload
//
// 上传文件接口
type AlibabaLegalSuitFileUploadAPIRequest struct {
	model.Params
	// 文件名称
	_fileName string
	// 签名
	_signature string
	// 文件
	_file *model.File
	// 时间搓
	_timeStamp int64
	// 文件大小
	_fileSize int64
}

// NewAlibabaLegalSuitFileUploadRequest 初始化AlibabaLegalSuitFileUploadAPIRequest对象
func NewAlibabaLegalSuitFileUploadRequest() *AlibabaLegalSuitFileUploadAPIRequest {
	return &AlibabaLegalSuitFileUploadAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalSuitFileUploadAPIRequest) Reset() {
	r._fileName = ""
	r._signature = ""
	r._file = nil
	r._timeStamp = 0
	r._fileSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitFileUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.file.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitFileUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalSuitFileUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileName is FileName Setter
// 文件名称
func (r *AlibabaLegalSuitFileUploadAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r AlibabaLegalSuitFileUploadAPIRequest) GetFileName() string {
	return r._fileName
}

// SetSignature is Signature Setter
// 签名
func (r *AlibabaLegalSuitFileUploadAPIRequest) SetSignature(_signature string) error {
	r._signature = _signature
	r.Set("signature", _signature)
	return nil
}

// GetSignature Signature Getter
func (r AlibabaLegalSuitFileUploadAPIRequest) GetSignature() string {
	return r._signature
}

// SetFile is File Setter
// 文件
func (r *AlibabaLegalSuitFileUploadAPIRequest) SetFile(_file *model.File) error {
	r._file = _file
	r.Set("file", _file)
	return nil
}

// GetFile File Getter
func (r AlibabaLegalSuitFileUploadAPIRequest) GetFile() *model.File {
	return r._file
}

// SetTimeStamp is TimeStamp Setter
// 时间搓
func (r *AlibabaLegalSuitFileUploadAPIRequest) SetTimeStamp(_timeStamp int64) error {
	r._timeStamp = _timeStamp
	r.Set("time_stamp", _timeStamp)
	return nil
}

// GetTimeStamp TimeStamp Getter
func (r AlibabaLegalSuitFileUploadAPIRequest) GetTimeStamp() int64 {
	return r._timeStamp
}

// SetFileSize is FileSize Setter
// 文件大小
func (r *AlibabaLegalSuitFileUploadAPIRequest) SetFileSize(_fileSize int64) error {
	r._fileSize = _fileSize
	r.Set("file_size", _fileSize)
	return nil
}

// GetFileSize FileSize Getter
func (r AlibabaLegalSuitFileUploadAPIRequest) GetFileSize() int64 {
	return r._fileSize
}

var poolAlibabaLegalSuitFileUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalSuitFileUploadRequest()
	},
}

// GetAlibabaLegalSuitFileUploadRequest 从 sync.Pool 获取 AlibabaLegalSuitFileUploadAPIRequest
func GetAlibabaLegalSuitFileUploadAPIRequest() *AlibabaLegalSuitFileUploadAPIRequest {
	return poolAlibabaLegalSuitFileUploadAPIRequest.Get().(*AlibabaLegalSuitFileUploadAPIRequest)
}

// ReleaseAlibabaLegalSuitFileUploadAPIRequest 将 AlibabaLegalSuitFileUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalSuitFileUploadAPIRequest(v *AlibabaLegalSuitFileUploadAPIRequest) {
	v.Reset()
	poolAlibabaLegalSuitFileUploadAPIRequest.Put(v)
}
