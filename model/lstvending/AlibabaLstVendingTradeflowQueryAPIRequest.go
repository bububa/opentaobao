package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstVendingTradeflowQueryAPIRequest
自动售卖机交易流水查询 API请求
alibaba.lst.vending.tradeflow.query

零售通自动售卖机交易流水查询接口，品牌商通过此接口同步商品交易数据。 */
type AlibabaLstVendingTradeflowQueryAPIRequest struct {
	model.Params
	// 交易流水查询条件
	_openTradeFlowQuery *OpenTradeFlowQuery
}

// New
