package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServiceSettleadjustmentGet 查询结算调整单单条记录
// tmall.service.settleadjustment.get
//
// 提供给服务商通过结算调整单id获取结算调整单信息
func TmallServiceSettleadjustmentGet(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServiceSettleadjustmentGetAPIRequest, resp *tmallservice.TmallServiceSettleadjustmentGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
