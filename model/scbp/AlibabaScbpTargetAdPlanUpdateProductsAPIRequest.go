package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdPlanUpdateProductsAPIRequest
定向推广 按照id操作推广计划的产品，包括新增，删除和更新 API请求
alibaba.scbp.target.ad.plan.update.products

定向推广 按照id操作推广计划的产品，包括新增，删除和更新 */
type AlibabaScbpTargetAdPlanUpdateProductsAPIRequest struct {
	model.Params
	// 系统生成
	_paramTopP4pModifyQuickCampaignProductDTO *TopP4pModifyQuickCampaignProductDto
}

// New
