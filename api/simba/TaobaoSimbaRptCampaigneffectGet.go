package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbarptcampaigneffectget 推广计划效果报表数据对象
// taobao.simba.rpt.campaigneffect.get
//
// 推广计划效果报表数据对象
func Taobaosimbarptcampaigneffectget(clt *core.SDKClient, req *simba.TaobaosimbarptcampaigneffectgetAPIRequest, session string) (*simba.TaobaosimbarptcampaigneffectgetAPIResponse, error) {
	var resp simba.TaobaosimbarptcampaigneffectgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
