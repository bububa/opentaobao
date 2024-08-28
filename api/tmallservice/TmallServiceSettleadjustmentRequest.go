package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServiceSettleadjustmentRequest 创建结算调整单
// tmall.service.settleadjustment.request
//
// 提供给服务商在对结算有异议时，发起结算调整单。
// 通过说明工单ID，调整费用值，调整原因进行新建结算调整单。
func TmallServiceSettleadjustmentRequest(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServiceSettleadjustmentRequestAPIRequest, resp *tmallservice.TmallServiceSettleadjustmentRequestAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
