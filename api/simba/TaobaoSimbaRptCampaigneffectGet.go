package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptCampaigneffectGet 推广计划效果报表数据对象
// taobao.simba.rpt.campaigneffect.get
//
// 推广计划效果报表数据对象
func TaobaoSimbaRptCampaigneffectGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRptCampaigneffectGetAPIRequest, resp *simba.TaobaoSimbaRptCampaigneffectGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
