package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdPlanProductListGetAPIRequest
定向推广-获取推广计划产品列表 API请求
alibaba.scbp.target.ad.plan.product.list.get

定向推广-获取推广计划产品列表 */
type AlibabaScbpTargetAdPlanProductListGetAPIRequest struct {
	model.Params
	// TopP4pQuickProductQuery
	_topP4pQuickProductQuery *TopP4pQuickProductQuery
}

// New
