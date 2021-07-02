package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdCampaignTagEffect 定向推广-获取推广计划定向效果数据
// alibaba.scbp.target.ad.campaign.tag.effect
//
// 定向推广-获取推广计划定向效果数据
func AlibabaScbpTargetAdCampaignTagEffect(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdCampaignTagEffectAPIRequest, session string) (*scbp.AlibabaScbpTargetAdCampaignTagEffectAPIResponse, error) {
	var resp scbp.AlibabaScbpTargetAdCampaignTagEffectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
