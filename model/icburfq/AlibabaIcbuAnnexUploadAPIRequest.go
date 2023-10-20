package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuannexuploadAPIRequest 上传附件获取附件files_str API请求
// alibaba.icbu.annex.upload
//
// 上传附件获取附件files_str
type AlibabaicbuannexuploadAPIRequest struct {
	model.Params
	// 文件名
	_fileName string
	// 来源
	_source string
	// 文件字节流
	_fileInputStreamBytes *model.File
}

// NewAlibabaicbuannexuploadRequest 初始化AlibabaicbuannexuploadAPIRequest对象
func NewAlibabaicbuannexuploadRequest() *AlibabaicbuannexuploadAPIRequest {
	return &AlibabaicbuannexuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuannexuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.annex.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuannexuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuannexuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileName is FileName Setter
// 文件名
func (r *AlibabaicbuannexuploadAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r AlibabaicbuannexuploadAPIRequest) GetFileName() string {
	return r._fileName
}

// SetSource is Source Setter
// 来源
func (r *AlibabaicbuannexuploadAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabaicbuannexuploadAPIRequest) GetSource() string {
	return r._source
}

// SetFileInputStreamBytes is FileInputStreamBytes Setter
// 文件字节流
func (r *AlibabaicbuannexuploadAPIRequest) SetFileInputStreamBytes(_fileInputStreamBytes *model.File) error {
	r._fileInputStreamBytes = _fileInputStreamBytes
	r.Set("file_input_stream_bytes", _fileInputStreamBytes)
	return nil
}

// GetFileInputStreamBytes FileInputStreamBytes Getter
func (r AlibabaicbuannexuploadAPIRequest) GetFileInputStreamBytes() *model.File {
	return r._fileInputStreamBytes
}
