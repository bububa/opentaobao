package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressTradeDsOrderGetAPIRequest
买家查询订单详情 API请求
aliexpress.trade.ds.order.get

买家查询订单详情，用于dropshipper */
type AliexpressTradeDsOrderGetAPIRequest struct {
	model.Params
	// 订单查询条件
	_singleOrderQuery *AeopSingleOrderQuery
}

// New
