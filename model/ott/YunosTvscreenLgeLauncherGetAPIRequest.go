package ott

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvscreenLgeLauncherGetAPIRequest LG用桌面信息获取 API请求
// yunos.tvscreen.lge.launcher.get
//
// LCTS一体机桌面后台,提供基于运营坑位适配的桌面服务,根据LG的需求定制输出格式
type YunosTvscreenLgeLauncherGetAPIRequest struct {
	model.Params
	// property of client
	_property string
	// ip of client
	_ip string
}

// NewYunosTvscreenLgeLauncherGetRequest 初始化YunosTvscreenLgeLauncherGetAPIRequest对象
func NewYunosTvscreenLgeLauncherGetRequest() *YunosTvscreenLgeLauncherGetAPIRequest {
	return &YunosTvscreenLgeLauncherGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvscreenLgeLauncherGetAPIRequest) Reset() {
	r._property = ""
	r._ip = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvscreenLgeLauncherGetAPIRequest) GetApiMethodName() string {
	return "yunos.tvscreen.lge.launcher.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvscreenLgeLauncherGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvscreenLgeLauncherGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProperty is Property Setter
// property of client
func (r *YunosTvscreenLgeLauncherGetAPIRequest) SetProperty(_property string) error {
	r._property = _property
	r.Set("property", _property)
	return nil
}

// GetProperty Property Getter
func (r YunosTvscreenLgeLauncherGetAPIRequest) GetProperty() string {
	return r._property
}

// SetIp is Ip Setter
// ip of client
func (r *YunosTvscreenLgeLauncherGetAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r YunosTvscreenLgeLauncherGetAPIRequest) GetIp() string {
	return r._ip
}

var poolYunosTvscreenLgeLauncherGetAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvscreenLgeLauncherGetRequest()
	},
}

// GetYunosTvscreenLgeLauncherGetRequest 从 sync.Pool 获取 YunosTvscreenLgeLauncherGetAPIRequest
func GetYunosTvscreenLgeLauncherGetAPIRequest() *YunosTvscreenLgeLauncherGetAPIRequest {
	return poolYunosTvscreenLgeLauncherGetAPIRequest.Get().(*YunosTvscreenLgeLauncherGetAPIRequest)
}

// ReleaseYunosTvscreenLgeLauncherGetAPIRequest 将 YunosTvscreenLgeLauncherGetAPIRequest 放入 sync.Pool
func ReleaseYunosTvscreenLgeLauncherGetAPIRequest(v *YunosTvscreenLgeLauncherGetAPIRequest) {
	v.Reset()
	poolYunosTvscreenLgeLauncherGetAPIRequest.Put(v)
}
