package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdTargetTagEstimateUv 标签人群预估
// alibaba.scbp.ad.target.tag.estimate.uv
//
// 标签人群预估
func AlibabaScbpAdTargetTagEstimateUv(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdTargetTagEstimateUvAPIRequest, resp *scbp.AlibabaScbpAdTargetTagEstimateUvAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
