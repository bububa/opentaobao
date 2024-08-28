package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallServiceSettleadjustmentSearch 服务商15分钟获取数据
// tmall.service.settleadjustment.search
//
// 天猫服务平台，按修改时间，时间间隔在15中内（包含15分钟），获取调整单数据
func TmallServiceSettleadjustmentSearch(ctx context.Context, clt *core.SDKClient, req *servicecenter.TmallServiceSettleadjustmentSearchAPIRequest, resp *servicecenter.TmallServiceSettleadjustmentSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
