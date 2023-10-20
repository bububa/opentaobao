package examination

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReserveReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationReserveReportAPIResponseModel).Reset()
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

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReserveReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.ResponseCode = ""
	m.ReportData = ""
}

var poolAlibabaAlihealthExaminationReserveReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationReserveReportAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationReserveReportAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationReserveReportAPIResponse
func GetAlibabaAlihealthExaminationReserveReportAPIResponse() *AlibabaAlihealthExaminationReserveReportAPIResponse {
	return poolAlibabaAlihealthExaminationReserveReportAPIResponse.Get().(*AlibabaAlihealthExaminationReserveReportAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationReserveReportAPIResponse 将 AlibabaAlihealthExaminationReserveReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationReserveReportAPIResponse(v *AlibabaAlihealthExaminationReserveReportAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationReserveReportAPIResponse.Put(v)
}
