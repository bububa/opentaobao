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

// New
