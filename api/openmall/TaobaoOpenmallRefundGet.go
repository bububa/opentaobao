package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundGet 获取OpenMall退款单详情
// taobao.openmall.refund.get
//
// 获取OpenMall退款单详情
func TaobaoOpenmallRefundGet(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundGetAPIRequest, resp *openmall.TaobaoOpenmallRefundGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
