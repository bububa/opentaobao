package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignAreaGet 取得一个推广计划的投放地域设置
// taobao.simba.campaign.area.get
//
// 取得一个推广计划的投放地域设置
func TaobaoSimbaCampaignAreaGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCampaignAreaGetAPIRequest, resp *simba.TaobaoSimbaCampaignAreaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
