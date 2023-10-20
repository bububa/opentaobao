package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappInteractBenefitItemGet 读取实物权益奖池对应绑定的专属下单商品
// taobao.miniapp.interact.benefit.item.get
//
// 读取实物权益奖池对应绑定的专属下单商品
func TaobaoMiniappInteractBenefitItemGet(clt *core.SDKClient, req *miniappopen.TaobaoMiniappInteractBenefitItemGetAPIRequest, resp *miniappopen.TaobaoMiniappInteractBenefitItemGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
