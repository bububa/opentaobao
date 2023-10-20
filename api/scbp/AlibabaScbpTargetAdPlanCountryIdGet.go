package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdPlanCountryIdGet 定向推广-国家标签ID获取
// alibaba.scbp.target.ad.plan.country.id.get
//
// 定向推广-国家标签ID获取
func AlibabaScbpTargetAdPlanCountryIdGet(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanCountryIdGetAPIRequest, resp *scbp.AlibabaScbpTargetAdPlanCountryIdGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
