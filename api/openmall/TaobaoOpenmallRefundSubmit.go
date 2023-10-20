package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundSubmit 提交OpenMall退款单物流
// taobao.openmall.refund.submit
//
// 提交OpenMall退款单物流
func TaobaoOpenmallRefundSubmit(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundSubmitAPIRequest, resp *openmall.TaobaoOpenmallRefundSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
