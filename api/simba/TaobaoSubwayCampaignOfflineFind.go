package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwaycampaignofflinefind 查询某计划离线多日汇总列表
// taobao.subway.campaign.offline.find
//
// 查询某计划离线列表
func Taobaosubwaycampaignofflinefind(clt *core.SDKClient, req *simba.TaobaosubwaycampaignofflinefindAPIRequest, session string) (*simba.TaobaosubwaycampaignofflinefindAPIResponse, error) {
	var resp simba.TaobaosubwaycampaignofflinefindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
