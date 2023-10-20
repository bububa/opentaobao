package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptargetadcampaigneffect 定向推广-获取计划维度推广效果
// alibaba.scbp.target.ad.campaign.effect
//
// 定向推广-获取计划维度推广效果
func Alibabascbptargetadcampaigneffect(clt *core.SDKClient, req *scbp.AlibabascbptargetadcampaigneffectAPIRequest, session string) (*scbp.AlibabascbptargetadcampaigneffectAPIResponse, error) {
	var resp scbp.AlibabascbptargetadcampaigneffectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
