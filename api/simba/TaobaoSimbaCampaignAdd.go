package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignAdd 创建一个推广计划
// taobao.simba.campaign.add
//
// 创建一个推广计划
func TaobaoSimbaCampaignAdd(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCampaignAddAPIRequest, resp *simba.TaobaoSimbaCampaignAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
