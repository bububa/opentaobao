package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbasalestarcampaignbudgetupdate 销量明星跟新预算相关接口
// taobao.simba.salestar.campaign.budget.update
//
// 更新一个推广计划的日限额
func Taobaosimbasalestarcampaignbudgetupdate(clt *core.SDKClient, req *simba.TaobaosimbasalestarcampaignbudgetupdateAPIRequest, session string) (*simba.TaobaosimbasalestarcampaignbudgetupdateAPIResponse, error) {
	var resp simba.TaobaosimbasalestarcampaignbudgetupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
