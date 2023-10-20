package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacampaignplatformupdate 更新一个推广计划的平台设置
// taobao.simba.campaign.platform.update
//
// 更新一个推广计划的平台设置
func Taobaosimbacampaignplatformupdate(clt *core.SDKClient, req *simba.TaobaosimbacampaignplatformupdateAPIRequest, session string) (*simba.TaobaosimbacampaignplatformupdateAPIResponse, error) {
	var resp simba.TaobaosimbacampaignplatformupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
