package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品报告 API请求
alibaba.scbp.ad.report.get.product.report

产品报告
*/
type AlibabaScbpAdReportGetProductReportAPIRequest struct {
    model.Params
    // 请求参数
    _productReportOperation   *ProductReportOperationDto
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdReportGetProductReportAPIRequest对象
func NewAlibabaScbpAdReportGetProductReportRequest() *AlibabaScbpAdReportGetProductReportAPIRequest{
    return &AlibabaScbpAdReportGetProductReportAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportGetProductReportAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.get.product.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportGetProductReportAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductReportOperation Setter
// 请求参数
func (r *AlibabaScbpAdReportGetProductReportAPIRequest) SetProductReportOperation(_productReportOperation *ProductReportOperationDto) error {
    r._productReportOperation = _productReportOperation
    r.Set("product_report_operation", _productReportOperation)
    return nil
}

// ProductReportOperation Getter
func (r AlibabaScbpAdReportGetProductReportAPIRequest) GetProductReportOperation() *ProductReportOperationDto {
    return r._productReportOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportGetProductReportAPIRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdReportGetProductReportAPIRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
