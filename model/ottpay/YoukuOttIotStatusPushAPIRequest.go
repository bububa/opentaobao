package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttIotStatusPushAPIRequest
iot设备状态变化通知接口 API请求
youku.ott.iot.status.push

ott iot设备状态通知 */
type YoukuOttIotStatusPushAPIRequest struct {
	model.Params
	// 变更信息
	_changeInfo string
}

// New
