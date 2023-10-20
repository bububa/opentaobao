package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportGetProductReportAPIResponse 产品报告 API返回值
// alibaba.scbp.ad.report.get.product.report
//
// 产品报告
type AlibabaScbpAdReportGetProductReportAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdReportGetProductReportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdReportGetProductReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdReportGetProductReportAPIResponseModel).Reset()
}

// AlibabaScbpAdReportGetProductReportAPIResponseModel is 产品报告 成功返回结果
type AlibabaScbpAdReportGetProductReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_report_get_product_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	Result *ProductReportDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdReportGetProductReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaScbpAdReportGetProductReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdReportGetProductReportAPIResponse)
	},
}

// GetAlibabaScbpAdReportGetProductReportAPIResponse 从 sync.Pool 获取 AlibabaScbpAdReportGetProductReportAPIResponse
func GetAlibabaScbpAdReportGetProductReportAPIResponse() *AlibabaScbpAdReportGetProductReportAPIResponse {
	return poolAlibabaScbpAdReportGetProductReportAPIResponse.Get().(*AlibabaScbpAdReportGetProductReportAPIResponse)
}

// ReleaseAlibabaScbpAdReportGetProductReportAPIResponse 将 AlibabaScbpAdReportGetProductReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdReportGetProductReportAPIResponse(v *AlibabaScbpAdReportGetProductReportAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdReportGetProductReportAPIResponse.Put(v)
}
