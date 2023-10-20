package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceapksAPIRequest 获取停开服apk列表 API请求
// yunos.tvpubadmin.device.apks
//
// 获取停开服apk列表
type YunostvpubadmindeviceapksAPIRequest struct {
	model.Params
	// 牌照
	_license int64
}

// NewYunostvpubadmindeviceapksRequest 初始化YunostvpubadmindeviceapksAPIRequest对象
func NewYunostvpubadmindeviceapksRequest() *YunostvpubadmindeviceapksAPIRequest {
	return &YunostvpubadmindeviceapksAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindeviceapksAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.apks"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindeviceapksAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindeviceapksAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLicense is License Setter
// 牌照
func (r *YunostvpubadmindeviceapksAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmindeviceapksAPIRequest) GetLicense() int64 {
	return r._license
}
