package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest 零配方案上报设备 API请求
// taobao.ailab.aicloud.smarthome.top.genielink.reportdevice
//
// 零配方案中设备联网成功之后上报设备
type TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 设备状态，online上线，offline下线
	_status string
	// 保留字段json字符串
	_extensions string
	// 供应商id
	_vendorId int64
}

// NewTaobaoailabaicloudsmarthometopgenielinkreportdeviceRequest 初始化TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest对象
func NewTaobaoailabaicloudsmarthometopgenielinkreportdeviceRequest() *TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest {
	return &TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.smarthome.top.genielink.reportdevice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetStatus is Status Setter
// 设备状态，online上线，offline下线
func (r *TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest) GetStatus() string {
	return r._status
}

// SetExtensions is Extensions Setter
// 保留字段json字符串
func (r *TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest) SetExtensions(_extensions string) error {
	r._extensions = _extensions
	r.Set("extensions", _extensions)
	return nil
}

// GetExtensions Extensions Getter
func (r TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest) GetExtensions() string {
	return r._extensions
}

// SetVendorId is VendorId Setter
// 供应商id
func (r *TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest) SetVendorId(_vendorId int64) error {
	r._vendorId = _vendorId
	r.Set("vendor_id", _vendorId)
	return nil
}

// GetVendorId VendorId Getter
func (r TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest) GetVendorId() int64 {
	return r._vendorId
}
