package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpcampaignget 查询单个计划详情
// taobao.universalbp.campaign.get
//
// 查询单个计划详情信息（不包括报表数据）
func Taobaouniversalbpcampaignget(clt *core.SDKClient, req *simba.TaobaouniversalbpcampaigngetAPIRequest, session string) (*simba.TaobaouniversalbpcampaigngetAPIResponse, error) {
	var resp simba.TaobaouniversalbpcampaigngetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
