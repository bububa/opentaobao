package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导医通报告解读临时消息接收 API请求
alibaba.alihealth.examination.report.diagnose.tempmessage.receive

导医通报告解读临时消息接收
*/
type AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest struct {
    model.Params
    // 入参对象
    reportDiagnoseImMessageRequest   *ReportDiagnoseImMessageRequest
}

// 初始化AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest() *AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest{
    return &AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.report.diagnose.tempmessage.receive"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReportDiagnoseImMessageRequest Setter
// 入参对象
func (r *AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest) SetReportDiagnoseImMessageRequest(reportDiagnoseImMessageRequest *ReportDiagnoseImMessageRequest) error {
    r.reportDiagnoseImMessageRequest = reportDiagnoseImMessageRequest
    r.Set("report_diagnose_im_message_request", reportDiagnoseImMessageRequest)
    return nil
}

// ReportDiagnoseImMessageRequest Getter
func (r AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest) GetReportDiagnoseImMessageRequest() *ReportDiagnoseImMessageRequest {
    return r.reportDiagnoseImMessageRequest
}
