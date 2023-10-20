package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacampaignbudgetupdate 更新一个推广计划的日限额
// taobao.simba.campaign.budget.update
//
// 更新一个推广计划的日限额
func Taobaosimbacampaignbudgetupdate(clt *core.SDKClient, req *simba.TaobaosimbacampaignbudgetupdateAPIRequest, session string) (*simba.TaobaosimbacampaignbudgetupdateAPIResponse, error) {
	var resp simba.TaobaosimbacampaignbudgetupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
