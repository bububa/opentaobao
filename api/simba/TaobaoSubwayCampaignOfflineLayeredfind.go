package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwaycampaignofflinelayeredfind 查询计划离线列表30天转化周期
// taobao.subway.campaign.offline.layeredfind
//
// 查询某计划离线列表
func Taobaosubwaycampaignofflinelayeredfind(clt *core.SDKClient, req *simba.TaobaosubwaycampaignofflinelayeredfindAPIRequest, session string) (*simba.TaobaosubwaycampaignofflinelayeredfindAPIResponse, error) {
	var resp simba.TaobaosubwaycampaignofflinelayeredfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
