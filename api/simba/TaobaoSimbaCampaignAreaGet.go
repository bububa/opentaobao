package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignAreaGet 取得一个推广计划的投放地域设置
// taobao.simba.campaign.area.get
//
// 取得一个推广计划的投放地域设置
func TaobaoSimbaCampaignAreaGet(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignAreaGetAPIRequest, session string) (*simba.TaobaoSimbaCampaignAreaGetAPIResponse, error) {
	var resp simba.TaobaoSimbaCampaignAreaGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
