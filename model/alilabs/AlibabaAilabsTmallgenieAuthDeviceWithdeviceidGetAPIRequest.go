package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIRequest 根据三方ID查询设备注册激活信息 API请求
// alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get
//
// 根据三方ID查询设备注册激活信息
type AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIRequest struct {
	model.Params
	// 设备产品ID
	_clientId string
	// mac地址
	_mac string
}

// NewAlibabaailabstmallgenieauthdevicewithdeviceidgetRequest 初始化AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIRequest对象
func NewAlibabaailabstmallgenieauthdevicewithdeviceidgetRequest() *AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIRequest {
	return &AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// 设备产品ID
func (r *AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIRequest) GetClientId() string {
	return r._clientId
}

// SetMac is Mac Setter
// mac地址
func (r *AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// GetMac Mac Getter
func (r AlibabaailabstmallgenieauthdevicewithdeviceidgetAPIRequest) GetMac() string {
	return r._mac
}
