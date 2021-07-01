package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFinanceOrderBackflowAPIRequest
财务订单回流 API请求
alibaba.wdk.finance.order.backflow

星巴克拉取财务订单回流数据 */
type AlibabaWdkFinanceOrderBackflowAPIRequest struct {
	model.Params
	// 财务订单回流入参
	_financeOrderDetailRequest *FinanceOrderDetailRequest
}

// New
