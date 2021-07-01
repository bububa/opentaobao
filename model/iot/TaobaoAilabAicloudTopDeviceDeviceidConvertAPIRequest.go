package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest
开放设备id转换内部设备id API请求
taobao.ailab.aicloud.top.device.deviceid.convert

将开放设备id转换为内部设备id */
type TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest struct {
	model.Params
	// 设备openId
	_deviceOpenId string
	// 技能id
	_skillId string
}

// New
