package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFlowAlipayPublishAPIRequest
流量钱包流量发放-面向支付宝用户 API请求
alibaba.aliqin.flow.alipay.publish

用户淘宝流量钱包商家给支付宝用户发放流量 */
type AlibabaAliqinFlowAlipayPublishAPIRequest struct {
	model.Params
	// 用户的支付宝ID
	_alipayId string
	// 外部流水号，用来做幂等校验
	_serial string
	// 发放的流量数，单位MB
	_flow string
	// 发放原因
	_reason string
}

// New
