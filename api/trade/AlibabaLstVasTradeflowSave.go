package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Alibabalstvastradeflowsave 交易信息回流
// alibaba.lst.vas.tradeflow.save
//
// 自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
func Alibabalstvastradeflowsave(clt *core.SDKClient, req *trade.AlibabalstvastradeflowsaveAPIRequest, session string) (*trade.AlibabalstvastradeflowsaveAPIResponse, error) {
	var resp trade.AlibabalstvastradeflowsaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
