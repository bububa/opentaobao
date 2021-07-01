package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkOrderFinanceBillQueryAPIRequest
资金合规商家账单 API请求
alibaba.wdk.order.finance.bill.query

拉取资金合规商家账单 */
type AlibabaWdkOrderFinanceBillQueryAPIRequest struct {
	model.Params
	// 入参
	_billQueryRequest *WdkOpenOrderFinanceBillQueryRequest
}

// New
