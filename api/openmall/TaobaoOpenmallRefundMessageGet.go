package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundMessageGet openmall获取退款单留言
// taobao.openmall.refund.message.get
//
// openmall获取退款单留言
func TaobaoOpenmallRefundMessageGet(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundMessageGetAPIRequest, resp *openmall.TaobaoOpenmallRefundMessageGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
