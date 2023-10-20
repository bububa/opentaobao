package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaibshenjingvisitorpaduploadfaceAPIRequest 访客PAD上传人脸 API请求
// alibaba.ib.shenjing.visitor.pad.uploadface
//
// 访客PAD端上传人脸。
type AlibabaibshenjingvisitorpaduploadfaceAPIRequest struct {
	model.Params
	// 访客ID
	_id string
	// 图片URL
	_image string
}

// NewAlibabaibshenjingvisitorpaduploadfaceRequest 初始化AlibabaibshenjingvisitorpaduploadfaceAPIRequest对象
func NewAlibabaibshenjingvisitorpaduploadfaceRequest() *AlibabaibshenjingvisitorpaduploadfaceAPIRequest {
	return &AlibabaibshenjingvisitorpaduploadfaceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaibshenjingvisitorpaduploadfaceAPIRequest) GetApiMethodName() string {
	return "alibaba.ib.shenjing.visitor.pad.uploadface"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaibshenjingvisitorpaduploadfaceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaibshenjingvisitorpaduploadfaceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 访客ID
func (r *AlibabaibshenjingvisitorpaduploadfaceAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaibshenjingvisitorpaduploadfaceAPIRequest) GetId() string {
	return r._id
}

// SetImage is Image Setter
// 图片URL
func (r *AlibabaibshenjingvisitorpaduploadfaceAPIRequest) SetImage(_image string) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r AlibabaibshenjingvisitorpaduploadfaceAPIRequest) GetImage() string {
	return r._image
}
