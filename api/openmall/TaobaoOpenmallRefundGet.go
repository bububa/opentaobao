package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundGet 获取OpenMall退款单详情
// taobao.openmall.refund.get
//
// 获取OpenMall退款单详情
func TaobaoOpenmallRefundGet(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundGetAPIRequest, resp *openmall.TaobaoOpenmallRefundGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
