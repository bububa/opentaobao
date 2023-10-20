package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacampaignareaupdate 更新一个推广计划的投放地域
// taobao.simba.campaign.area.update
//
// 更新一个推广计划的投放地域
func Taobaosimbacampaignareaupdate(clt *core.SDKClient, req *simba.TaobaosimbacampaignareaupdateAPIRequest, session string) (*simba.TaobaosimbacampaignareaupdateAPIResponse, error) {
	var resp simba.TaobaosimbacampaignareaupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
