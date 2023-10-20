package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadcampaignfindcampaigneffect 批量查询计划效果数据
// alibaba.scbp.ad.campaign.find.campaign.effect
//
// 批量查询计划效果数据
func Alibabascbpadcampaignfindcampaigneffect(clt *core.SDKClient, req *scbp.AlibabascbpadcampaignfindcampaigneffectAPIRequest, session string) (*scbp.AlibabascbpadcampaignfindcampaigneffectAPIResponse, error) {
	var resp scbp.AlibabascbpadcampaignfindcampaigneffectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
