package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitripaxintransfundconfirm 确认资金单
// taobao.alitrip.axin.trans.fund.confirm
//
// 通过外部订单号进行资金结算
func Taobaoalitripaxintransfundconfirm(clt *core.SDKClient, req *axintrade.TaobaoalitripaxintransfundconfirmAPIRequest, session string) (*axintrade.TaobaoalitripaxintransfundconfirmAPIResponse, error) {
	var resp axintrade.TaobaoalitripaxintransfundconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
