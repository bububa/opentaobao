package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuPhotobankUploadAPIRequest 图片银行图片上传开放接口 API请求
// alibaba.icbu.photobank.upload
//
// 图片银行图片上传开放接口
type AlibabaIcbuPhotobankUploadAPIRequest struct {
	model.Params
	// 扩展参数信息,如ICVID
	_extraContext string
	// 上传图片名称
	_fileName string
	// 上传图片所在分组
	_groupId string
	// 图片字节数组
	_imageBytes *model.File
}

// NewAlibabaIcbuPhotobankUploadRequest 初始化AlibabaIcbuPhotobankUploadAPIRequest对象
func NewAlibabaIcbuPhotobankUploadRequest() *AlibabaIcbuPhotobankUploadAPIRequest {
	return &AlibabaIcbuPhotobankUploadAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuPhotobankUploadAPIRequest) Reset() {
	r._extraContext = ""
	r._fileName = ""
	r._groupId = ""
	r._imageBytes = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuPhotobankUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.photobank.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuPhotobankUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuPhotobankUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtraContext is ExtraContext Setter
// 扩展参数信息,如ICVID
func (r *AlibabaIcbuPhotobankUploadAPIRequest) SetExtraContext(_extraContext string) error {
	r._extraContext = _extraContext
	r.Set("extra_context", _extraContext)
	return nil
}

// GetExtraContext ExtraContext Getter
func (r AlibabaIcbuPhotobankUploadAPIRequest) GetExtraContext() string {
	return r._extraContext
}

// SetFileName is FileName Setter
// 上传图片名称
func (r *AlibabaIcbuPhotobankUploadAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r AlibabaIcbuPhotobankUploadAPIRequest) GetFileName() string {
	return r._fileName
}

// SetGroupId is GroupId Setter
// 上传图片所在分组
func (r *AlibabaIcbuPhotobankUploadAPIRequest) SetGroupId(_groupId string) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaIcbuPhotobankUploadAPIRequest) GetGroupId() string {
	return r._groupId
}

// SetImageBytes is ImageBytes Setter
// 图片字节数组
func (r *AlibabaIcbuPhotobankUploadAPIRequest) SetImageBytes(_imageBytes *model.File) error {
	r._imageBytes = _imageBytes
	r.Set("image_bytes", _imageBytes)
	return nil
}

// GetImageBytes ImageBytes Getter
func (r AlibabaIcbuPhotobankUploadAPIRequest) GetImageBytes() *model.File {
	return r._imageBytes
}

var poolAlibabaIcbuPhotobankUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuPhotobankUploadRequest()
	},
}

// GetAlibabaIcbuPhotobankUploadRequest 从 sync.Pool 获取 AlibabaIcbuPhotobankUploadAPIRequest
func GetAlibabaIcbuPhotobankUploadAPIRequest() *AlibabaIcbuPhotobankUploadAPIRequest {
	return poolAlibabaIcbuPhotobankUploadAPIRequest.Get().(*AlibabaIcbuPhotobankUploadAPIRequest)
}

// ReleaseAlibabaIcbuPhotobankUploadAPIRequest 将 AlibabaIcbuPhotobankUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuPhotobankUploadAPIRequest(v *AlibabaIcbuPhotobankUploadAPIRequest) {
	v.Reset()
	poolAlibabaIcbuPhotobankUploadAPIRequest.Put(v)
}
