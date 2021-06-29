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
    _orderId   string
    // 手机号码，显示后四位
    _mobilePhone   string
    // 证件号，显示前1，后1
    _idCardNo   string
    // 性别
    _gender   string
    // 报告地址
    _reportUrl   string
    // 主诉问题
    _question   string
    // 咨询人名称
    _patientName   string
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
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetOrderId() string {
    return r._orderId
}
// MobilePhone Setter
// 手机号码，显示后四位
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetMobilePhone(_mobilePhone string) error {
    r._mobilePhone = _mobilePhone
    r.Set("mobile_phone", _mobilePhone)
    return nil
}

// MobilePhone Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetMobilePhone() string {
    return r._mobilePhone
}
// IdCardNo Setter
// 证件号，显示前1，后1
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetIdCardNo(_idCardNo string) error {
    r._idCardNo = _idCardNo
    r.Set("id_card_no", _idCardNo)
    return nil
}

// IdCardNo Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetIdCardNo() string {
    return r._idCardNo
}
// Gender Setter
// 性别
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetGender(_gender string) error {
    r._gender = _gender
    r.Set("gender", _gender)
    return nil
}

// Gender Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetGender() string {
    return r._gender
}
// ReportUrl Setter
// 报告地址
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetReportUrl(_reportUrl string) error {
    r._reportUrl = _reportUrl
    r.Set("report_url", _reportUrl)
    return nil
}

// ReportUrl Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetReportUrl() string {
    return r._reportUrl
}
// Question Setter
// 主诉问题
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetQuestion(_question string) error {
    r._question = _question
    r.Set("question", _question)
    return nil
}

// Question Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetQuestion() string {
    return r._question
}
// PatientName Setter
// 咨询人名称
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) SetPatientName(_patientName string) error {
    r._patientName = _patientName
    r.Set("patient_name", _patientName)
    return nil
}

// PatientName Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest) GetPatientName() string {
    return r._patientName
}
