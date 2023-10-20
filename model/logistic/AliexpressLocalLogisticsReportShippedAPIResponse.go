package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLocalLogisticsReportShippedAPIResponse report as shipped API返回值
// aliexpress.local.logistics.report.shipped
//
// report as shipped
type AliexpressLocalLogisticsReportShippedAPIResponse struct {
	model.CommonResponse
	AliexpressLocalLogisticsReportShippedAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressLocalLogisticsReportShippedAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressLocalLogisticsReportShippedAPIResponseModel).Reset()
}

// AliexpressLocalLogisticsReportShippedAPIResponseModel is report as shipped 成功返回结果
type AliexpressLocalLogisticsReportShippedAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_local_logistics_report_shipped_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// response info
	Data []ReportShippedDto `json:"data,omitempty" xml:"data>report_shipped_dto,omitempty"`
	// error message
	ErrorResultMessage string `json:"error_result_message,omitempty" xml:"error_result_message,omitempty"`
	// error code
	ErrorResultCode string `json:"error_result_code,omitempty" xml:"error_result_code,omitempty"`
	// is success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressLocalLogisticsReportShippedAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = m.Data[:0]
	m.ErrorResultMessage = ""
	m.ErrorResultCode = ""
	m.IsSuccess = false
}

var poolAliexpressLocalLogisticsReportShippedAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressLocalLogisticsReportShippedAPIResponse)
	},
}

// GetAliexpressLocalLogisticsReportShippedAPIResponse 从 sync.Pool 获取 AliexpressLocalLogisticsReportShippedAPIResponse
func GetAliexpressLocalLogisticsReportShippedAPIResponse() *AliexpressLocalLogisticsReportShippedAPIResponse {
	return poolAliexpressLocalLogisticsReportShippedAPIResponse.Get().(*AliexpressLocalLogisticsReportShippedAPIResponse)
}

// ReleaseAliexpressLocalLogisticsReportShippedAPIResponse 将 AliexpressLocalLogisticsReportShippedAPIResponse 保存到 sync.Pool
func ReleaseAliexpressLocalLogisticsReportShippedAPIResponse(v *AliexpressLocalLogisticsReportShippedAPIResponse) {
	v.Reset()
	poolAliexpressLocalLogisticsReportShippedAPIResponse.Put(v)
}
