package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpocrAPIRequest 聚安全实人认证证件OCR识别功能接口 API请求
// alibaba.security.jaq.rp.ocr
//
// 聚安全实人认证证件OCR识别功能接口
type AlibabasecurityjaqrpocrAPIRequest struct {
	model.Params
	// token
	_verifyToken string
	// 要识别的信息
	_imageUrls string
}

// NewAlibabasecurityjaqrpocrRequest 初始化AlibabasecurityjaqrpocrAPIRequest对象
func NewAlibabasecurityjaqrpocrRequest() *AlibabasecurityjaqrpocrAPIRequest {
	return &AlibabasecurityjaqrpocrAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqrpocrAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.ocr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqrpocrAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqrpocrAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// token
func (r *AlibabasecurityjaqrpocrAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabasecurityjaqrpocrAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// SetImageUrls is ImageUrls Setter
// 要识别的信息
func (r *AlibabasecurityjaqrpocrAPIRequest) SetImageUrls(_imageUrls string) error {
	r._imageUrls = _imageUrls
	r.Set("image_urls", _imageUrls)
	return nil
}

// GetImageUrls ImageUrls Getter
func (r AlibabasecurityjaqrpocrAPIRequest) GetImageUrls() string {
	return r._imageUrls
}
