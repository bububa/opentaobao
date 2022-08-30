package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayCampaignOfflineLayeredfind 查询计划离线列表30天转化周期
// taobao.subway.campaign.offline.layeredfind
//
// 查询某计划离线列表
func TaobaoSubwayCampaignOfflineLayeredfind(clt *core.SDKClient, req *simba.TaobaoSubwayCampaignOfflineLayeredfindAPIRequest, session string) (*simba.TaobaoSubwayCampaignOfflineLayeredfindAPIResponse, error) {
	var resp simba.TaobaoSubwayCampaignOfflineLayeredfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
