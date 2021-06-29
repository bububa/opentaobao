package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检报告人工解读订单 API请求
alibaba.alihealth.examination.report.diagnose.order.submit

体检报告人工解读订单信息推送给ISV，进行人工解读
*/
type AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest struct {
    model.Params
    // 订单ID
    orderId   string
    // 手机号码，显示后四位
    mobilePhone   string
    // 证件号，显示前1，后1
    idCardNo   string
    // 性别
    gender   string
    // 报告地址
    reportUrl   string
    // 主诉问题
    question   string
    // 咨询人名称
    patientName   string
}

// 初始化AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest{
    return &AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.report.diagnose.order.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单ID
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetOrderId() string {
    return r.orderId
}
// MobilePhone Setter
// 手机号码，显示后四位
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetMobilePhone(mobilePhone string) error {
    r.mobilePhone = mobilePhone
    r.Set("mobile_phone", mobilePhone)
    return nil
}

// MobilePhone Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetMobilePhone() string {
    return r.mobilePhone
}
// IdCardNo Setter
// 证件号，显示前1，后1
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetIdCardNo(idCardNo string) error {
    r.idCardNo = idCardNo
    r.Set("id_card_no", idCardNo)
    return nil
}

// IdCardNo Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetIdCardNo() string {
    return r.idCardNo
}
// Gender Setter
// 性别
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetGender(gender string) error {
    r.gender = gender
    r.Set("gender", gender)
    return nil
}

// Gender Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetGender() string {
    return r.gender
}
// ReportUrl Setter
// 报告地址
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetReportUrl(reportUrl string) error {
    r.reportUrl = reportUrl
    r.Set("report_url", reportUrl)
    return nil
}

// ReportUrl Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetReportUrl() string {
    return r.reportUrl
}
// Question Setter
// 主诉问题
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetQuestion(question string) error {
    r.question = question
    r.Set("question", question)
    return nil
}

// Question Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetQuestion() string {
    return r.question
}
// PatientName Setter
// 咨询人名称
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetPatientName(patientName string) error {
    r.patientName = patientName
    r.Set("patient_name", patientName)
    return nil
}

// PatientName Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetPatientName() string {
    return r.patientName
}
