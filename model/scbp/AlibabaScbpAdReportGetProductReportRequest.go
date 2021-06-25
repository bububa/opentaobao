package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品报告 APIRequest
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

func NewAlibabaScbpAdReportGetProductReportRequest() *AlibabaScbpAdReportGetProductReportRequest{
    return &AlibabaScbpAdReportGetProductReportRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdReportGetProductReportRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.get.product.report"
}

func (r AlibabaScbpAdReportGetProductReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdReportGetProductReportRequest) SetProductReportOperation(productReportOperation *ProductReportOperationDto) error {
    r.productReportOperation = productReportOperation
    r.Set("product_report_operation", productReportOperation)
    return nil
}

func (r AlibabaScbpAdReportGetProductReportRequest) GetProductReportOperation() *ProductReportOperationDto {
    return r.productReportOperation
}

func (r *AlibabaScbpAdReportGetProductReportRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdReportGetProductReportRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

