package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqRpCloudStartAPIRequest 实人认证云开始认证 API请求
// alibaba.security.jaq.rp.cloud.start
//
// 聚安全实人认证开始
type AlibabaSecurityJaqRpCloudStartAPIRequest struct {
	model.Params
	// token
	_verifyToken string
	// 扩展信息
	_extraData string
	// 客户端信息，如果是服务端接入，里面的参数可为空
	_clientInfo *RpClientInfo
}

// NewAlibabaSecurityJaqRpCloudStartRequest 初始化AlibabaSecurityJaqRpCloudStartAPIRequest对象
func NewAlibabaSecurityJaqRpCloudStartRequest() *AlibabaSecurityJaqRpCloudStartAPIRequest {
	return &AlibabaSecurityJaqRpCloudStartAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqRpCloudStartAPIRequest) Reset() {
	r._verifyToken = ""
	r._extraData = ""
	r._clientInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudStartAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.start"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudStartAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqRpCloudStartAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpCloudStartAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudStartAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// SetExtraData is ExtraData Setter
// 扩展信息
func (r *AlibabaSecurityJaqRpCloudStartAPIRequest) SetExtraData(_extraData string) error {
	r._extraData = _extraData
	r.Set("extra_data", _extraData)
	return nil
}

// GetExtraData ExtraData Getter
func (r AlibabaSecurityJaqRpCloudStartAPIRequest) GetExtraData() string {
	return r._extraData
}

// SetClientInfo is ClientInfo Setter
// 客户端信息，如果是服务端接入，里面的参数可为空
func (r *AlibabaSecurityJaqRpCloudStartAPIRequest) SetClientInfo(_clientInfo *RpClientInfo) error {
	r._clientInfo = _clientInfo
	r.Set("client_info", _clientInfo)
	return nil
}

// GetClientInfo ClientInfo Getter
func (r AlibabaSecurityJaqRpCloudStartAPIRequest) GetClientInfo() *RpClientInfo {
	return r._clientInfo
}

var poolAlibabaSecurityJaqRpCloudStartAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqRpCloudStartRequest()
	},
}

// GetAlibabaSecurityJaqRpCloudStartRequest 从 sync.Pool 获取 AlibabaSecurityJaqRpCloudStartAPIRequest
func GetAlibabaSecurityJaqRpCloudStartAPIRequest() *AlibabaSecurityJaqRpCloudStartAPIRequest {
	return poolAlibabaSecurityJaqRpCloudStartAPIRequest.Get().(*AlibabaSecurityJaqRpCloudStartAPIRequest)
}

// ReleaseAlibabaSecurityJaqRpCloudStartAPIRequest 将 AlibabaSecurityJaqRpCloudStartAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqRpCloudStartAPIRequest(v *AlibabaSecurityJaqRpCloudStartAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqRpCloudStartAPIRequest.Put(v)
}
