package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest 体检报告人工解读订单 API请求
// alibaba.alihealth.examination.report.diagnose.order.submit
//
// 体检报告人工解读订单信息推送给ISV，进行人工解读
type AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest struct {
	model.Params
	// 订单ID
	_orderId string
	// 手机号码，显示后四位
	_mobilePhone string
	// 证件号，显示前1，后1
	_idCardNo string
	// 性别
	_gender string
	// 报告地址
	_reportUrl string
	// 主诉问题
	_question string
	// 咨询人名称
	_patientName string
}

// NewAlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest 初始化AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseOrderSubmitRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest {
	return &AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.order.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetMobilePhone is MobilePhone Setter
// 手机号码，显示后四位
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) SetMobilePhone(_mobilePhone string) error {
	r._mobilePhone = _mobilePhone
	r.Set("mobile_phone", _mobilePhone)
	return nil
}

// GetMobilePhone MobilePhone Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) GetMobilePhone() string {
	return r._mobilePhone
}

// SetIdCardNo is IdCardNo Setter
// 证件号，显示前1，后1
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) SetIdCardNo(_idCardNo string) error {
	r._idCardNo = _idCardNo
	r.Set("id_card_no", _idCardNo)
	return nil
}

// GetIdCardNo IdCardNo Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) GetIdCardNo() string {
	return r._idCardNo
}

// SetGender is Gender Setter
// 性别
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) SetGender(_gender string) error {
	r._gender = _gender
	r.Set("gender", _gender)
	return nil
}

// GetGender Gender Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) GetGender() string {
	return r._gender
}

// SetReportUrl is ReportUrl Setter
// 报告地址
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) SetReportUrl(_reportUrl string) error {
	r._reportUrl = _reportUrl
	r.Set("report_url", _reportUrl)
	return nil
}

// GetReportUrl ReportUrl Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) GetReportUrl() string {
	return r._reportUrl
}

// SetQuestion is Question Setter
// 主诉问题
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) SetQuestion(_question string) error {
	r._question = _question
	r.Set("question", _question)
	return nil
}

// GetQuestion Question Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) GetQuestion() string {
	return r._question
}

// SetPatientName is PatientName Setter
// 咨询人名称
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) SetPatientName(_patientName string) error {
	r._patientName = _patientName
	r.Set("patient_name", _patientName)
	return nil
}

// GetPatientName PatientName Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest) GetPatientName() string {
	return r._patientName
}
