package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpStartAPIRequest 聚安全实人认证开始 API请求
// alibaba.security.jaq.rp.start
//
// 聚安全实人认证开始
type AlibabaSecurityJaqRpStartAPIRequest struct {
	model.Params
	// token
	_verifyToken string
	// 扩展信息
	_extraData string
	// 客户端信息，如果是服务端接入，里面的参数可为空
	_clientInfo *RpClientInfo
}

// NewAlibabaSecurityJaqRpStartRequest 初始化AlibabaSecurityJaqRpStartAPIRequest对象
func NewAlibabaSecurityJaqRpStartRequest() *AlibabaSecurityJaqRpStartAPIRequest {
	return &AlibabaSecurityJaqRpStartAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqRpStartAPIRequest) Reset() {
	r._verifyToken = ""
	r._extraData = ""
	r._clientInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpStartAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.start"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpStartAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpStartAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpStartAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabaSecurityJaqRpStartAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// SetExtraData is ExtraData Setter
// 扩展信息
func (r *AlibabaSecurityJaqRpStartAPIRequest) SetExtraData(_extraData string) error {
	r._extraData = _extraData
	r.Set("extra_data", _extraData)
	return nil
}

// GetExtraData ExtraData Getter
func (r AlibabaSecurityJaqRpStartAPIRequest) GetExtraData() string {
	return r._extraData
}

// SetClientInfo is ClientInfo Setter
// 客户端信息，如果是服务端接入，里面的参数可为空
func (r *AlibabaSecurityJaqRpStartAPIRequest) SetClientInfo(_clientInfo *RpClientInfo) error {
	r._clientInfo = _clientInfo
	r.Set("client_info", _clientInfo)
	return nil
}

// GetClientInfo ClientInfo Getter
func (r AlibabaSecurityJaqRpStartAPIRequest) GetClientInfo() *RpClientInfo {
	return r._clientInfo
}

var poolAlibabaSecurityJaqRpStartAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqRpStartRequest()
	},
}

// GetAlibabaSecurityJaqRpStartRequest 从 sync.Pool 获取 AlibabaSecurityJaqRpStartAPIRequest
func GetAlibabaSecurityJaqRpStartAPIRequest() *AlibabaSecurityJaqRpStartAPIRequest {
	return poolAlibabaSecurityJaqRpStartAPIRequest.Get().(*AlibabaSecurityJaqRpStartAPIRequest)
}

// ReleaseAlibabaSecurityJaqRpStartAPIRequest 将 AlibabaSecurityJaqRpStartAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqRpStartAPIRequest(v *AlibabaSecurityJaqRpStartAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqRpStartAPIRequest.Put(v)
}
