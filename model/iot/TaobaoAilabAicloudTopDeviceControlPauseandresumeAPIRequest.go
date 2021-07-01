package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest
设备播放暂停 API请求
taobao.ailab.aicloud.top.device.control.pauseandresume

设备播放暂停 */
type TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
	// 是暂停还是继续
	_param2 bool
}

// New
