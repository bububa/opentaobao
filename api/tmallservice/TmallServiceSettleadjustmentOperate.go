package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServiceSettleadjustmentOperate 天猫服务调整单操作
// tmall.service.settleadjustment.operate
//
// 提供给服务商在对结算有异议时，发起结算调整单。
// 通过说明调整单ID，调整费用值，调整原因进行结算调整单修改。
func TmallServiceSettleadjustmentOperate(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServiceSettleadjustmentOperateAPIRequest, resp *tmallservice.TmallServiceSettleadjustmentOperateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
