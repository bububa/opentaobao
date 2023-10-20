package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadcampaignfindcampaignpage 分页查询计划
// alibaba.scbp.ad.campaign.find.campaign.page
//
// 分页查询计划
func Alibabascbpadcampaignfindcampaignpage(clt *core.SDKClient, req *scbp.AlibabascbpadcampaignfindcampaignpageAPIRequest, session string) (*scbp.AlibabascbpadcampaignfindcampaignpageAPIResponse, error) {
	var resp scbp.AlibabascbpadcampaignfindcampaignpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
