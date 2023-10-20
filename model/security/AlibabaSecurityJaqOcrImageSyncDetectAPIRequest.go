package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqocrimagesyncdetectAPIRequest 聚安全图文识别同步检测接口 API请求
// alibaba.security.jaq.ocr.image.sync.detect
//
// 图像字符识别同步检测接口
type AlibabasecurityjaqocrimagesyncdetectAPIRequest struct {
	model.Params
	// 待检测图像链接
	_imageUrl string
}

// NewAlibabasecurityjaqocrimagesyncdetectRequest 初始化AlibabasecurityjaqocrimagesyncdetectAPIRequest对象
func NewAlibabasecurityjaqocrimagesyncdetectRequest() *AlibabasecurityjaqocrimagesyncdetectAPIRequest {
	return &AlibabasecurityjaqocrimagesyncdetectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqocrimagesyncdetectAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.ocr.image.sync.detect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqocrimagesyncdetectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqocrimagesyncdetectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 待检测图像链接
func (r *AlibabasecurityjaqocrimagesyncdetectAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AlibabasecurityjaqocrimagesyncdetectAPIRequest) GetImageUrl() string {
	return r._imageUrl
}
