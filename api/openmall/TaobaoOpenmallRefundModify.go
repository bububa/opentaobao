package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundModify 修改OpenMall退款申请
// taobao.openmall.refund.modify
//
// 修改OpenMall退款申请
func TaobaoOpenmallRefundModify(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundModifyAPIRequest, resp *openmall.TaobaoOpenmallRefundModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
