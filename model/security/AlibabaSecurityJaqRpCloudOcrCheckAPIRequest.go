package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudOcrCheckAPIRequest ocr同时实名校验 API请求
// alibaba.security.jaq.rp.cloud.ocr.check
//
// 聚安全实人认证证件OCR识别功能接口
type AlibabaSecurityJaqRpCloudOcrCheckAPIRequest struct {
	model.Params
	// token
	_verifyToken string
	// 要识别的信息
	_imageUrls string
}

// NewAlibabaSecurityJaqRpCloudOcrCheckRequest 初始化AlibabaSecurityJaqRpCloudOcrCheckAPIRequest对象
func NewAlibabaSecurityJaqRpCloudOcrCheckRequest() *AlibabaSecurityJaqRpCloudOcrCheckAPIRequest {
	return &AlibabaSecurityJaqRpCloudOcrCheckAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) Reset() {
	r._verifyToken = ""
	r._imageUrls = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.ocr.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// SetImageUrls is ImageUrls Setter
// 要识别的信息
func (r *AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) SetImageUrls(_imageUrls string) error {
	r._imageUrls = _imageUrls
	r.Set("image_urls", _imageUrls)
	return nil
}

// GetImageUrls ImageUrls Getter
func (r AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) GetImageUrls() string {
	return r._imageUrls
}

var poolAlibabaSecurityJaqRpCloudOcrCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqRpCloudOcrCheckRequest()
	},
}

// GetAlibabaSecurityJaqRpCloudOcrCheckRequest 从 sync.Pool 获取 AlibabaSecurityJaqRpCloudOcrCheckAPIRequest
func GetAlibabaSecurityJaqRpCloudOcrCheckAPIRequest() *AlibabaSecurityJaqRpCloudOcrCheckAPIRequest {
	return poolAlibabaSecurityJaqRpCloudOcrCheckAPIRequest.Get().(*AlibabaSecurityJaqRpCloudOcrCheckAPIRequest)
}

// ReleaseAlibabaSecurityJaqRpCloudOcrCheckAPIRequest 将 AlibabaSecurityJaqRpCloudOcrCheckAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqRpCloudOcrCheckAPIRequest(v *AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqRpCloudOcrCheckAPIRequest.Put(v)
}
