package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdCampaignCreateForbiddenKeyword 创建屏蔽词
// alibaba.scbp.ad.campaign.create.forbidden.keyword
//
// 创建屏蔽词
func AlibabaScbpAdCampaignCreateForbiddenKeyword(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest, resp *scbp.AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
