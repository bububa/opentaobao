package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqPornImageSyncDetectAPIRequest 聚安全智能鉴黄同步检测接口 API请求
// alibaba.security.jaq.porn.image.sync.detect
//
// 同步黄图图像检测接口
type AlibabaSecurityJaqPornImageSyncDetectAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// NewAlibabaSecurityJaqPornImageSyncDetectRequest 初始化AlibabaSecurityJaqPornImageSyncDetectAPIRequest对象
func NewAlibabaSecurityJaqPornImageSyncDetectRequest() *AlibabaSecurityJaqPornImageSyncDetectAPIRequest {
	return &AlibabaSecurityJaqPornImageSyncDetectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqPornImageSyncDetectAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.porn.image.sync.detect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqPornImageSyncDetectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqPornImageSyncDetectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AlibabaSecurityJaqPornImageSyncDetectAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AlibabaSecurityJaqPornImageSyncDetectAPIRequest) GetImageUrl() string {
	return r._imageUrl
}
