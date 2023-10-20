package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdPlanProductListGet 定向推广-获取推广计划产品列表
// alibaba.scbp.target.ad.plan.product.list.get
//
// 定向推广-获取推广计划产品列表
func AlibabaScbpTargetAdPlanProductListGet(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanProductListGetAPIRequest, resp *scbp.AlibabaScbpTargetAdPlanProductListGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
