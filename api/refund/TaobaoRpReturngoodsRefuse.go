package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

/* TaobaoRpReturngoodsRefuse
卖家拒绝退货
taobao.rp.returngoods.refuse

卖家拒绝退货，目前仅支持天猫退货。 */
func TaobaoRpReturngoodsRefuse(clt *core.SDKClient, req *refund.TaobaoRpReturngoodsRefuseAPIRequest, session string) (*refund.TaobaoRpReturngoodsRefuseAPIResponse, error) {
	var resp refund.TaobaoRpReturngoodsRefuseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
