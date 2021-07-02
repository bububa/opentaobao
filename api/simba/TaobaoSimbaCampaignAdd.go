package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignAdd 创建一个推广计划
// taobao.simba.campaign.add
//
// 创建一个推广计划
func TaobaoSimbaCampaignAdd(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignAddAPIRequest, session string) (*simba.TaobaoSimbaCampaignAddAPIResponse, error) {
	var resp simba.TaobaoSimbaCampaignAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
