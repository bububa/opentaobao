package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpcampaignfindlist 查询全量计划列表(不分页)
// taobao.universalbp.campaign.findlist
//
// 查询场景内的全量计划列表
func Taobaouniversalbpcampaignfindlist(clt *core.SDKClient, req *simba.TaobaouniversalbpcampaignfindlistAPIRequest, session string) (*simba.TaobaouniversalbpcampaignfindlistAPIResponse, error) {
	var resp simba.TaobaouniversalbpcampaignfindlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
