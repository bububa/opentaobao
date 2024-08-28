package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayCampaignOfflineLayeredfind 查询计划离线列表30天转化周期
// taobao.subway.campaign.offline.layeredfind
//
// 查询某计划离线列表
func TaobaoSubwayCampaignOfflineLayeredfind(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayCampaignOfflineLayeredfindAPIRequest, resp *simba.TaobaoSubwayCampaignOfflineLayeredfindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
