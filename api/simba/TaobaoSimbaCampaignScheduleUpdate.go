package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacampaignscheduleupdate 更新一个推广计划的分时折扣设置
// taobao.simba.campaign.schedule.update
//
// 更新一个推广计划的分时折扣设置
func Taobaosimbacampaignscheduleupdate(clt *core.SDKClient, req *simba.TaobaosimbacampaignscheduleupdateAPIRequest, session string) (*simba.TaobaosimbacampaignscheduleupdateAPIResponse, error) {
	var resp simba.TaobaosimbacampaignscheduleupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
