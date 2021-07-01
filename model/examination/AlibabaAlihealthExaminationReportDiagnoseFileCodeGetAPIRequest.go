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
type AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest struct {
    model.Params
    // 报告id
    _reportId   int64
    // 订单id
    _orderId   string
    // 医生id
    _doctorId   string
}

// 初始化AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseFileCodeGetRequest() *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest{
    return &AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.report.diagnose.file.code.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReportId Setter
// 报告id
func (r *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest) SetReportId(_reportId int64) error {
    r._reportId = _reportId
    r.Set("report_id", _reportId)
    return nil
}

// ReportId Getter
func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest) GetReportId() int64 {
    return r._reportId
}
// OrderId Setter
// 订单id
func (r *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest) GetOrderId() string {
    return r._orderId
}
// DoctorId Setter
// 医生id
func (r *AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest) SetDoctorId(_doctorId string) error {
    r._doctorId = _doctorId
    r.Set("doctor_id", _doctorId)
    return nil
}

// DoctorId Getter
func (r AlibabaAlihealthExaminationReportDiagnoseFileCodeGetAPIRequest) GetDoctorId() string {
    return r._doctorId
}
