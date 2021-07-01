package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

/* AlibabaScbpTargetAdPlanList
定向推广-查询定向推广计划列表并返回计划基础信息
alibaba.scbp.target.ad.plan.list

定向推广-查询定向推广计划列表并返回计划基础信息 */
func AlibabaScbpTargetAdPlanList(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanListAPIRequest, session string) (*scbp.AlibabaScbpTargetAdPlanListAPIResponse, error) {
	var resp scbp.AlibabaScbpTargetAdPlanListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
