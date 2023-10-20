package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptargetadcampaignproducteffect 定向推广-获取计划中产品推广效果
// alibaba.scbp.target.ad.campaign.product.effect
//
// 定向推广-获取计划中产品推广效果
func Alibabascbptargetadcampaignproducteffect(clt *core.SDKClient, req *scbp.AlibabascbptargetadcampaignproducteffectAPIRequest, session string) (*scbp.AlibabascbptargetadcampaignproducteffectAPIResponse, error) {
	var resp scbp.AlibabascbptargetadcampaignproducteffectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
