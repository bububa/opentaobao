package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkTradeDiscountBillGetAPIRequest
订单优惠账单查询 API请求
alibaba.wdk.trade.discount.bill.get

商家查询订单优惠账单 */
type AlibabaWdkTradeDiscountBillGetAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *OrderDiscountBillQueryRequest
}

// New
