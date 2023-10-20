package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalsuitfileuploadAPIRequest 诉讼文件上传接口 API请求
// alibaba.legal.suit.file.upload
//
// 上传文件接口
type AlibabalegalsuitfileuploadAPIRequest struct {
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

// NewAlibabalegalsuitfileuploadRequest 初始化AlibabalegalsuitfileuploadAPIRequest对象
func NewAlibabalegalsuitfileuploadRequest() *AlibabalegalsuitfileuploadAPIRequest {
	return &AlibabalegalsuitfileuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalsuitfileuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.file.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalsuitfileuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalsuitfileuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileName is FileName Setter
// 文件名称
func (r *AlibabalegalsuitfileuploadAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r AlibabalegalsuitfileuploadAPIRequest) GetFileName() string {
	return r._fileName
}

// SetSignature is Signature Setter
// 签名
func (r *AlibabalegalsuitfileuploadAPIRequest) SetSignature(_signature string) error {
	r._signature = _signature
	r.Set("signature", _signature)
	return nil
}

// GetSignature Signature Getter
func (r AlibabalegalsuitfileuploadAPIRequest) GetSignature() string {
	return r._signature
}

// SetFile is File Setter
// 文件
func (r *AlibabalegalsuitfileuploadAPIRequest) SetFile(_file *model.File) error {
	r._file = _file
	r.Set("file", _file)
	return nil
}

// GetFile File Getter
func (r AlibabalegalsuitfileuploadAPIRequest) GetFile() *model.File {
	return r._file
}

// SetTimeStamp is TimeStamp Setter
// 时间搓
func (r *AlibabalegalsuitfileuploadAPIRequest) SetTimeStamp(_timeStamp int64) error {
	r._timeStamp = _timeStamp
	r.Set("time_stamp", _timeStamp)
	return nil
}

// GetTimeStamp TimeStamp Getter
func (r AlibabalegalsuitfileuploadAPIRequest) GetTimeStamp() int64 {
	return r._timeStamp
}

// SetFileSize is FileSize Setter
// 文件大小
func (r *AlibabalegalsuitfileuploadAPIRequest) SetFileSize(_fileSize int64) error {
	r._fileSize = _fileSize
	r.Set("file_size", _fileSize)
	return nil
}

// GetFileSize FileSize Getter
func (r AlibabalegalsuitfileuploadAPIRequest) GetFileSize() int64 {
	return r._fileSize
}
