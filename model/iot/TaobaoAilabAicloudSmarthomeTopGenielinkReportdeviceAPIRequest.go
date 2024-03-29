package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest 零配方案上报设备 API请求
// taobao.ailab.aicloud.smarthome.top.genielink.reportdevice
//
// 零配方案中设备联网成功之后上报设备
type TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest struct {
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

// NewTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest 初始化TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest对象
func NewTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest() *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest {
	return &TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest) Reset() {
	r._deviceId = ""
	r._status = ""
	r._extensions = ""
	r._vendorId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.smarthome.top.genielink.reportdevice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetStatus is Status Setter
// 设备状态，online上线，offline下线
func (r *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest) GetStatus() string {
	return r._status
}

// SetExtensions is Extensions Setter
// 保留字段json字符串
func (r *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest) SetExtensions(_extensions string) error {
	r._extensions = _extensions
	r.Set("extensions", _extensions)
	return nil
}

// GetExtensions Extensions Getter
func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest) GetExtensions() string {
	return r._extensions
}

// SetVendorId is VendorId Setter
// 供应商id
func (r *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest) SetVendorId(_vendorId int64) error {
	r._vendorId = _vendorId
	r.Set("vendor_id", _vendorId)
	return nil
}

// GetVendorId VendorId Getter
func (r TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest) GetVendorId() int64 {
	return r._vendorId
}

var poolTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest()
	},
}

// GetTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceRequest 从 sync.Pool 获取 TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest
func GetTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest() *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest {
	return poolTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest.Get().(*TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest)
}

// ReleaseTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest 将 TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest(v *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest.Put(v)
}
