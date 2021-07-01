package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkChannelOrderCreateAPIRequest
创建订单 API请求
alibaba.wdk.channel.order.create

外部商家创建订单 */
type AlibabaWdkChannelOrderCreateAPIRequest struct {
	model.Params
	// 订单信息
	_orderInfo *OrderInfo
}

// New
