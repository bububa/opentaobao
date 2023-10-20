package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdPlanUpdateProducts 定向推广 按照id操作推广计划的产品，包括新增，删除和更新
// alibaba.scbp.target.ad.plan.update.products
//
// 定向推广 按照id操作推广计划的产品，包括新增，删除和更新
func AlibabaScbpTargetAdPlanUpdateProducts(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanUpdateProductsAPIRequest, resp *scbp.AlibabaScbpTargetAdPlanUpdateProductsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
