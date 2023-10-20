package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkRexoutDeviceIotRegisteridAPIRequest 通过设备ID获取三元组-外部 API请求
// wdk.rexout.device.iot.registerid
//
// 通过设备ID获取三元组-外部
type WdkRexoutDeviceIotRegisteridAPIRequest struct {
	model.Params
	// 平台key
	_accessKey string
	// 平台签名
	_accessSign string
	// 签名时间戳,毫秒,加密时使用时间加密则必传
	_accessTime string
	// 设备ID
	_uuid string
}

// NewWdkRexoutDeviceIotRegisteridRequest 初始化WdkRexoutDeviceIotRegisteridAPIRequest对象
func NewWdkRexoutDeviceIotRegisteridRequest() *WdkRexoutDeviceIotRegisteridAPIRequest {
	return &WdkRexoutDeviceIotRegisteridAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkRexoutDeviceIotRegisteridAPIRequest) GetApiMethodName() string {
	return "wdk.rexout.device.iot.registerid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkRexoutDeviceIotRegisteridAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkRexoutDeviceIotRegisteridAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessKey is AccessKey Setter
// 平台key
func (r *WdkRexoutDeviceIotRegisteridAPIRequest) SetAccessKey(_accessKey string) error {
	r._accessKey = _accessKey
	r.Set("access_key", _accessKey)
	return nil
}

// GetAccessKey AccessKey Getter
func (r WdkRexoutDeviceIotRegisteridAPIRequest) GetAccessKey() string {
	return r._accessKey
}

// SetAccessSign is AccessSign Setter
// 平台签名
func (r *WdkRexoutDeviceIotRegisteridAPIRequest) SetAccessSign(_accessSign string) error {
	r._accessSign = _accessSign
	r.Set("access_sign", _accessSign)
	return nil
}

// GetAccessSign AccessSign Getter
func (r WdkRexoutDeviceIotRegisteridAPIRequest) GetAccessSign() string {
	return r._accessSign
}

// SetAccessTime is AccessTime Setter
// 签名时间戳,毫秒,加密时使用时间加密则必传
func (r *WdkRexoutDeviceIotRegisteridAPIRequest) SetAccessTime(_accessTime string) error {
	r._accessTime = _accessTime
	r.Set("access_time", _accessTime)
	return nil
}

// GetAccessTime AccessTime Getter
func (r WdkRexoutDeviceIotRegisteridAPIRequest) GetAccessTime() string {
	return r._accessTime
}

// SetUuid is Uuid Setter
// 设备ID
func (r *WdkRexoutDeviceIotRegisteridAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r WdkRexoutDeviceIotRegisteridAPIRequest) GetUuid() string {
	return r._uuid
}
