package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个产品的报表 API请求
alibaba.scbp.effect.product.single.get

单个产品的报表
*/
type AlibabaScbpEffectProductSingleGetRequest struct {
    model.Params
    // ProductQuery
    p4pProductReportQuery   *ProductQuery
}

// 初始化AlibabaScbpEffectProductSingleGetRequest对象
func NewAlibabaScbpEffectProductSingleGetRequest() *AlibabaScbpEffectProductSingleGetRequest{
    return &AlibabaScbpEffectProductSingleGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectProductSingleGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.effect.product.single.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectProductSingleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// P4pProductReportQuery Setter
// ProductQuery
func (r *AlibabaScbpEffectProductSingleGetRequest) SetP4pProductReportQuery(p4pProductReportQuery *ProductQuery) error {
    r.p4pProductReportQuery = p4pProductReportQuery
    r.Set("p4p_product_report_query", p4pProductReportQuery)
    return nil
}

// P4pProductReportQuery Getter
func (r AlibabaScbpEffectProductSingleGetRequest) GetP4pProductReportQuery() *ProductQuery {
    return r.p4pProductReportQuery
}
