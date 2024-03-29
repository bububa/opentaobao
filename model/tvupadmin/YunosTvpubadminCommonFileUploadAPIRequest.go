package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminCommonFileUploadAPIRequest 文件上传API API请求
// yunos.tvpubadmin.common.file.upload
//
// 文件上传服务
type YunosTvpubadminCommonFileUploadAPIRequest struct {
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

// NewYunosTvpubadminCommonFileUploadRequest 初始化YunosTvpubadminCommonFileUploadAPIRequest对象
func NewYunosTvpubadminCommonFileUploadRequest() *YunosTvpubadminCommonFileUploadAPIRequest {
	return &YunosTvpubadminCommonFileUploadAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminCommonFileUploadAPIRequest) Reset() {
	r._originalFilename = ""
	r._size = ""
	r._contentType = ""
	r._uploadPath = ""
	r._bytes = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminCommonFileUploadAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.common.file.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminCommonFileUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminCommonFileUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOriginalFilename is OriginalFilename Setter
// 原文件名
func (r *YunosTvpubadminCommonFileUploadAPIRequest) SetOriginalFilename(_originalFilename string) error {
	r._originalFilename = _originalFilename
	r.Set("original_filename", _originalFilename)
	return nil
}

// GetOriginalFilename OriginalFilename Getter
func (r YunosTvpubadminCommonFileUploadAPIRequest) GetOriginalFilename() string {
	return r._originalFilename
}

// SetSize is Size Setter
// 文件大小
func (r *YunosTvpubadminCommonFileUploadAPIRequest) SetSize(_size string) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r YunosTvpubadminCommonFileUploadAPIRequest) GetSize() string {
	return r._size
}

// SetContentType is ContentType Setter
// 文件类型
func (r *YunosTvpubadminCommonFileUploadAPIRequest) SetContentType(_contentType string) error {
	r._contentType = _contentType
	r.Set("content_type", _contentType)
	return nil
}

// GetContentType ContentType Getter
func (r YunosTvpubadminCommonFileUploadAPIRequest) GetContentType() string {
	return r._contentType
}

// SetUploadPath is UploadPath Setter
// 上传地址
func (r *YunosTvpubadminCommonFileUploadAPIRequest) SetUploadPath(_uploadPath string) error {
	r._uploadPath = _uploadPath
	r.Set("upload_path", _uploadPath)
	return nil
}

// GetUploadPath UploadPath Getter
func (r YunosTvpubadminCommonFileUploadAPIRequest) GetUploadPath() string {
	return r._uploadPath
}

// SetBytes is Bytes Setter
// 文件字节流
func (r *YunosTvpubadminCommonFileUploadAPIRequest) SetBytes(_bytes *model.File) error {
	r._bytes = _bytes
	r.Set("bytes", _bytes)
	return nil
}

// GetBytes Bytes Getter
func (r YunosTvpubadminCommonFileUploadAPIRequest) GetBytes() *model.File {
	return r._bytes
}

var poolYunosTvpubadminCommonFileUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminCommonFileUploadRequest()
	},
}

// GetYunosTvpubadminCommonFileUploadRequest 从 sync.Pool 获取 YunosTvpubadminCommonFileUploadAPIRequest
func GetYunosTvpubadminCommonFileUploadAPIRequest() *YunosTvpubadminCommonFileUploadAPIRequest {
	return poolYunosTvpubadminCommonFileUploadAPIRequest.Get().(*YunosTvpubadminCommonFileUploadAPIRequest)
}

// ReleaseYunosTvpubadminCommonFileUploadAPIRequest 将 YunosTvpubadminCommonFileUploadAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminCommonFileUploadAPIRequest(v *YunosTvpubadminCommonFileUploadAPIRequest) {
	v.Reset()
	poolYunosTvpubadminCommonFileUploadAPIRequest.Put(v)
}
