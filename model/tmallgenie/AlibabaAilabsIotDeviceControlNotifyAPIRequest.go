package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsIotDeviceControlNotifyAPIRequest
天猫精灵IoT异步控制回调接口 API请求
alibaba.ailabs.iot.device.control.notify

用于天猫精灵IoT云云接入控制结果的异步回调 */
type AlibabaAilabsIotDeviceControlNotifyAPIRequest struct {
	model.Params
	// 入参
	_notifyControlParams *NotifyVehicleControlParams
}

// NewAlibabaAilabsIotDeviceControlNotifyRequest 初始化AlibabaAilabsIotDeviceControlNotifyAPIRequest对象
func NewAlibabaAilabsIotDeviceControlNotifyRequest() *AlibabaAilabsIotDeviceControlNotifyAPIRequest {
	return &AlibabaAilabsIotDeviceControlNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotDeviceControlNotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.device.control.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotDeviceControlNotifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is NotifyControlParams Setter
// 入参
func (r *AlibabaAilabsIotDeviceControlNotifyAPIRequest) SetNotifyControlParams(_notifyControlParams *NotifyVehicleControlParams) error {
	r._notifyControlParams = _notifyControlParams
	r.Set("notify_control_params", _notifyControlParams)
	return nil
}

// Get NotifyControlParams Getter
func (r AlibabaAilabsIotDeviceControlNotifyAPIRequest) GetNotifyControlParams() *NotifyVehicleControlParams {
	return r._notifyControlParams
}
