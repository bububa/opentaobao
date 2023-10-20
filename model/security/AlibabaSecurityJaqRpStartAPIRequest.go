package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqrpstartAPIRequest 聚安全实人认证开始 API请求
// alibaba.security.jaq.rp.start
//
// 聚安全实人认证开始
type AlibabasecurityjaqrpstartAPIRequest struct {
	model.Params
	// token
	_verifyToken string
	// 扩展信息
	_extraData string
	// 客户端信息，如果是服务端接入，里面的参数可为空
	_clientInfo *RpClientInfo
}

// NewAlibabasecurityjaqrpstartRequest 初始化AlibabasecurityjaqrpstartAPIRequest对象
func NewAlibabasecurityjaqrpstartRequest() *AlibabasecurityjaqrpstartAPIRequest {
	return &AlibabasecurityjaqrpstartAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqrpstartAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.rp.start"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqrpstartAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqrpstartAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyToken is VerifyToken Setter
// token
func (r *AlibabasecurityjaqrpstartAPIRequest) SetVerifyToken(_verifyToken string) error {
	r._verifyToken = _verifyToken
	r.Set("verify_token", _verifyToken)
	return nil
}

// GetVerifyToken VerifyToken Getter
func (r AlibabasecurityjaqrpstartAPIRequest) GetVerifyToken() string {
	return r._verifyToken
}

// SetExtraData is ExtraData Setter
// 扩展信息
func (r *AlibabasecurityjaqrpstartAPIRequest) SetExtraData(_extraData string) error {
	r._extraData = _extraData
	r.Set("extra_data", _extraData)
	return nil
}

// GetExtraData ExtraData Getter
func (r AlibabasecurityjaqrpstartAPIRequest) GetExtraData() string {
	return r._extraData
}

// SetClientInfo is ClientInfo Setter
// 客户端信息，如果是服务端接入，里面的参数可为空
func (r *AlibabasecurityjaqrpstartAPIRequest) SetClientInfo(_clientInfo *RpClientInfo) error {
	r._clientInfo = _clientInfo
	r.Set("client_info", _clientInfo)
	return nil
}

// GetClientInfo ClientInfo Getter
func (r AlibabasecurityjaqrpstartAPIRequest) GetClientInfo() *RpClientInfo {
	return r._clientInfo
}
