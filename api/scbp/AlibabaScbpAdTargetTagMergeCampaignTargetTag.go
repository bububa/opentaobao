package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdTargetTagMergeCampaignTargetTag 标签增删改
// alibaba.scbp.ad.target.tag.merge.campaign.target.tag
//
// 标签增删改
func AlibabaScbpAdTargetTagMergeCampaignTargetTag(clt *core.SDKClient, req *scbp.AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIRequest, resp *scbp.AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
