package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsaligenieiotdevicecontrolresultAPIRequest 设备控制结果 API请求
// alibaba.ailabs.aligenie.iot.device.control.result
//
// 智能IOT解决外部厂商在云云模式在用户通过天猫精灵下发设备指令过程中，厂商指令完成，回调结果通知
type AlibabaailabsaligenieiotdevicecontrolresultAPIRequest struct {
	model.Params
	// 请求token
	_requestToken string
	// 设备ID
	_deviceId string
	// 厂商执行返回ackCode
	_ackCode string
	// 操作类型 1：控制操作 0:查询
	_type int64
	// 控制成功
	_control bool
}

// NewAlibabaailabsaligenieiotdevicecontrolresultRequest 初始化AlibabaailabsaligenieiotdevicecontrolresultAPIRequest对象
func NewAlibabaailabsaligenieiotdevicecontrolresultRequest() *AlibabaailabsaligenieiotdevicecontrolresultAPIRequest {
	return &AlibabaailabsaligenieiotdevicecontrolresultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsaligenieiotdevicecontrolresultAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.iot.device.control.result"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsaligenieiotdevicecontrolresultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsaligenieiotdevicecontrolresultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestToken is RequestToken Setter
// 请求token
func (r *AlibabaailabsaligenieiotdevicecontrolresultAPIRequest) SetRequestToken(_requestToken string) error {
	r._requestToken = _requestToken
	r.Set("request_token", _requestToken)
	return nil
}

// GetRequestToken RequestToken Getter
func (r AlibabaailabsaligenieiotdevicecontrolresultAPIRequest) GetRequestToken() string {
	return r._requestToken
}

// SetDeviceId is DeviceId Setter
// 设备ID
func (r *AlibabaailabsaligenieiotdevicecontrolresultAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaailabsaligenieiotdevicecontrolresultAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetAckCode is AckCode Setter
// 厂商执行返回ackCode
func (r *AlibabaailabsaligenieiotdevicecontrolresultAPIRequest) SetAckCode(_ackCode string) error {
	r._ackCode = _ackCode
	r.Set("ack_code", _ackCode)
	return nil
}

// GetAckCode AckCode Getter
func (r AlibabaailabsaligenieiotdevicecontrolresultAPIRequest) GetAckCode() string {
	return r._ackCode
}

// SetType is Type Setter
// 操作类型 1：控制操作 0:查询
func (r *AlibabaailabsaligenieiotdevicecontrolresultAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaailabsaligenieiotdevicecontrolresultAPIRequest) GetType() int64 {
	return r._type
}

// SetControl is Control Setter
// 控制成功
func (r *AlibabaailabsaligenieiotdevicecontrolresultAPIRequest) SetControl(_control bool) error {
	r._control = _control
	r.Set("control", _control)
	return nil
}

// GetControl Control Getter
func (r AlibabaailabsaligenieiotdevicecontrolresultAPIRequest) GetControl() bool {
	return r._control
}
