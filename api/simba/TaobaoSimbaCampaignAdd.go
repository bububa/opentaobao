package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacampaignadd 创建一个推广计划
// taobao.simba.campaign.add
//
// 创建一个推广计划
func Taobaosimbacampaignadd(clt *core.SDKClient, req *simba.TaobaosimbacampaignaddAPIRequest, session string) (*simba.TaobaosimbacampaignaddAPIResponse, error) {
	var resp simba.TaobaosimbacampaignaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
