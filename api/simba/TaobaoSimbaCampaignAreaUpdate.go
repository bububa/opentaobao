package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignAreaUpdate 更新一个推广计划的投放地域
// taobao.simba.campaign.area.update
//
// 更新一个推广计划的投放地域
func TaobaoSimbaCampaignAreaUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignAreaUpdateAPIRequest, session string) (*simba.TaobaoSimbaCampaignAreaUpdateAPIResponse, error) {
	var resp simba.TaobaoSimbaCampaignAreaUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
