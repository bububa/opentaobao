package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCampaignFindsubcampaignid 查询无界版计划对应的原场景计划id
// taobao.universalbp.campaign.findsubcampaignid
//
// 查询该场景下，无界版计划对应的原场景的计划
func TaobaoUniversalbpCampaignFindsubcampaignid(clt *core.SDKClient, req *simba.TaobaoUniversalbpCampaignFindsubcampaignidAPIRequest, session string) (*simba.TaobaoUniversalbpCampaignFindsubcampaignidAPIResponse, error) {
	var resp simba.TaobaoUniversalbpCampaignFindsubcampaignidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
