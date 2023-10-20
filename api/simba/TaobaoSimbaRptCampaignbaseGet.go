package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbarptcampaignbaseget 推广计划报表基础数据对象
// taobao.simba.rpt.campaignbase.get
//
// 推广计划报表基础数据对象
func Taobaosimbarptcampaignbaseget(clt *core.SDKClient, req *simba.TaobaosimbarptcampaignbasegetAPIRequest, session string) (*simba.TaobaosimbarptcampaignbasegetAPIResponse, error) {
	var resp simba.TaobaosimbarptcampaignbasegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
