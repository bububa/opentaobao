package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceControlCustomAPIRequest
设备控制自定义扩展接口 API请求
taobao.ailab.aicloud.top.device.control.custom

设备控制自定义扩展接口 */
type TaobaoAilabAicloudTopDeviceControlCustomAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
	// 参数key-value列表
	_param2 []string
}

// New
