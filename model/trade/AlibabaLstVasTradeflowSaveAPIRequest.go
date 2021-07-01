package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstVasTradeflowSaveAPIRequest
交易信息回流 API请求
alibaba.lst.vas.tradeflow.save

自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。 */
type AlibabaLstVasTradeflowSaveAPIRequest struct {
	model.Params
	// 交易流水信息
	_tradeFlowModelList *TradeFlowModel
}

// New
