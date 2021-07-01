package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest
设备儿童锁 API请求
taobao.ailab.aicloud.top.device.control.childlock

设备儿童锁 */
type TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
	// 是否打开
	_param2 bool
}

// New
