package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest
获取设备详细信息 API请求
taobao.ailab.aicloud.top.device.detailinfo.get

获取设备详细信息 */
type TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest struct {
	model.Params
	// 三方用户id或淘宝openId
	_originUserId string
	// 账号秘钥
	_schemaKey string
	// 三方传extUser，淘宝传openTaoBaoUser
	_userType string
	// 设备id
	_deviceId string
}

// New
