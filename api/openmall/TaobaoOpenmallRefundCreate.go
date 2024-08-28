package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundCreate 创建OpenMall退款单
// taobao.openmall.refund.create
//
// 创建OpenMall退款单
// 如存在未完结的退款单，则返回该退款单ID
func TaobaoOpenmallRefundCreate(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundCreateAPIRequest, resp *openmall.TaobaoOpenmallRefundCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
