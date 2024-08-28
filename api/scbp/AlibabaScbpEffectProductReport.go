package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpEffectProductReport 所有产品报表
// alibaba.scbp.effect.product.report
//
// 所有产品报表
func AlibabaScbpEffectProductReport(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpEffectProductReportAPIRequest, resp *scbp.AlibabaScbpEffectProductReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
