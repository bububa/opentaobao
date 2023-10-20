package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordlistcampaignkeywords 获取计划关键词
// alibaba.scbp.ad.keyword.list.campaign.keywords
//
// 获取计划关键词
func Alibabascbpadkeywordlistcampaignkeywords(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordlistcampaignkeywordsAPIRequest, session string) (*scbp.AlibabascbpadkeywordlistcampaignkeywordsAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordlistcampaignkeywordsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
