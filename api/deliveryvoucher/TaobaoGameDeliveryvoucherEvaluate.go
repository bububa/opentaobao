package deliveryvoucher

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// TaobaoGameDeliveryvoucherEvaluate 卡券评价回传
// taobao.game.deliveryvoucher.evaluate
//
// 卡券ISV回传商品评价
func TaobaoGameDeliveryvoucherEvaluate(clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherEvaluateAPIRequest, resp *deliveryvoucher.TaobaoGameDeliveryvoucherEvaluateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
