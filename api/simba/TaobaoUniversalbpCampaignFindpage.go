package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCampaignFindpage 查询计划分页列表
// taobao.universalbp.campaign.findpage
//
// 分页查询场景内的计划列表
func TaobaoUniversalbpCampaignFindpage(clt *core.SDKClient, req *simba.TaobaoUniversalbpCampaignFindpageAPIRequest, session string) (*simba.TaobaoUniversalbpCampaignFindpageAPIResponse, error) {
	var resp simba.TaobaoUniversalbpCampaignFindpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
