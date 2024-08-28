package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundModify 修改OpenMall退款申请
// taobao.openmall.refund.modify
//
// 修改OpenMall退款申请
func TaobaoOpenmallRefundModify(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundModifyAPIRequest, resp *openmall.TaobaoOpenmallRefundModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
