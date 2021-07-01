package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkChannelOrderUserrefundAPIRequest
用户发起售后退款(整单/部分) API请求
alibaba.wdk.channel.order.userrefund

用户发起售后退款(整单/部分) */
type AlibabaWdkChannelOrderUserrefundAPIRequest struct {
	model.Params
	// 退款信息
	_orderUserRefundInfo *OrderUserRefundInfo
}

// New
