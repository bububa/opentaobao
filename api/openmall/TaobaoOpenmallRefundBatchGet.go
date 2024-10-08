package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundBatchGet 批量获取openmall退款单
// taobao.openmall.refund.batch.get
//
// 批量获取openmall退款单
// 注意：该接口信息存在延迟，如需实时详情请访问taobao.openmall.refund.get
func TaobaoOpenmallRefundBatchGet(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundBatchGetAPIRequest, resp *openmall.TaobaoOpenmallRefundBatchGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
