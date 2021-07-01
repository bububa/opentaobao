package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdPlanListAPIRequest
定向推广-查询定向推广计划列表并返回计划基础信息 API请求
alibaba.scbp.target.ad.plan.list

定向推广-查询定向推广计划列表并返回计划基础信息 */
type AlibabaScbpTargetAdPlanListAPIRequest struct {
	model.Params
	// TopP4pQuickCampaignQuery
	_topP4pQuickCampaignQuery *TopP4pQuickCampaignQueryDto
}

// New
