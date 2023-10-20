package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbartrptcampaignget 获取推广计划实时报表数据
// taobao.simba.rtrpt.campaign.get
//
// 获取推广计划实时报表数据
func Taobaosimbartrptcampaignget(clt *core.SDKClient, req *simba.TaobaosimbartrptcampaigngetAPIRequest, session string) (*simba.TaobaosimbartrptcampaigngetAPIResponse, error) {
	var resp simba.TaobaosimbartrptcampaigngetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
