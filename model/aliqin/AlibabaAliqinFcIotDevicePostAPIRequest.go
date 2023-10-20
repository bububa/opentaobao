package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfciotdevicepostAPIRequest 商家提交设备信息 API请求
// alibaba.aliqin.fc.iot.device.post
//
// 物联网商家设备信息录入
type AlibabaaliqinfciotdevicepostAPIRequest struct {
	model.Params
	// 15位imei号
	_imei string
	// 设备类型（将来扩展）
	_deviceType string
	// 备注
	_comments string
	// 扩展字段
	_midPatChannel string
}

// NewAlibabaaliqinfciotdevicepostRequest 初始化AlibabaaliqinfciotdevicepostAPIRequest对象
func NewAlibabaaliqinfciotdevicepostRequest() *AlibabaaliqinfciotdevicepostAPIRequest {
	return &AlibabaaliqinfciotdevicepostAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinfciotdevicepostAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.device.post"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinfciotdevicepostAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinfciotdevicepostAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImei is Imei Setter
// 15位imei号
func (r *AlibabaaliqinfciotdevicepostAPIRequest) SetImei(_imei string) error {
	r._imei = _imei
	r.Set("imei", _imei)
	return nil
}

// GetImei Imei Getter
func (r AlibabaaliqinfciotdevicepostAPIRequest) GetImei() string {
	return r._imei
}

// SetDeviceType is DeviceType Setter
// 设备类型（将来扩展）
func (r *AlibabaaliqinfciotdevicepostAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r AlibabaaliqinfciotdevicepostAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetComments is Comments Setter
// 备注
func (r *AlibabaaliqinfciotdevicepostAPIRequest) SetComments(_comments string) error {
	r._comments = _comments
	r.Set("comments", _comments)
	return nil
}

// GetComments Comments Getter
func (r AlibabaaliqinfciotdevicepostAPIRequest) GetComments() string {
	return r._comments
}

// SetMidPatChannel is MidPatChannel Setter
// 扩展字段
func (r *AlibabaaliqinfciotdevicepostAPIRequest) SetMidPatChannel(_midPatChannel string) error {
	r._midPatChannel = _midPatChannel
	r.Set("mid_pat_channel", _midPatChannel)
	return nil
}

// GetMidPatChannel MidPatChannel Getter
func (r AlibabaaliqinfciotdevicepostAPIRequest) GetMidPatChannel() string {
	return r._midPatChannel
}
