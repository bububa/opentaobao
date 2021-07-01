package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

/* AlibabaScbpTargetAdPlanCountryIdGet
定向推广-国家标签ID获取
alibaba.scbp.target.ad.plan.country.id.get

定向推广-国家标签ID获取 */
func AlibabaScbpTargetAdPlanCountryIdGet(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanCountryIdGetAPIRequest, session string) (*scbp.AlibabaScbpTargetAdPlanCountryIdGetAPIResponse, error) {
	var resp scbp.AlibabaScbpTargetAdPlanCountryIdGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
