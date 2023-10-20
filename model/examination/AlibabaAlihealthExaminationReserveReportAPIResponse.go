package examination

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReserveReportAPIResponse 体检机构对接_体检报告查询 API返回值
// alibaba.alihealth.examination.reserve.report
//
// 体检机构对接_体检报告获取
type AlibabaAlihealthExaminationReserveReportAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReserveReportAPIResponseModel
}

// AlibabaAlihealthExaminationReserveReportAPIResponseModel is 体检机构对接_体检报告查询 成功返回结果
type AlibabaAlihealthExaminationReserveReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果编码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 文件数据流
	ReportData string `json:"report_data,omitempty" xml:"report_data,omitempty"`
}
