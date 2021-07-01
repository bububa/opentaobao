package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdPlanOperationAPIRequest
定向推广-计划开启/暂停/删除 API请求
alibaba.scbp.target.ad.plan.operation

定向推广-计划开启/暂停/删除 */
type AlibabaScbpTargetAdPlanOperationAPIRequest struct {
	model.Params
	// TopP4pModifyQuickCampaignDTO
	_topP4pModifyQuickCampaignDTO *TopP4pModifyQuickCampaignDto
}

// New
