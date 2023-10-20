package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpOcrCheckAPIRequest ocr同时实名校验 API请求
// alibaba.security.jaq.rp.ocr.check
//
// 聚安全实人认证证件OCR识别功能接口
type AlibabaSecurityJaqRpOcrCheckAPIRequest struct {
	model.Params
	// token
	_verifyToken string
	// 要识别的信息
	_imageUrls string
}

// NewAlibabaSecurityJaqRpOcrCheckRequest 初始化AlibabaSecurityJaqRpOcrCheckAPIRequest对象
func NewAlibabaSecurityJaqRpOcrCheckRequest() *AlibabaSecurityJaqRpOcrCheckAPIRequest {
	return &AlibabaSecurityJaqRpOcrCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpOcrCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.ocr.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpOcrCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpOcrCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpOcrCheckAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabaSecurityJaqRpOcrCheckAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// SetImageUrls is ImageUrls Setter
// 要识别的信息
func (r *AlibabaSecurityJaqRpOcrCheckAPIRequest) SetImageUrls(_imageUrls string) error {
	r._imageUrls = _imageUrls
	r.Set("image_urls", _imageUrls)
	return nil
}

// GetImageUrls ImageUrls Getter
func (r AlibabaSecurityJaqRpOcrCheckAPIRequest) GetImageUrls() string {
	return r._imageUrls
}
