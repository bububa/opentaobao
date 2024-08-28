package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRtrptCampaignGet 获取推广计划实时报表数据
// taobao.simba.rtrpt.campaign.get
//
// 获取推广计划实时报表数据
func TaobaoSimbaRtrptCampaignGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRtrptCampaignGetAPIRequest, resp *simba.TaobaoSimbaRtrptCampaignGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
