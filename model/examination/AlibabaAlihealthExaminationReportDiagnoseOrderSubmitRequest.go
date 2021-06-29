package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检报告人工解读订单 APIRequest
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

func NewAlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest{
    return &AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.report.diagnose.order.submit"
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetMobilePhone(mobilePhone string) error {
    r.mobilePhone = mobilePhone
    r.Set("mobile_phone", mobilePhone)
    return nil
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetMobilePhone() string {
    return r.mobilePhone
}

func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetIdCardNo(idCardNo string) error {
    r.idCardNo = idCardNo
    r.Set("id_card_no", idCardNo)
    return nil
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetIdCardNo() string {
    return r.idCardNo
}

func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetGender(gender string) error {
    r.gender = gender
    r.Set("gender", gender)
    return nil
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetGender() string {
    return r.gender
}

func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetReportUrl(reportUrl string) error {
    r.reportUrl = reportUrl
    r.Set("report_url", reportUrl)
    return nil
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetReportUrl() string {
    return r.reportUrl
}

func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetQuestion(question string) error {
    r.question = question
    r.Set("question", question)
    return nil
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetQuestion() string {
    return r.question
}

func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetPatientName(patientName string) error {
    r.patientName = patientName
    r.Set("patient_name", patientName)
    return nil
}

func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetPatientName() string {
    return r.patientName
}

