package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdTargetTagFindCampaignTargetTag 查询标签数据
// alibaba.scbp.ad.target.tag.find.campaign.target.tag
//
// 查询标签数据
func AlibabaScbpAdTargetTagFindCampaignTargetTag(clt *core.SDKClient, req *scbp.AlibabaScbpAdTargetTagFindCampaignTargetTagAPIRequest, resp *scbp.AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
