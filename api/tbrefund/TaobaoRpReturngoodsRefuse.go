package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRpReturngoodsRefuse 卖家拒绝退货
// taobao.rp.returngoods.refuse
//
// 卖家拒绝退货，目前仅支持天猫退货。
func TaobaoRpReturngoodsRefuse(clt *core.SDKClient, req *tbrefund.TaobaoRpReturngoodsRefuseAPIRequest, resp *tbrefund.TaobaoRpReturngoodsRefuseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
