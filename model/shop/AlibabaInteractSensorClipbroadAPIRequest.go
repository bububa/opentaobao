package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorClipbroadAPIRequest
Weex页面设置或读取剪切板 API请求
alibaba.interact.sensor.clipbroad

Weex页面设置或读取剪切板 */
type AlibabaInteractSensorClipbroadAPIRequest struct {
	model.Params
	// 客户端鉴权使用，实际不会发送或接收数据
	_unNamed string
}

// New
