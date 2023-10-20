package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitripaxintranspayregisteraudit 阿信支付入驻审核通知
// taobao.alitrip.axin.trans.pay.register.audit
//
// 阿信支付入驻审核通知
func Taobaoalitripaxintranspayregisteraudit(clt *core.SDKClient, req *axintrade.TaobaoalitripaxintranspayregisterauditAPIRequest, session string) (*axintrade.TaobaoalitripaxintranspayregisterauditAPIResponse, error) {
	var resp axintrade.TaobaoalitripaxintranspayregisterauditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
