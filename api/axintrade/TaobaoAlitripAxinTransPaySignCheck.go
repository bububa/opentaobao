package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitripaxintranspaysigncheck 阿信支付宝验签服务
// taobao.alitrip.axin.trans.pay.sign.check
//
// 阿信支付宝验签服务
func Taobaoalitripaxintranspaysigncheck(clt *core.SDKClient, req *axintrade.TaobaoalitripaxintranspaysigncheckAPIRequest, session string) (*axintrade.TaobaoalitripaxintranspaysigncheckAPIResponse, error) {
	var resp axintrade.TaobaoalitripaxintranspaysigncheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
