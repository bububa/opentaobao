package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacampaignareaget 取得一个推广计划的投放地域设置
// taobao.simba.campaign.area.get
//
// 取得一个推广计划的投放地域设置
func Taobaosimbacampaignareaget(clt *core.SDKClient, req *simba.TaobaosimbacampaignareagetAPIRequest, session string) (*simba.TaobaosimbacampaignareagetAPIResponse, error) {
	var resp simba.TaobaosimbacampaignareagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
