package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdPlanAddAPIRequest
定向推广-新建计划 API请求
alibaba.scbp.target.ad.plan.add

定向推广-新建单条计划 */
type AlibabaScbpTargetAdPlanAddAPIRequest struct {
	model.Params
	// 定向推广基础信息
	_topP4pBasicQuickCampaign *BasicQuickCampaign
}

// New
