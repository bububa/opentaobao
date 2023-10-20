package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpEffectProductReport 所有产品报表
// alibaba.scbp.effect.product.report
//
// 所有产品报表
func AlibabaScbpEffectProductReport(clt *core.SDKClient, req *scbp.AlibabaScbpEffectProductReportAPIRequest, resp *scbp.AlibabaScbpEffectProductReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
