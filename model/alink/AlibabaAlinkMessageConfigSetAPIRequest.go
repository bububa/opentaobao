package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlinkMessageConfigSetAPIRequest
消息提醒开关 API请求
alibaba.alink.message.config.set

阿里智能消息开关 */
type AlibabaAlinkMessageConfigSetAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
	// 0：开启，1：关闭
	_pushDisabled string
}

// New
