package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoOrderPushAPIRequest
推送订单 API请求
alibaba.ele.fengniao.order.push

推送淘宝订单至蜂鸟开放平台配送 */
type AlibabaEleFengniaoOrderPushAPIRequest struct {
	model.Params
	// 参数param
	_param *Param
}

// New
