package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceControlLampAPIRequest
台灯控制 API请求
taobao.ailab.aicloud.top.device.control.lamp

台灯控制 */
type TaobaoAilabAicloudTopDeviceControlLampAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
	// 是否打开
	_param2 bool
	// 目标名称
	_param3 string
}

// New
