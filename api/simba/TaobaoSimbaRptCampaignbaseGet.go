package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptCampaignbaseGet 推广计划报表基础数据对象
// taobao.simba.rpt.campaignbase.get
//
// 推广计划报表基础数据对象
func TaobaoSimbaRptCampaignbaseGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRptCampaignbaseGetAPIRequest, resp *simba.TaobaoSimbaRptCampaignbaseGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
