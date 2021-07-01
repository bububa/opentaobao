package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV获取报告文件查看验证码 API返回值 
alibaba.alihealth.examination.report.diagnose.file.code.get

体检报告人工解读_ISV获取报告文件验证码进行查看报告文件
*/
type AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponseModel
}

// ISV获取报告文件查看验证码 成功返回结果
type AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_file_code_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
