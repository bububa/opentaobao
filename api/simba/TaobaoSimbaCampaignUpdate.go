package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignUpdate 更新一个推广计划
// taobao.simba.campaign.update
//
// 更新一个推广计划，可以设置推广计划名字，修改推广计划上下线状态。
func TaobaoSimbaCampaignUpdate(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCampaignUpdateAPIRequest, resp *simba.TaobaoSimbaCampaignUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
