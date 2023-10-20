package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacampaignplatformget 取得一个推广计划的投放平台设置
// taobao.simba.campaign.platform.get
//
// 获得一个推广计划的投放平台设置
func Taobaosimbacampaignplatformget(clt *core.SDKClient, req *simba.TaobaosimbacampaignplatformgetAPIRequest, session string) (*simba.TaobaosimbacampaignplatformgetAPIResponse, error) {
	var resp simba.TaobaosimbacampaignplatformgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
