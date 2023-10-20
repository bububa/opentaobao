package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpcampaignfindpage 查询计划分页列表
// taobao.universalbp.campaign.findpage
//
// 分页查询场景内的计划列表
func Taobaouniversalbpcampaignfindpage(clt *core.SDKClient, req *simba.TaobaouniversalbpcampaignfindpageAPIRequest, session string) (*simba.TaobaouniversalbpcampaignfindpageAPIResponse, error) {
	var resp simba.TaobaouniversalbpcampaignfindpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
