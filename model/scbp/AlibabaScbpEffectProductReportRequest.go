package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
所有产品报表 APIRequest
alibaba.scbp.effect.product.report

所有产品报表
*/
type AlibabaScbpEffectProductReportRequest struct {
    model.Params

    // ProductQuery
    p4pProductReportQuery   *ProductQuery 

}

func NewAlibabaScbpEffectProductReportRequest() *AlibabaScbpEffectProductReportRequest{
    return &AlibabaScbpEffectProductReportRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpEffectProductReportRequest) GetApiMethodName() string {
    return "alibaba.scbp.effect.product.report"
}

func (r AlibabaScbpEffectProductReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpEffectProductReportRequest) SetP4pProductReportQuery(p4pProductReportQuery *ProductQuery) error {
    r.p4pProductReportQuery = p4pProductReportQuery
    r.Set("p4p_product_report_query", p4pProductReportQuery)
    return nil
}

func (r AlibabaScbpEffectProductReportRequest) GetP4pProductReportQuery() *ProductQuery {
    return r.p4pProductReportQuery
}

