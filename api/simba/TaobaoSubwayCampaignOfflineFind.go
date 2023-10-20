package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayCampaignOfflineFind 查询某计划离线多日汇总列表
// taobao.subway.campaign.offline.find
//
// 查询某计划离线列表
func TaobaoSubwayCampaignOfflineFind(clt *core.SDKClient, req *simba.TaobaoSubwayCampaignOfflineFindAPIRequest, resp *simba.TaobaoSubwayCampaignOfflineFindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
