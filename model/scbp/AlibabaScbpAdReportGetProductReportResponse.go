package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品报告 APIResponse
alibaba.scbp.ad.report.get.product.report

产品报告
*/
type AlibabaScbpAdReportGetProductReportAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdReportGetProductReportResponse `json:"alibaba_scbp_ad_report_get_product_report_response,omitempty"` 
    AlibabaScbpAdReportGetProductReportResponse
}

/* model for simplify = false
type AlibabaScbpAdReportGetProductReportResponse struct {

    // 返回数据
    
    Result  *struct {
        ProductReportDto  *ProductReportDto `json:"product_report_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdReportGetProductReportResponse struct {

    // 返回数据
    Result   *ProductReportDto `json:"result,omitempty"`

}
