package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkrexoutdeviceinfogetAPIRequest 获取设备详情-外部对接 API请求
// wdk.rexout.device.info.get
//
// 获取设备详情-外部对接
type WdkrexoutdeviceinfogetAPIRequest struct {
	model.Params
	// ak
	_accessKey string
	// 签名
	_accessSign string
	// 设备ID
	_uuid string
}

// NewWdkrexoutdeviceinfogetRequest 初始化WdkrexoutdeviceinfogetAPIRequest对象
func NewWdkrexoutdeviceinfogetRequest() *WdkrexoutdeviceinfogetAPIRequest {
	return &WdkrexoutdeviceinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkrexoutdeviceinfogetAPIRequest) GetApiMethodName() string {
	return "wdk.rexout.device.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkrexoutdeviceinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkrexoutdeviceinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessKey is AccessKey Setter
// ak
func (r *WdkrexoutdeviceinfogetAPIRequest) SetAccessKey(_accessKey string) error {
	r._accessKey = _accessKey
	r.Set("access_key", _accessKey)
	return nil
}

// GetAccessKey AccessKey Getter
func (r WdkrexoutdeviceinfogetAPIRequest) GetAccessKey() string {
	return r._accessKey
}

// SetAccessSign is AccessSign Setter
// 签名
func (r *WdkrexoutdeviceinfogetAPIRequest) SetAccessSign(_accessSign string) error {
	r._accessSign = _accessSign
	r.Set("access_sign", _accessSign)
	return nil
}

// GetAccessSign AccessSign Getter
func (r WdkrexoutdeviceinfogetAPIRequest) GetAccessSign() string {
	return r._accessSign
}

// SetUuid is Uuid Setter
// 设备ID
func (r *WdkrexoutdeviceinfogetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r WdkrexoutdeviceinfogetAPIRequest) GetUuid() string {
	return r._uuid
}
