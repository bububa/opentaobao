package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceApksAPIRequest 获取停开服apk列表 API请求
// yunos.tvpubadmin.device.apks
//
// 获取停开服apk列表
type YunosTvpubadminDeviceApksAPIRequest struct {
	model.Params
	// 牌照
	_license int64
}

// NewYunosTvpubadminDeviceApksRequest 初始化YunosTvpubadminDeviceApksAPIRequest对象
func NewYunosTvpubadminDeviceApksRequest() *YunosTvpubadminDeviceApksAPIRequest {
	return &YunosTvpubadminDeviceApksAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceApksAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.apks"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceApksAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceApksAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLicense is License Setter
// 牌照
func (r *YunosTvpubadminDeviceApksAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminDeviceApksAPIRequest) GetLicense() int64 {
	return r._license
}
