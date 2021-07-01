package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdPlanTagGetAPIRequest
定向推广-获取计划的定向溢价数据 API请求
alibaba.scbp.target.ad.plan.tag.get

定向推广-获取计划的定向溢价数据 */
type AlibabaScbpTargetAdPlanTagGetAPIRequest struct {
	model.Params
	// 推广计划Id
	_campaignId int64
}

// New
