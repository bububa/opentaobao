package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TmallServiceSettleadjustmentModify 修改结算调整单
// tmall.service.settleadjustment.modify
//
// 提供给服务商在对结算有异议时，发起结算调整单。
// 通过说明调整单ID，调整费用值，调整原因进行结算调整单修改。
func TmallServiceSettleadjustmentModify(ctx context.Context, clt *core.SDKClient, req *user.TmallServiceSettleadjustmentModifyAPIRequest, resp *user.TmallServiceSettleadjustmentModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
