package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

/* AlibabaScbpTargetAdPlanUpdate
定向推广-更新推广计划的基础信息
alibaba.scbp.target.ad.plan.update

定向推广-更新推广计划的基础信息 */
func AlibabaScbpTargetAdPlanUpdate(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanUpdateAPIRequest, session string) (*scbp.AlibabaScbpTargetAdPlanUpdateAPIResponse, error) {
	var resp scbp.AlibabaScbpTargetAdPlanUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
