package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfciotdeviceisexistAPIRequest 判断设备是否存在 API请求
// alibaba.aliqin.fc.iot.device.isexist
//
// 判断设备是否存在
type AlibabaaliqinfciotdeviceisexistAPIRequest struct {
	model.Params
	// 设备编号
	_imei string
	// 设备类型（预留将来扩展）
	_deviceType string
	// 渠道扩展编码（预留）
	_midPatChannel string
}

// NewAlibabaaliqinfciotdeviceisexistRequest 初始化AlibabaaliqinfciotdeviceisexistAPIRequest对象
func NewAlibabaaliqinfciotdeviceisexistRequest() *AlibabaaliqinfciotdeviceisexistAPIRequest {
	return &AlibabaaliqinfciotdeviceisexistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinfciotdeviceisexistAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.device.isexist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinfciotdeviceisexistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinfciotdeviceisexistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImei is Imei Setter
// 设备编号
func (r *AlibabaaliqinfciotdeviceisexistAPIRequest) SetImei(_imei string) error {
	r._imei = _imei
	r.Set("imei", _imei)
	return nil
}

// GetImei Imei Getter
func (r AlibabaaliqinfciotdeviceisexistAPIRequest) GetImei() string {
	return r._imei
}

// SetDeviceType is DeviceType Setter
// 设备类型（预留将来扩展）
func (r *AlibabaaliqinfciotdeviceisexistAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r AlibabaaliqinfciotdeviceisexistAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetMidPatChannel is MidPatChannel Setter
// 渠道扩展编码（预留）
func (r *AlibabaaliqinfciotdeviceisexistAPIRequest) SetMidPatChannel(_midPatChannel string) error {
	r._midPatChannel = _midPatChannel
	r.Set("mid_pat_channel", _midPatChannel)
	return nil
}

// GetMidPatChannel MidPatChannel Getter
func (r AlibabaaliqinfciotdeviceisexistAPIRequest) GetMidPatChannel() string {
	return r._midPatChannel
}
