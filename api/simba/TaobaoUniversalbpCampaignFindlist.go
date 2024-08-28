package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCampaignFindlist 查询全量计划列表(不分页)
// taobao.universalbp.campaign.findlist
//
// 查询场景内的全量计划列表
func TaobaoUniversalbpCampaignFindlist(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpCampaignFindlistAPIRequest, resp *simba.TaobaoUniversalbpCampaignFindlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
