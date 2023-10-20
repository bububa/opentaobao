package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpQueryAPIRequest 聚安全实人认证查询认证结果 API请求
// alibaba.security.jaq.rp.query
//
// 聚安全实人认证查询认证结果
type AlibabaSecurityJaqRpQueryAPIRequest struct {
	model.Params
	// token
	_verifyToken string
}

// NewAlibabaSecurityJaqRpQueryRequest 初始化AlibabaSecurityJaqRpQueryAPIRequest对象
func NewAlibabaSecurityJaqRpQueryRequest() *AlibabaSecurityJaqRpQueryAPIRequest {
	return &AlibabaSecurityJaqRpQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqRpQueryAPIRequest) Reset() {
	r._verifyToken = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpQueryAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabaSecurityJaqRpQueryAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

var poolAlibabaSecurityJaqRpQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqRpQueryRequest()
	},
}

// GetAlibabaSecurityJaqRpQueryRequest 从 sync.Pool 获取 AlibabaSecurityJaqRpQueryAPIRequest
func GetAlibabaSecurityJaqRpQueryAPIRequest() *AlibabaSecurityJaqRpQueryAPIRequest {
	return poolAlibabaSecurityJaqRpQueryAPIRequest.Get().(*AlibabaSecurityJaqRpQueryAPIRequest)
}

// ReleaseAlibabaSecurityJaqRpQueryAPIRequest 将 AlibabaSecurityJaqRpQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqRpQueryAPIRequest(v *AlibabaSecurityJaqRpQueryAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqRpQueryAPIRequest.Put(v)
}
