package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIRequest 根据mac查询设备的安全二维码 API请求
// alibaba.ailabs.tmallgenie.auth.device.withmac.qrcode.get
//
// 根据mac查询二维码详细信息
type AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIRequest struct {
	model.Params
	// 产品ID
	_clientId string
	// 设备mac地址
	_mac string
}

// NewAlibabaailabstmallgenieauthdevicewithmacqrcodegetRequest 初始化AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIRequest对象
func NewAlibabaailabstmallgenieauthdevicewithmacqrcodegetRequest() *AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIRequest {
	return &AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.withmac.qrcode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// 产品ID
func (r *AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIRequest) GetClientId() string {
	return r._clientId
}

// SetMac is Mac Setter
// 设备mac地址
func (r *AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// GetMac Mac Getter
func (r AlibabaailabstmallgenieauthdevicewithmacqrcodegetAPIRequest) GetMac() string {
	return r._mac
}
