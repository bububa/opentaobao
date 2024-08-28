package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignPlatformGet 取得一个推广计划的投放平台设置
// taobao.simba.campaign.platform.get
//
// 获得一个推广计划的投放平台设置
func TaobaoSimbaCampaignPlatformGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCampaignPlatformGetAPIRequest, resp *simba.TaobaoSimbaCampaignPlatformGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
