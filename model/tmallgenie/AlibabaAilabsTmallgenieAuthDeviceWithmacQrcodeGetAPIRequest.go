package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest 根据mac查询设备的安全二维码 API请求
// alibaba.ailabs.tmallgenie.auth.device.withmac.qrcode.get
//
// 根据mac查询二维码详细信息
type AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest struct {
	model.Params
	// 产品ID
	_clientId string
	// 设备mac地址
	_mac string
}

// NewAlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetRequest 初始化AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetRequest() *AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.withmac.qrcode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// 产品ID
func (r *AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest) GetClientId() string {
	return r._clientId
}

// SetMac is Mac Setter
// 设备mac地址
func (r *AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// GetMac Mac Getter
func (r AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest) GetMac() string {
	return r._mac
}
