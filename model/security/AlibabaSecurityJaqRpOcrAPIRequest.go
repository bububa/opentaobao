package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpOcrAPIRequest 聚安全实人认证证件OCR识别功能接口 API请求
// alibaba.security.jaq.rp.ocr
//
// 聚安全实人认证证件OCR识别功能接口
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqRpOcrAPIRequest) Reset() {
	r._verifyToken = ""
	r._imageUrls = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpOcrAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.ocr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpOcrAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpOcrAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpOcrAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabaSecurityJaqRpOcrAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// SetImageUrls is ImageUrls Setter
// 要识别的信息
func (r *AlibabaSecurityJaqRpOcrAPIRequest) SetImageUrls(_imageUrls string) error {
	r._imageUrls = _imageUrls
	r.Set("image_urls", _imageUrls)
	return nil
}

// GetImageUrls ImageUrls Getter
func (r AlibabaSecurityJaqRpOcrAPIRequest) GetImageUrls() string {
	return r._imageUrls
}

var poolAlibabaSecurityJaqRpOcrAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqRpOcrRequest()
	},
}

// GetAlibabaSecurityJaqRpOcrRequest 从 sync.Pool 获取 AlibabaSecurityJaqRpOcrAPIRequest
func GetAlibabaSecurityJaqRpOcrAPIRequest() *AlibabaSecurityJaqRpOcrAPIRequest {
	return poolAlibabaSecurityJaqRpOcrAPIRequest.Get().(*AlibabaSecurityJaqRpOcrAPIRequest)
}

// ReleaseAlibabaSecurityJaqRpOcrAPIRequest 将 AlibabaSecurityJaqRpOcrAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqRpOcrAPIRequest(v *AlibabaSecurityJaqRpOcrAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqRpOcrAPIRequest.Put(v)
}
