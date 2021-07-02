package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundModify 修改OpenMall退款申请
// taobao.openmall.refund.modify
//
// 修改OpenMall退款申请
func TaobaoOpenmallRefundModify(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundModifyAPIRequest, session string) (*openmall.TaobaoOpenmallRefundModifyAPIResponse, error) {
	var resp openmall.TaobaoOpenmallRefundModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
