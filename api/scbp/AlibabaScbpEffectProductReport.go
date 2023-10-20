package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpEffectProductReport 所有产品报表
// alibaba.scbp.effect.product.report
//
// 所有产品报表
func AlibabaScbpEffectProductReport(clt *core.SDKClient, req *scbp.AlibabaScbpEffectProductReportAPIRequest, session string) (*scbp.AlibabaScbpEffectProductReportAPIResponse, error) {
	var resp scbp.AlibabaScbpEffectProductReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
