package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdCampaignFindForbiddenKeyword 查询屏蔽词
// alibaba.scbp.ad.campaign.find.forbidden.keyword
//
// 查询屏蔽词
func AlibabaScbpAdCampaignFindForbiddenKeyword(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest, resp *scbp.AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
