package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIbShenjingVisitorPadUploadfaceAPIRequest 访客PAD上传人脸 API请求
// alibaba.ib.shenjing.visitor.pad.uploadface
//
// 访客PAD端上传人脸。
type AlibabaIbShenjingVisitorPadUploadfaceAPIRequest struct {
	model.Params
	// 访客ID
	_id string
	// 图片URL
	_image string
}

// NewAlibabaIbShenjingVisitorPadUploadfaceRequest 初始化AlibabaIbShenjingVisitorPadUploadfaceAPIRequest对象
func NewAlibabaIbShenjingVisitorPadUploadfaceRequest() *AlibabaIbShenjingVisitorPadUploadfaceAPIRequest {
	return &AlibabaIbShenjingVisitorPadUploadfaceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIbShenjingVisitorPadUploadfaceAPIRequest) GetApiMethodName() string {
	return "alibaba.ib.shenjing.visitor.pad.uploadface"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIbShenjingVisitorPadUploadfaceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetId is Id Setter
// 访客ID
func (r *AlibabaIbShenjingVisitorPadUploadfaceAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaIbShenjingVisitorPadUploadfaceAPIRequest) GetId() string {
	return r._id
}

// SetImage is Image Setter
// 图片URL
func (r *AlibabaIbShenjingVisitorPadUploadfaceAPIRequest) SetImage(_image string) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r AlibabaIbShenjingVisitorPadUploadfaceAPIRequest) GetImage() string {
	return r._image
}
