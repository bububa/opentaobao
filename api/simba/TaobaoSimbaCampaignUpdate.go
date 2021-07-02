package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignUpdate 更新一个推广计划
// taobao.simba.campaign.update
//
// 更新一个推广计划，可以设置推广计划名字，修改推广计划上下线状态。
func TaobaoSimbaCampaignUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignUpdateAPIRequest, session string) (*simba.TaobaoSimbaCampaignUpdateAPIResponse, error) {
	var resp simba.TaobaoSimbaCampaignUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
