package ott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvscreenlgelaunchergetAPIRequest LG用桌面信息获取 API请求
// yunos.tvscreen.lge.launcher.get
//
// LCTS一体机桌面后台,提供基于运营坑位适配的桌面服务,根据LG的需求定制输出格式
type YunostvscreenlgelaunchergetAPIRequest struct {
	model.Params
	// property of client
	_property string
	// ip of client
	_ip string
}

// NewYunostvscreenlgelaunchergetRequest 初始化YunostvscreenlgelaunchergetAPIRequest对象
func NewYunostvscreenlgelaunchergetRequest() *YunostvscreenlgelaunchergetAPIRequest {
	return &YunostvscreenlgelaunchergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvscreenlgelaunchergetAPIRequest) GetApiMethodName() string {
	return "yunos.tvscreen.lge.launcher.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvscreenlgelaunchergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvscreenlgelaunchergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProperty is Property Setter
// property of client
func (r *YunostvscreenlgelaunchergetAPIRequest) SetProperty(_property string) error {
	r._property = _property
	r.Set("property", _property)
	return nil
}

// GetProperty Property Getter
func (r YunostvscreenlgelaunchergetAPIRequest) GetProperty() string {
	return r._property
}

// SetIp is Ip Setter
// ip of client
func (r *YunostvscreenlgelaunchergetAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r YunostvscreenlgelaunchergetAPIRequest) GetIp() string {
	return r._ip
}
