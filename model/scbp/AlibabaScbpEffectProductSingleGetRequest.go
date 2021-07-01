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
type AlibabaScbpEffectProductSingleGetAPIRequest struct {
    model.Params
    // ProductQuery
    _p4pProductReportQuery   *ProductQuery
}

// 初始化AlibabaScbpEffectProductSingleGetAPIRequest对象
func NewAlibabaScbpEffectProductSingleGetRequest() *AlibabaScbpEffectProductSingleGetAPIRequest{
    return &AlibabaScbpEffectProductSingleGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectProductSingleGetAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.effect.product.single.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectProductSingleGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// P4pProductReportQuery Setter
// ProductQuery
func (r *AlibabaScbpEffectProductSingleGetAPIRequest) SetP4pProductReportQuery(_p4pProductReportQuery *ProductQuery) error {
    r._p4pProductReportQuery = _p4pProductReportQuery
    r.Set("p4p_product_report_query", _p4pProductReportQuery)
    return nil
}

// P4pProductReportQuery Getter
func (r AlibabaScbpEffectProductSingleGetAPIRequest) GetP4pProductReportQuery() *ProductQuery {
    return r._p4pProductReportQuery
}
