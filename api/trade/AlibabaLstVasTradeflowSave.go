package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaLstVasTradeflowSave 交易信息回流
// alibaba.lst.vas.tradeflow.save
//
// 自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
func AlibabaLstVasTradeflowSave(clt *core.SDKClient, req *trade.AlibabaLstVasTradeflowSaveAPIRequest, resp *trade.AlibabaLstVasTradeflowSaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
