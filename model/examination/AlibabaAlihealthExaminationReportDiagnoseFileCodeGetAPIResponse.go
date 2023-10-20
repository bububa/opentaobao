package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse ISV获取报告文件查看验证码 API返回值
// alibaba.alihealth.examination.report.diagnose.file.code.get
//
// 体检报告人工解读_ISV获取报告文件验证码进行查看报告文件
type AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponseModel is ISV获取报告文件查看验证码 成功返回结果
type AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_file_code_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse
func GetAlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse() *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse {
	return poolAlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse.Get().(*AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse 将 AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse(v *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse.Put(v)
}
