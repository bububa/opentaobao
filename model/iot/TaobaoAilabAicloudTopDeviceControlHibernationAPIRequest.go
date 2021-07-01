package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest
定时休眠 API请求
taobao.ailab.aicloud.top.device.control.hibernation

定时休眠 */
type TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest struct {
	model.Params
	// 用户信息
	_param0 *OpenBaseInfo
	// 设备id
	_param1 string
	// N秒后休眠
	_param2 string
}

// New
