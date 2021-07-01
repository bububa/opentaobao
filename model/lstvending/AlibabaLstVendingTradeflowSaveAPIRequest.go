package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstVendingTradeflowSaveAPIRequest
自动售卖机交易信息回流 API请求
alibaba.lst.vending.tradeflow.save

自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。 */
type AlibabaLstVendingTradeflowSaveAPIRequest struct {
	model.Params
	// 交易流水信息
	_tradeFlowDTOList []VendingTradeFlowDto
}

// New
