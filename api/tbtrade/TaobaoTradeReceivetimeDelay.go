package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTradeReceivetimeDelay 延长交易收货时间
// taobao.trade.receivetime.delay
//
// 延长交易收货时间
func TaobaoTradeReceivetimeDelay(clt *core.SDKClient, req *tbtrade.TaobaoTradeReceivetimeDelayAPIRequest, session string) (*tbtrade.TaobaoTradeReceivetimeDelayAPIResponse, error) {
	var resp tbtrade.TaobaoTradeReceivetimeDelayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
