package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdPlanUpdateTags 定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新
// alibaba.scbp.target.ad.plan.update.tags
//
// 定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新
func AlibabaScbpTargetAdPlanUpdateTags(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanUpdateTagsAPIRequest, resp *scbp.AlibabaScbpTargetAdPlanUpdateTagsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
