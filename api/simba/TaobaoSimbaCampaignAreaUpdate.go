package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignAreaUpdate 更新一个推广计划的投放地域
// taobao.simba.campaign.area.update
//
// 更新一个推广计划的投放地域
func TaobaoSimbaCampaignAreaUpdate(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCampaignAreaUpdateAPIRequest, resp *simba.TaobaoSimbaCampaignAreaUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
