package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV获取报告文件查看验证码 API请求
alibaba.alihealth.examination.report.diagnose.file.code.get

体检报告人工解读_ISV获取报告文件验证码进行查看报告文件
*/
type AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest struct {
    model.Params
    // 报告id
    reportId   int64
    // 订单id
    orderId   string
    // 医生id
    doctorId   string
}

// 初始化AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest() *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest{
    return &AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.report.diagnose.file.code.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReportId Setter
// 报告id
func (r *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) SetReportId(reportId int64) error {
    r.reportId = reportId
    r.Set("report_id", reportId)
    return nil
}

// ReportId Getter
func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) GetReportId() int64 {
    return r.reportId
}
// OrderId Setter
// 订单id
func (r *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) GetOrderId() string {
    return r.orderId
}
// DoctorId Setter
// 医生id
func (r *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) SetDoctorId(doctorId string) error {
    r.doctorId = doctorId
    r.Set("doctor_id", doctorId)
    return nil
}

// DoctorId Getter
func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) GetDoctorId() string {
    return r.doctorId
}
