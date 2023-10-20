package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundMessageSubmit 提交退款单留言
// taobao.openmall.refund.message.submit
//
// OpenMall业务提交退款单留言
func TaobaoOpenmallRefundMessageSubmit(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundMessageSubmitAPIRequest, resp *openmall.TaobaoOpenmallRefundMessageSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
