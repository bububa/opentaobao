package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpresslocallogisticsreportshippedAPIResponse report as shipped API返回值
// aliexpress.local.logistics.report.shipped
//
// report as shipped
type AliexpresslocallogisticsreportshippedAPIResponse struct {
	model.CommonResponse
	AliexpresslocallogisticsreportshippedAPIResponseModel
}

// AliexpresslocallogisticsreportshippedAPIResponseModel is report as shipped 成功返回结果
type AliexpresslocallogisticsreportshippedAPIResponseModel struct {
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
