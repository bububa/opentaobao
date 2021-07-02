package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqOcrImageSyncDetectAPIRequest 聚安全图文识别同步检测接口 API请求
// alibaba.security.jaq.ocr.image.sync.detect
//
// 图像字符识别同步检测接口
type AlibabaSecurityJaqOcrImageSyncDetectAPIRequest struct {
	model.Params
	// 待检测图像链接
	_imageUrl string
}

// NewAlibabaSecurityJaqOcrImageSyncDetectRequest 初始化AlibabaSecurityJaqOcrImageSyncDetectAPIRequest对象
func NewAlibabaSecurityJaqOcrImageSyncDetectRequest() *AlibabaSecurityJaqOcrImageSyncDetectAPIRequest {
	return &AlibabaSecurityJaqOcrImageSyncDetectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqOcrImageSyncDetectAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.ocr.image.sync.detect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqOcrImageSyncDetectAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ImageUrl Setter
// 待检测图像链接
func (r *AlibabaSecurityJaqOcrImageSyncDetectAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// Get ImageUrl Getter
func (r AlibabaSecurityJaqOcrImageSyncDetectAPIRequest) GetImageUrl() string {
	return r._imageUrl
}
