package aedata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateOrderGetAPIRequest
AE流量订单详情获取API API请求
aliexpress.affiliate.order.get

联盟推广订单效果获取API */
type AliexpressAffiliateOrderGetAPIRequest struct {
	model.Params
	// 安全签名
	_appSignature string
	// 返回的字段列表
	_fields string
	// 订单ID列表，以逗号分隔，当前只支持子订单ID查询
	_orderIds string
}

// New
