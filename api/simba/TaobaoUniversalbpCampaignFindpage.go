package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCampaignFindpage 查询计划分页列表
// taobao.universalbp.campaign.findpage
//
// 分页查询场景内的计划列表
func TaobaoUniversalbpCampaignFindpage(clt *core.SDKClient, req *simba.TaobaoUniversalbpCampaignFindpageAPIRequest, resp *simba.TaobaoUniversalbpCampaignFindpageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
