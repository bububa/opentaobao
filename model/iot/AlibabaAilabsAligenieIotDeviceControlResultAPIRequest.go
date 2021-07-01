package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieIotDeviceControlResultAPIRequest
设备控制结果 API请求
alibaba.ailabs.aligenie.iot.device.control.result

智能IOT解决外部厂商在云云模式在用户通过天猫精灵下发设备指令过程中，厂商指令完成，回调结果通知 */
type AlibabaAilabsAligenieIotDeviceControlResultAPIRequest struct {
	model.Params
	// 请求token
	_requestToken string
	// 设备ID
	_deviceId string
	// 操作类型 1：控制操作 0:查询
	_type int64
	// 控制成功
	_control bool
	// 厂商执行返回ackCode
	_ackCode string
}

// New
