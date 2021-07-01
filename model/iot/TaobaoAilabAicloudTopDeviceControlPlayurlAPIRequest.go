package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest
点播url API请求
taobao.ailab.aicloud.top.device.control.playurl

点播url */
type TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
	// url
	_param2 string
}

// New
