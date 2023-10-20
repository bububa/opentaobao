package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuphotobankuploadAPIRequest 图片银行图片上传开放接口 API请求
// alibaba.icbu.photobank.upload
//
// 图片银行图片上传开放接口
type AlibabaicbuphotobankuploadAPIRequest struct {
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

// NewAlibabaicbuphotobankuploadRequest 初始化AlibabaicbuphotobankuploadAPIRequest对象
func NewAlibabaicbuphotobankuploadRequest() *AlibabaicbuphotobankuploadAPIRequest {
	return &AlibabaicbuphotobankuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuphotobankuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.photobank.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuphotobankuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuphotobankuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtraContext is ExtraContext Setter
// 扩展参数信息,如ICVID
func (r *AlibabaicbuphotobankuploadAPIRequest) SetExtraContext(_extraContext string) error {
	r._extraContext = _extraContext
	r.Set("extra_context", _extraContext)
	return nil
}

// GetExtraContext ExtraContext Getter
func (r AlibabaicbuphotobankuploadAPIRequest) GetExtraContext() string {
	return r._extraContext
}

// SetFileName is FileName Setter
// 上传图片名称
func (r *AlibabaicbuphotobankuploadAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r AlibabaicbuphotobankuploadAPIRequest) GetFileName() string {
	return r._fileName
}

// SetGroupId is GroupId Setter
// 上传图片所在分组
func (r *AlibabaicbuphotobankuploadAPIRequest) SetGroupId(_groupId string) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaicbuphotobankuploadAPIRequest) GetGroupId() string {
	return r._groupId
}

// SetImageBytes is ImageBytes Setter
// 图片字节数组
func (r *AlibabaicbuphotobankuploadAPIRequest) SetImageBytes(_imageBytes *model.File) error {
	r._imageBytes = _imageBytes
	r.Set("image_bytes", _imageBytes)
	return nil
}

// GetImageBytes ImageBytes Getter
func (r AlibabaicbuphotobankuploadAPIRequest) GetImageBytes() *model.File {
	return r._imageBytes
}
