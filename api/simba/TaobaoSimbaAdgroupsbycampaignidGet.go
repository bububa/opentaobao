package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupsbycampaignidGet 批量得到推广计划下的推广单元
// taobao.simba.adgroupsbycampaignid.get
//
// 根据推广计划ID分页获取推广计划下的推广单元信息
func TaobaoSimbaAdgroupsbycampaignidGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupsbycampaignidGetAPIRequest, resp *simba.TaobaoSimbaAdgroupsbycampaignidGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
