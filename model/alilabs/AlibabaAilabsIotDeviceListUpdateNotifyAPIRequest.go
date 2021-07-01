package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest
设备列表更新通知 API请求
alibaba.ailabs.iot.device.list.update.notify

用于人工智能实验室IoT合作厂商上报三方接入IoT设备列表更新时的通知 */
type AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest struct {
	model.Params
	// 用户OAuth授权后的token
	_token string
	// 厂商在天猫精灵开放平台申请的技能id
	_skillId string
	// 更新类型 1：设备更新 2：场景更新
	_type string
}

// New
