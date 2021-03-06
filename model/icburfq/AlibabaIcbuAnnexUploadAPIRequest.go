package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuAnnexUploadAPIRequest 上传附件获取附件files_str API请求
// alibaba.icbu.annex.upload
//
// 上传附件获取附件files_str
type AlibabaIcbuAnnexUploadAPIRequest struct {
	model.Params
	// 文件名
	_fileName string
	// 来源
	_source string
	// 文件字节流
	_fileInputStreamBytes *model.File
}

// NewAlibabaIcbuAnnexUploadRequest 初始化AlibabaIcbuAnnexUploadAPIRequest对象
func NewAlibabaIcbuAnnexUploadRequest() *AlibabaIcbuAnnexUploadAPIRequest {
	return &AlibabaIcbuAnnexUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuAnnexUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.annex.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuAnnexUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFileName is FileName Setter
// 文件名
func (r *AlibabaIcbuAnnexUploadAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r AlibabaIcbuAnnexUploadAPIRequest) GetFileName() string {
	return r._fileName
}

// SetSource is Source Setter
// 来源
func (r *AlibabaIcbuAnnexUploadAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r AlibabaIcbuAnnexUploadAPIRequest) GetSource() string {
	return r._source
}

// SetFileInputStreamBytes is FileInputStreamBytes Setter
// 文件字节流
func (r *AlibabaIcbuAnnexUploadAPIRequest) SetFileInputStreamBytes(_fileInputStreamBytes *model.File) error {
	r._fileInputStreamBytes = _fileInputStreamBytes
	r.Set("file_input_stream_bytes", _fileInputStreamBytes)
	return nil
}

// GetFileInputStreamBytes FileInputStreamBytes Getter
func (r AlibabaIcbuAnnexUploadAPIRequest) GetFileInputStreamBytes() *model.File {
	return r._fileInputStreamBytes
}
