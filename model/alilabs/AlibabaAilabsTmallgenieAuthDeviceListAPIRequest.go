package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceListAPIRequest 获取用户设备列表 API请求
// alibaba.ailabs.tmallgenie.auth.device.list
//
// 通过此接口获取用户绑定的设备信息列表
type AlibabaAilabsTmallgenieAuthDeviceListAPIRequest struct {
	model.Params
	// 客户id
	_clientId string
	// 用户开放id
	_userOpenId string
}

// NewAlibabaAilabsTmallgenieAuthDeviceListRequest 初始化AlibabaAilabsTmallgenieAuthDeviceListAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceListRequest() *AlibabaAilabsTmallgenieAuthDeviceListAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetClientId is ClientId Setter
// 客户id
func (r *AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) GetClientId() string {
	return r._clientId
}

// SetUserOpenId is UserOpenId Setter
// 用户开放id
func (r *AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// GetUserOpenId UserOpenId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}
