package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacampaignbudgetget 取得一个推广计划的日限额
// taobao.simba.campaign.budget.get
//
// 取得一个推广计划的日限额
func Taobaosimbacampaignbudgetget(clt *core.SDKClient, req *simba.TaobaosimbacampaignbudgetgetAPIRequest, session string) (*simba.TaobaosimbacampaignbudgetgetAPIResponse, error) {
	var resp simba.TaobaosimbacampaignbudgetgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
