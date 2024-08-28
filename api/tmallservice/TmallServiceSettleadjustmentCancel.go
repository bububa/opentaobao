package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServiceSettleadjustmentCancel 取消结算调整单
// tmall.service.settleadjustment.cancel
//
// 提供给服务商在对取消已经发起的结算调整单。
// 通过说明调整单ID进行结算调整单取消。
func TmallServiceSettleadjustmentCancel(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServiceSettleadjustmentCancelAPIRequest, resp *tmallservice.TmallServiceSettleadjustmentCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
