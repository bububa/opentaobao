package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordListCampaignKeywords 获取计划关键词
// alibaba.scbp.ad.keyword.list.campaign.keywords
//
// 获取计划关键词
func AlibabaScbpAdKeywordListCampaignKeywords(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordListCampaignKeywordsAPIRequest, resp *scbp.AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
