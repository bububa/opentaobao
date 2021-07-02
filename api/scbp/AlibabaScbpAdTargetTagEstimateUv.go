package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdTargetTagEstimateUv 标签人群预估
// alibaba.scbp.ad.target.tag.estimate.uv
//
// 标签人群预估
func AlibabaScbpAdTargetTagEstimateUv(clt *core.SDKClient, req *scbp.AlibabaScbpAdTargetTagEstimateUvAPIRequest, session string) (*scbp.AlibabaScbpAdTargetTagEstimateUvAPIResponse, error) {
	var resp scbp.AlibabaScbpAdTargetTagEstimateUvAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
