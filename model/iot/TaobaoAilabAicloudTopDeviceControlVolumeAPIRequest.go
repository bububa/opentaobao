package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest
设备音量 API请求
taobao.ailab.aicloud.top.device.control.volume

设备音量 */
type TaobaoAilabAicloudTopDeviceControlVolumeAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
	// 音量0-100
	_param2 int64
}

// New
