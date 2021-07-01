package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdPlanUpdateAPIRequest
定向推广-更新推广计划的基础信息 API请求
alibaba.scbp.target.ad.plan.update

定向推广-更新推广计划的基础信息 */
type AlibabaScbpTargetAdPlanUpdateAPIRequest struct {
	model.Params
	// TopP4pBasicQuickCampaign
	_topP4pBasicQuickCampaign *TopP4pBasicQuickCampaign
}

// New
