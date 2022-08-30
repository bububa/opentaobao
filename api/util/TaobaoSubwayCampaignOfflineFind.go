package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoSubwayCampaignOfflineFind 查询某计划离线多日汇总列表
// taobao.subway.campaign.offline.find
//
// 查询某计划离线列表
func TaobaoSubwayCampaignOfflineFind(clt *core.SDKClient, req *util.TaobaoSubwayCampaignOfflineFindAPIRequest, session string) (*util.TaobaoSubwayCampaignOfflineFindAPIResponse, error) {
	var resp util.TaobaoSubwayCampaignOfflineFindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
