package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpcampaignfindsubcampaignid 查询无界版计划对应的原场景计划id
// taobao.universalbp.campaign.findsubcampaignid
//
// 查询该场景下，无界版计划对应的原场景的计划
func Taobaouniversalbpcampaignfindsubcampaignid(clt *core.SDKClient, req *simba.TaobaouniversalbpcampaignfindsubcampaignidAPIRequest, session string) (*simba.TaobaouniversalbpcampaignfindsubcampaignidAPIResponse, error) {
	var resp simba.TaobaouniversalbpcampaignfindsubcampaignidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
