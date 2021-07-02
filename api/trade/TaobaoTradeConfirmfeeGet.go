package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoTradeConfirmfeeGet 获取交易确认收货费用
// taobao.trade.confirmfee.get
//
// 获取交易确认收货费用，可以获取主订单或子订单的确认收货费用
func TaobaoTradeConfirmfeeGet(clt *core.SDKClient, req *trade.TaobaoTradeConfirmfeeGetAPIRequest, session string) (*trade.TaobaoTradeConfirmfeeGetAPIResponse, error) {
	var resp trade.TaobaoTradeConfirmfeeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
