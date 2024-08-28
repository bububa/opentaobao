package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundClose 关闭OpenMall退款单
// taobao.openmall.refund.close
//
// 关闭OpenMall退款单
func TaobaoOpenmallRefundClose(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundCloseAPIRequest, resp *openmall.TaobaoOpenmallRefundCloseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
