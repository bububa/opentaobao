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
    _productReportOperation   *ProductReportOperationDTO
    // 用户信息
    _topContext   *TopContextDTO
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
func (r *AlibabaScbpAdReportGetProductReportRequest) SetProductReportOperation(_productReportOperation *ProductReportOperationDTO) error {
    r._productReportOperation = _productReportOperation
    r.Set("product_report_operation", _productReportOperation)
    return nil
}

// ProductReportOperation Getter
func (r AlibabaScbpAdReportGetProductReportRequest) GetProductReportOperation() *ProductReportOperationDTO {
    return r._productReportOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportGetProductReportRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdReportGetProductReportRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
