package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceGetstatusAPIRequest
获取设备状态 API请求
taobao.ailab.aicloud.top.device.getstatus

获取设备状态 */
type TaobaoAilabAicloudTopDeviceGetstatusAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
}

// New
