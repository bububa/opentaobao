package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdPlanUpdateTagsAPIRequest
定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新 API请求
alibaba.scbp.target.ad.plan.update.tags

定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新 */
type AlibabaScbpTargetAdPlanUpdateTagsAPIRequest struct {
	model.Params
	// 系统生成
	_paramTopP4pModifyQuickCampaignTagDTO *TopP4pModifyQuickCampaignTagDto
}

// New
