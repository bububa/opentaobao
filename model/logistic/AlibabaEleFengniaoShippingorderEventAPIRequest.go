package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoShippingorderEventAPIRequest
查询运单事件信息 API请求
alibaba.ele.fengniao.shippingorder.event

查询运单事件信息 */
type AlibabaEleFengniaoShippingorderEventAPIRequest struct {
	model.Params
	// appid
	_appId string
	// 外部订单号
	_partnerOrderCode string
}

// New
