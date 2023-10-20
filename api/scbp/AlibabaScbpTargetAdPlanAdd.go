package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdPlanAdd 定向推广-新建计划
// alibaba.scbp.target.ad.plan.add
//
// 定向推广-新建单条计划
func AlibabaScbpTargetAdPlanAdd(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanAddAPIRequest, resp *scbp.AlibabaScbpTargetAdPlanAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
