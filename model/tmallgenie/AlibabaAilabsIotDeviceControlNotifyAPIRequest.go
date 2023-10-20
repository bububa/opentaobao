package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotDeviceControlNotifyAPIRequest 天猫精灵IoT异步控制回调接口 API请求
// alibaba.ailabs.iot.device.control.notify
//
// 用于天猫精灵IoT云云接入控制结果的异步回调
type AlibabaAilabsIotDeviceControlNotifyAPIRequest struct {
	model.Params
	// 入参
	_notifyControlParams *NotifyVehicleControlParams
}

// NewAlibabaAilabsIotDeviceControlNotifyRequest 初始化AlibabaAilabsIotDeviceControlNotifyAPIRequest对象
func NewAlibabaAilabsIotDeviceControlNotifyRequest() *AlibabaAilabsIotDeviceControlNotifyAPIRequest {
	return &AlibabaAilabsIotDeviceControlNotifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsIotDeviceControlNotifyAPIRequest) Reset() {
	r._notifyControlParams = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotDeviceControlNotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.device.control.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotDeviceControlNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsIotDeviceControlNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNotifyControlParams is NotifyControlParams Setter
// 入参
func (r *AlibabaAilabsIotDeviceControlNotifyAPIRequest) SetNotifyControlParams(_notifyControlParams *NotifyVehicleControlParams) error {
	r._notifyControlParams = _notifyControlParams
	r.Set("notify_control_params", _notifyControlParams)
	return nil
}

// GetNotifyControlParams NotifyControlParams Getter
func (r AlibabaAilabsIotDeviceControlNotifyAPIRequest) GetNotifyControlParams() *NotifyVehicleControlParams {
	return r._notifyControlParams
}

var poolAlibabaAilabsIotDeviceControlNotifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsIotDeviceControlNotifyRequest()
	},
}

// GetAlibabaAilabsIotDeviceControlNotifyRequest 从 sync.Pool 获取 AlibabaAilabsIotDeviceControlNotifyAPIRequest
func GetAlibabaAilabsIotDeviceControlNotifyAPIRequest() *AlibabaAilabsIotDeviceControlNotifyAPIRequest {
	return poolAlibabaAilabsIotDeviceControlNotifyAPIRequest.Get().(*AlibabaAilabsIotDeviceControlNotifyAPIRequest)
}

// ReleaseAlibabaAilabsIotDeviceControlNotifyAPIRequest 将 AlibabaAilabsIotDeviceControlNotifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsIotDeviceControlNotifyAPIRequest(v *AlibabaAilabsIotDeviceControlNotifyAPIRequest) {
	v.Reset()
	poolAlibabaAilabsIotDeviceControlNotifyAPIRequest.Put(v)
}
