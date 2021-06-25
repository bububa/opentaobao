package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个产品的报表 APIRequest
alibaba.scbp.effect.product.single.get

单个产品的报表
*/
type AlibabaScbpEffectProductSingleGetRequest struct {
    model.Params

    // ProductQuery
    p4pProductReportQuery   *ProductQuery 

}

func NewAlibabaScbpEffectProductSingleGetRequest() *AlibabaScbpEffectProductSingleGetRequest{
    return &AlibabaScbpEffectProductSingleGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpEffectProductSingleGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.effect.product.single.get"
}

func (r AlibabaScbpEffectProductSingleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpEffectProductSingleGetRequest) SetP4pProductReportQuery(p4pProductReportQuery *ProductQuery) error {
    r.p4pProductReportQuery = p4pProductReportQuery
    r.Set("p4p_product_report_query", p4pProductReportQuery)
    return nil
}

func (r AlibabaScbpEffectProductSingleGetRequest) GetP4pProductReportQuery() *ProductQuery {
    return r.p4pProductReportQuery
}

