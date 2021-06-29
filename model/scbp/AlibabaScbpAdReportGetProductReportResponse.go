package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品报告 API返回值 
alibaba.scbp.ad.report.get.product.report

产品报告
*/
type AlibabaScbpAdReportGetProductReportAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdReportGetProductReportResponse
}

// 产品报告 成功返回结果
type AlibabaScbpAdReportGetProductReportResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_report_get_product_report_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回数据
    Result   *ProductReportDTO `json:"result,omitempty" xml:"result,omitempty"`
}
