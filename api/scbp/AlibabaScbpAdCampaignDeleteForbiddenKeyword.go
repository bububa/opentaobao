package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdCampaignDeleteForbiddenKeyword 删除屏蔽词
// alibaba.scbp.ad.campaign.delete.forbidden.keyword
//
// 删除屏蔽词
func AlibabaScbpAdCampaignDeleteForbiddenKeyword(clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest, resp *scbp.AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
