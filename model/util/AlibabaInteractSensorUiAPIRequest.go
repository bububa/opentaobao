package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorUiAPIRequest
基本ui操作 API请求
alibaba.interact.sensor.ui

Weex 基本UI操作 */
type AlibabaInteractSensorUiAPIRequest struct {
	model.Params
	// 仅作客户端鉴权使用，不会发送接收请求
	_unNamed string
}

// New
