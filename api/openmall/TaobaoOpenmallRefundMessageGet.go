package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundMessageGet openmall获取退款单留言
// taobao.openmall.refund.message.get
//
// openmall获取退款单留言
func TaobaoOpenmallRefundMessageGet(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundMessageGetAPIRequest, resp *openmall.TaobaoOpenmallRefundMessageGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
