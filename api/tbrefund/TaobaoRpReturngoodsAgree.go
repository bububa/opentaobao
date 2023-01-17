package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRpReturngoodsAgree 卖家同意退货
// taobao.rp.returngoods.agree
//
// 卖家同意退货，支持淘宝和天猫的订单。
func TaobaoRpReturngoodsAgree(clt *core.SDKClient, req *tbrefund.TaobaoRpReturngoodsAgreeAPIRequest, session string) (*tbrefund.TaobaoRpReturngoodsAgreeAPIResponse, error) {
	var resp tbrefund.TaobaoRpReturngoodsAgreeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
