package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqRpOcrAPIRequest
聚安全实人认证证件OCR识别功能接口 API请求
alibaba.security.jaq.rp.ocr

聚安全实人认证证件OCR识别功能接口 */
type AlibabaSecurityJaqRpOcrAPIRequest struct {
	model.Params
	// token
	_verifyToken string
	// 要识别的信息
	_imageUrls string
}

// NewAlibabaSecurityJaqRpOcrRequest 初始化AlibabaSecurityJaqRpOcrAPIRequest对象
func NewAlibabaSecurityJaqRpOcrRequest() *AlibabaSecurityJaqRpOcrAPIRequest {
	return &AlibabaSecurityJaqRpOcrAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpOcrAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.ocr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpOcrAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpOcrAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// Get VerifyToken Getter
func (r AlibabaSecurityJaqRpOcrAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// Set is ImageUrls Setter
// 要识别的信息
func (r *AlibabaSecurityJaqRpOcrAPIRequest) SetImageUrls(_imageUrls string) error {
	r._imageUrls = _imageUrls
	r.Set("image_urls", _imageUrls)
	return nil
}

// Get ImageUrls Getter
func (r AlibabaSecurityJaqRpOcrAPIRequest) GetImageUrls() string {
	return r._imageUrls
}
