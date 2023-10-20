package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsiotdevicecontrolnotifyAPIRequest 天猫精灵IoT异步控制回调接口 API请求
// alibaba.ailabs.iot.device.control.notify
//
// 用于天猫精灵IoT云云接入控制结果的异步回调
type AlibabaailabsiotdevicecontrolnotifyAPIRequest struct {
	model.Params
	// 入参
	_notifyControlParams *NotifyVehicleControlParams
}

// NewAlibabaailabsiotdevicecontrolnotifyRequest 初始化AlibabaailabsiotdevicecontrolnotifyAPIRequest对象
func NewAlibabaailabsiotdevicecontrolnotifyRequest() *AlibabaailabsiotdevicecontrolnotifyAPIRequest {
	return &AlibabaailabsiotdevicecontrolnotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsiotdevicecontrolnotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.device.control.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsiotdevicecontrolnotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsiotdevicecontrolnotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNotifyControlParams is NotifyControlParams Setter
// 入参
func (r *AlibabaailabsiotdevicecontrolnotifyAPIRequest) SetNotifyControlParams(_notifyControlParams *NotifyVehicleControlParams) error {
	r._notifyControlParams = _notifyControlParams
	r.Set("notify_control_params", _notifyControlParams)
	return nil
}

// GetNotifyControlParams NotifyControlParams Getter
func (r AlibabaailabsiotdevicecontrolnotifyAPIRequest) GetNotifyControlParams() *NotifyVehicleControlParams {
	return r._notifyControlParams
}
