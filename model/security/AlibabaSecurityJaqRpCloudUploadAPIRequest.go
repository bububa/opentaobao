package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudUploadAPIRequest 实人认证云上传接口 API请求
// alibaba.security.jaq.rp.cloud.upload
//
// 聚安全实人认证上传认证信息
type AlibabaSecurityJaqRpCloudUploadAPIRequest struct {
	model.Params
	// []
	_elements []Elements
	// 认证token
	_verifyToken string
}

// NewAlibabaSecurityJaqRpCloudUploadRequest 初始化AlibabaSecurityJaqRpCloudUploadAPIRequest对象
func NewAlibabaSecurityJaqRpCloudUploadRequest() *AlibabaSecurityJaqRpCloudUploadAPIRequest {
	return &AlibabaSecurityJaqRpCloudUploadAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqRpCloudUploadAPIRequest) Reset() {
	r._elements = r._elements[:0]
	r._verifyToken = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpCloudUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetElements is Elements Setter
// []
func (r *AlibabaSecurityJaqRpCloudUploadAPIRequest) SetElements(_elements []Elements) error {
	r._elements = _elements
	r.Set("elements", _elements)
	return nil
}

// GetElements Elements Getter
func (r AlibabaSecurityJaqRpCloudUploadAPIRequest) GetElements() []Elements {
	return r._elements
}

// SetVerifyToken is VerifyToken Setter
// 认证token
func (r *AlibabaSecurityJaqRpCloudUploadAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudUploadAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

var poolAlibabaSecurityJaqRpCloudUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqRpCloudUploadRequest()
	},
}

// GetAlibabaSecurityJaqRpCloudUploadRequest 从 sync.Pool 获取 AlibabaSecurityJaqRpCloudUploadAPIRequest
func GetAlibabaSecurityJaqRpCloudUploadAPIRequest() *AlibabaSecurityJaqRpCloudUploadAPIRequest {
	return poolAlibabaSecurityJaqRpCloudUploadAPIRequest.Get().(*AlibabaSecurityJaqRpCloudUploadAPIRequest)
}

// ReleaseAlibabaSecurityJaqRpCloudUploadAPIRequest 将 AlibabaSecurityJaqRpCloudUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqRpCloudUploadAPIRequest(v *AlibabaSecurityJaqRpCloudUploadAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqRpCloudUploadAPIRequest.Put(v)
}
