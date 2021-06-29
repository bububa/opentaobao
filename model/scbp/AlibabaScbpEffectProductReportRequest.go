package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
所有产品报表 API请求
alibaba.scbp.effect.product.report

所有产品报表
*/
type AlibabaScbpEffectProductReportRequest struct {
    model.Params
    // ProductQuery
    _p4pProductReportQuery   *ProductQuery
}

// 初始化AlibabaScbpEffectProductReportRequest对象
func NewAlibabaScbpEffectProductReportRequest() *AlibabaScbpEffectProductReportRequest{
    return &AlibabaScbpEffectProductReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectProductReportRequest) GetApiMethodName() string {
    return "alibaba.scbp.effect.product.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectProductReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// P4pProductReportQuery Setter
// ProductQuery
func (r *AlibabaScbpEffectProductReportRequest) SetP4pProductReportQuery(_p4pProductReportQuery *ProductQuery) error {
    r._p4pProductReportQuery = _p4pProductReportQuery
    r.Set("p4p_product_report_query", _p4pProductReportQuery)
    return nil
}

// P4pProductReportQuery Getter
func (r AlibabaScbpEffectProductReportRequest) GetP4pProductReportQuery() *ProductQuery {
    return r._p4pProductReportQuery
}
