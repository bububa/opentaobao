package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdCampaignFindForbiddenKeyword 查询屏蔽词
// alibaba.scbp.ad.campaign.find.forbidden.keyword
//
// 查询屏蔽词
func AlibabaScbpAdCampaignFindForbiddenKeyword(clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest, resp *scbp.AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
