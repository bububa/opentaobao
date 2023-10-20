package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCampaignGet 查询单个计划详情
// taobao.universalbp.campaign.get
//
// 查询单个计划详情信息（不包括报表数据）
func TaobaoUniversalbpCampaignGet(clt *core.SDKClient, req *simba.TaobaoUniversalbpCampaignGetAPIRequest, resp *simba.TaobaoUniversalbpCampaignGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
