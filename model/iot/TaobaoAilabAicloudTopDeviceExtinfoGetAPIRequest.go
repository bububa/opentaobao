package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest
获取设备扩展信息 API请求
taobao.ailab.aicloud.top.device.extinfo.get

获取设备扩展信息 */
type TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest struct {
	model.Params
	// 三方id、淘宝openId
	_originUserId string
	// 账号秘钥
	_schemaKey string
	// 类型：openTaoBao, extUser
	_userType string
	// 设备id
	_deviceId string
}

// New
