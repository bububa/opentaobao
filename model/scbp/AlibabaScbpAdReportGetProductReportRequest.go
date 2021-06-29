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
type AlibabaScbpAdReportGetProductReportRequest struct {
    model.Params
    // 请求参数
    productReportOperation   *ProductReportOperationDto
    // 用户信息
    topContext   *TopContextDto
}

// 初始化AlibabaScbpAdReportGetProductReportRequest对象
func NewAlibabaScbpAdReportGetProductReportRequest() *AlibabaScbpAdReportGetProductReportRequest{
    return &AlibabaScbpAdReportGetProductReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportGetProductReportRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.get.product.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportGetProductReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductReportOperation Setter
// 请求参数
func (r *AlibabaScbpAdReportGetProductReportRequest) SetProductReportOperation(productReportOperation *ProductReportOperationDto) error {
    r.productReportOperation = productReportOperation
    r.Set("product_report_operation", productReportOperation)
    return nil
}

// ProductReportOperation Getter
func (r AlibabaScbpAdReportGetProductReportRequest) GetProductReportOperation() *ProductReportOperationDto {
    return r.productReportOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportGetProductReportRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdReportGetProductReportRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
