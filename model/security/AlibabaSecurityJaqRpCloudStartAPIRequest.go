package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpcloudstartAPIRequest 实人认证云开始认证 API请求
// alibaba.security.jaq.rp.cloud.start
//
// 聚安全实人认证开始
type AlibabasecurityjaqrpcloudstartAPIRequest struct {
	model.Params
	// token
	_verifyToken string
	// 扩展信息
	_extraData string
	// 客户端信息，如果是服务端接入，里面的参数可为空
	_clientInfo *RpClientInfo
}

// NewAlibabasecurityjaqrpcloudstartRequest 初始化AlibabasecurityjaqrpcloudstartAPIRequest对象
func NewAlibabasecurityjaqrpcloudstartRequest() *AlibabasecurityjaqrpcloudstartAPIRequest {
	return &AlibabasecurityjaqrpcloudstartAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqrpcloudstartAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.cloud.start"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqrpcloudstartAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqrpcloudstartAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// token
func (r *AlibabasecurityjaqrpcloudstartAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabasecurityjaqrpcloudstartAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// SetExtraData is ExtraData Setter
// 扩展信息
func (r *AlibabasecurityjaqrpcloudstartAPIRequest) SetExtraData(_extraData string) error {
	r._extraData = _extraData
	r.Set("extra_data", _extraData)
	return nil
}

// GetExtraData ExtraData Getter
func (r AlibabasecurityjaqrpcloudstartAPIRequest) GetExtraData() string {
	return r._extraData
}

// SetClientInfo is ClientInfo Setter
// 客户端信息，如果是服务端接入，里面的参数可为空
func (r *AlibabasecurityjaqrpcloudstartAPIRequest) SetClientInfo(_clientInfo *RpClientInfo) error {
	r._clientInfo = _clientInfo
	r.Set("client_info", _clientInfo)
	return nil
}

// GetClientInfo ClientInfo Getter
func (r AlibabasecurityjaqrpcloudstartAPIRequest) GetClientInfo() *RpClientInfo {
	return r._clientInfo
}
