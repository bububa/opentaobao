package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品报告 APIResponse
alibaba.scbp.ad.report.get.product.report

产品报告
*/
type AlibabaScbpAdReportGetProductReportAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_report_get_product_report_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回数据
    
    Result   *ProductReportDto `json:"result,omitempty" xml:"