package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppOfficialVerifyAPIRequest 聚安全验证官方应用接口 API请求
// alibaba.security.jaq.app.official.verify
//
// 接入用户来查询应用是否为官方应用
type AlibabaSecurityJaqAppOfficialVerifyAPIRequest struct {
	model.Params
	// 验证参数
	_officialAppVerifyRequest *OfficialAppVerifyRequest
}

// NewAlibabaSecurityJaqAppOfficialVerifyRequest 初始化AlibabaSecurityJaqAppOfficialVerifyAPIRequest对象
func NewAlibabaSecurityJaqAppOfficialVerifyRequest() *AlibabaSecurityJaqAppOfficialVerifyAPIRequest {
	return &AlibabaSecurityJaqAppOfficialVerifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqAppOfficialVerifyAPIRequest) Reset() {
	r._officialAppVerifyRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppOfficialVerifyAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.official.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppOfficialVerifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqAppOfficialVerifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfficialAppVerifyRequest is OfficialAppVerifyRequest Setter
// 验证参数
func (r *AlibabaSecurityJaqAppOfficialVerifyAPIRequest) SetOfficialAppVerifyRequest(_officialAppVerifyRequest *OfficialAppVerifyRequest) error {
	r._officialAppVerifyRequest = _officialAppVerifyRequest
	r.Set("official_app_verify_request", _officialAppVerifyRequest)
	return nil
}

// GetOfficialAppVerifyRequest OfficialAppVerifyRequest Getter
func (r AlibabaSecurityJaqAppOfficialVerifyAPIRequest) GetOfficialAppVerifyRequest() *OfficialAppVerifyRequest {
	return r._officialAppVerifyRequest
}

var poolAlibabaSecurityJaqAppOfficialVerifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqAppOfficialVerifyRequest()
	},
}

// GetAlibabaSecurityJaqAppOfficialVerifyRequest 从 sync.Pool 获取 AlibabaSecurityJaqAppOfficialVerifyAPIRequest
func GetAlibabaSecurityJaqAppOfficialVerifyAPIRequest() *AlibabaSecurityJaqAppOfficialVerifyAPIRequest {
	return poolAlibabaSecurityJaqAppOfficialVerifyAPIRequest.Get().(*AlibabaSecurityJaqAppOfficialVerifyAPIRequest)
}

// ReleaseAlibabaSecurityJaqAppOfficialVerifyAPIRequest 将 AlibabaSecurityJaqAppOfficialVerifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqAppOfficialVerifyAPIRequest(v *AlibabaSecurityJaqAppOfficialVerifyAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqAppOfficialVerifyAPIRequest.Put(v)
}
