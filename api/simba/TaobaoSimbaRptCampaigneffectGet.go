package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptCampaigneffectGet 推广计划效果报表数据对象
// taobao.simba.rpt.campaigneffect.get
//
// 推广计划效果报表数据对象
func TaobaoSimbaRptCampaigneffectGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptCampaigneffectGetAPIRequest, session string) (*simba.TaobaoSimbaRptCampaigneffectGetAPIResponse, error) {
	var resp simba.TaobaoSimbaRptCampaigneffectGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
