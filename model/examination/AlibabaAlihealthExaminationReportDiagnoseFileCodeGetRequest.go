package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV获取报告文件查看验证码 APIRequest
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

func NewAlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest() *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest{
    return &AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.report.diagnose.file.code.get"
}

func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) SetReportId(reportId int64) error {
    r.reportId = reportId
    r.Set("report_id", reportId)
    return nil
}

func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) GetReportId() int64 {
    return r.reportId
}

func (r *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) SetDoctorId(doctorId string) error {
    r.doctorId = doctorId
    r.Set("doctor_id", doctorId)
    return nil
}

func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest) GetDoctorId() string {
    return r.doctorId
}

