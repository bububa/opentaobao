package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgCpaActivityDetail 淘宝客-推广者-CPA活动执行明细
// taobao.tbk.dg.cpa.activity.detail
//
// 淘宝客获取CPA活动具体执行效果的明细数据（含预估和结算维度）
func TaobaoTbkDgCpaActivityDetail(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkDgCpaActivityDetailAPIRequest, resp *tbk.TaobaoTbkDgCpaActivityDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
