package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitripaxintranspayregisterreapply 阿信支付入驻重新申请
// taobao.alitrip.axin.trans.pay.register.reapply
//
// 阿信支付入驻重新申请
// 用于支付平台驳回，商户提交时的业务场景
func Taobaoalitripaxintranspayregisterreapply(clt *core.SDKClient, req *axintrade.TaobaoalitripaxintranspayregisterreapplyAPIRequest, session string) (*axintrade.TaobaoalitripaxintranspayregisterreapplyAPIResponse, error) {
	var resp axintrade.TaobaoalitripaxintranspayregisterreapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
