package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunviapiimagesegsegmentheadAPIRequest 头像分割 API请求
// aliyun.viapi.imageseg.segmenthead
//
// 输入一张图片，对图中人头区域进行抠图解析，输出人头png透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunviapiimagesegsegmentheadAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// NewAliyunviapiimagesegsegmentheadRequest 初始化AliyunviapiimagesegsegmentheadAPIRequest对象
func NewAliyunviapiimagesegsegmentheadRequest() *AliyunviapiimagesegsegmentheadAPIRequest {
	return &AliyunviapiimagesegsegmentheadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunviapiimagesegsegmentheadAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.imageseg.segmenthead"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunviapiimagesegsegmentheadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunviapiimagesegsegmentheadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunviapiimagesegsegmentheadAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunviapiimagesegsegmentheadAPIRequest) GetImageUrl() string {
	return r._imageUrl
}
