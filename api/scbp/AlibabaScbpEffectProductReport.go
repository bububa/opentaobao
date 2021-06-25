package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
所有产品报表 
alibaba.scbp.effect.product.report

所有产品报表
*/
func AlibabaScbpEffectProductReport(clt *core.SDKClient, req *scbp.AlibabaScbpEffectProductReportRequest, session string) (*scbp.AlibabaScbpEffectProductReportResponse, error) {
    var resp scbp.AlibabaScbpEffectProductReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
