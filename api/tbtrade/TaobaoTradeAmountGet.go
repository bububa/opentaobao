package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// Taobaotradeamountget 交易帐务查询
// taobao.trade.amount.get
//
// 卖家查询该笔交易的资金帐务相关的数据；
// 1. 只供卖家使用，买家不可使用
// 2. 可查询所有的状态的交易，但不同状态时交易的相关数据可能会有不同
func Taobaotradeamountget(clt *core.SDKClient, req *tbtrade.TaobaotradeamountgetAPIRequest, session string) (*tbtrade.TaobaotradeamountgetAPIResponse, error) {
	var resp tbtrade.TaobaotradeamountgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
