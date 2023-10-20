package elife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/elife"
)

// TaobaoElifeLifecardConsume 品牌惠卡券核销
// taobao.elife.lifecard.consume
//
// 用户线上购买生活汇品牌惠虚拟消费卡，线下购物时，商家码枪核销，涉及用户虚拟卡余额扣减操作
func TaobaoElifeLifecardConsume(clt *core.SDKClient, req *elife.TaobaoElifeLifecardConsumeAPIRequest, resp *elife.TaobaoElifeLifecardConsumeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
