package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacampaignscheduleget 取得一个推广计划的分时折扣设置
// taobao.simba.campaign.schedule.get
//
// 取得一个推广计划的分时折扣设置
func Taobaosimbacampaignscheduleget(clt *core.SDKClient, req *simba.TaobaosimbacampaignschedulegetAPIRequest, session string) (*simba.TaobaosimbacampaignschedulegetAPIResponse, error) {
	var resp simba.TaobaosimbacampaignschedulegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
