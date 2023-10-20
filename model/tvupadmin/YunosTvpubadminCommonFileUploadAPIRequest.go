package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincommonfileuploadAPIRequest 文件上传API API请求
// yunos.tvpubadmin.common.file.upload
//
// 文件上传服务
type YunostvpubadmincommonfileuploadAPIRequest struct {
	model.Params
	// 原文件名
	_originalFilename string
	// 文件大小
	_size string
	// 文件类型
	_contentType string
	// 上传地址
	_uploadPath string
	// 文件字节流
	_bytes *model.File
}

// NewYunostvpubadmincommonfileuploadRequest 初始化YunostvpubadmincommonfileuploadAPIRequest对象
func NewYunostvpubadmincommonfileuploadRequest() *YunostvpubadmincommonfileuploadAPIRequest {
	return &YunostvpubadmincommonfileuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincommonfileuploadAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.common.file.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincommonfileuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincommonfileuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOriginalFilename is OriginalFilename Setter
// 原文件名
func (r *YunostvpubadmincommonfileuploadAPIRequest) SetOriginalFilename(_originalFilename string) error {
	r._originalFilename = _originalFilename
	r.Set("original_filename", _originalFilename)
	return nil
}

// GetOriginalFilename OriginalFilename Getter
func (r YunostvpubadmincommonfileuploadAPIRequest) GetOriginalFilename() string {
	return r._originalFilename
}

// SetSize is Size Setter
// 文件大小
func (r *YunostvpubadmincommonfileuploadAPIRequest) SetSize(_size string) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r YunostvpubadmincommonfileuploadAPIRequest) GetSize() string {
	return r._size
}

// SetContentType is ContentType Setter
// 文件类型
func (r *YunostvpubadmincommonfileuploadAPIRequest) SetContentType(_contentType string) error {
	r._contentType = _contentType
	r.Set("content_type", _contentType)
	return nil
}

// GetContentType ContentType Getter
func (r YunostvpubadmincommonfileuploadAPIRequest) GetContentType() string {
	return r._contentType
}

// SetUploadPath is UploadPath Setter
// 上传地址
func (r *YunostvpubadmincommonfileuploadAPIRequest) SetUploadPath(_uploadPath string) error {
	r._uploadPath = _uploadPath
	r.Set("upload_path", _uploadPath)
	return nil
}

// GetUploadPath UploadPath Getter
func (r YunostvpubadmincommonfileuploadAPIRequest) GetUploadPath() string {
	return r._uploadPath
}

// SetBytes is Bytes Setter
// 文件字节流
func (r *YunostvpubadmincommonfileuploadAPIRequest) SetBytes(_bytes *model.File) error {
	r._bytes = _bytes
	r.Set("bytes", _bytes)
	return nil
}

// GetBytes Bytes Getter
func (r YunostvpubadmincommonfileuploadAPIRequest) GetBytes() *model.File {
	return r._bytes
}
