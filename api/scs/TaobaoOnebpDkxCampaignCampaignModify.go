package scs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxCampaignCampaignModify 修改计划
// taobao.onebp.dkx.campaign.campaign.modify
//
// 修改计划。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
func TaobaoOnebpDkxCampaignCampaignModify(ctx context.Context, clt *core.SDKClient, req *scs.TaobaoOnebpDkxCampaignCampaignModifyAPIRequest, resp *scs.TaobaoOnebpDkxCampaignCampaignModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
