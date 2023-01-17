package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappInteractBenefitItemGet 读取实物权益奖池对应绑定的专属下单商品
// taobao.miniapp.interact.benefit.item.get
//
// 读取实物权益奖池对应绑定的专属下单商品
func TaobaoMiniappInteractBenefitItemGet(clt *core.SDKClient, req *miniappopen.TaobaoMiniappInteractBenefitItemGetAPIRequest, session string) (*miniappopen.TaobaoMiniappInteractBenefitItemGetAPIResponse, error) {
	var resp miniappopen.TaobaoMiniappInteractBenefitItemGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
