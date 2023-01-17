package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRpReturngoodsRefuse 卖家拒绝退货
// taobao.rp.returngoods.refuse
//
// 卖家拒绝退货，目前仅支持天猫退货。
func TaobaoRpReturngoodsRefuse(clt *core.SDKClient, req *tbrefund.TaobaoRpReturngoodsRefuseAPIRequest, session string) (*tbrefund.TaobaoRpReturngoodsRefuseAPIResponse, error) {
	var resp tbrefund.TaobaoRpReturngoodsRefuseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
