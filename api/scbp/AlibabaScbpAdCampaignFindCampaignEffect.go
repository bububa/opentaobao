package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdCampaignFindCampaignEffect 批量查询计划效果数据
// alibaba.scbp.ad.campaign.find.campaign.effect
//
// 批量查询计划效果数据
func AlibabaScbpAdCampaignFindCampaignEffect(clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignFindCampaignEffectAPIRequest, session string) (*scbp.AlibabaScbpAdCampaignFindCampaignEffectAPIResponse, error) {
	var resp scbp.AlibabaScbpAdCampaignFindCampaignEffectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
