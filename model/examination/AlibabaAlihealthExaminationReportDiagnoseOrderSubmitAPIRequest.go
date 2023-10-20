package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest 体检报告人工解读订单 API请求
// alibaba.alihealth.examination.report.diagnose.order.submit
//
// 体检报告人工解读订单信息推送给ISV，进行人工解读
type AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest struct {
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

// NewAlibabaalihealthexaminationreportdiagnoseordersubmitRequest 初始化AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest对象
func NewAlibabaalihealthexaminationreportdiagnoseordersubmitRequest() *AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest {
	return &AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.order.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetMobilePhone is MobilePhone Setter
// 手机号码，显示后四位
func (r *AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) SetMobilePhone(_mobilePhone string) error {
	r._mobilePhone = _mobilePhone
	r.Set("mobile_phone", _mobilePhone)
	return nil
}

// GetMobilePhone MobilePhone Getter
func (r AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) GetMobilePhone() string {
	return r._mobilePhone
}

// SetIdCardNo is IdCardNo Setter
// 证件号，显示前1，后1
func (r *AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) SetIdCardNo(_idCardNo string) error {
	r._idCardNo = _idCardNo
	r.Set("id_card_no", _idCardNo)
	return nil
}

// GetIdCardNo IdCardNo Getter
func (r AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) GetIdCardNo() string {
	return r._idCardNo
}

// SetGender is Gender Setter
// 性别
func (r *AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) SetGender(_gender string) error {
	r._gender = _gender
	r.Set("gender", _gender)
	return nil
}

// GetGender Gender Getter
func (r AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) GetGender() string {
	return r._gender
}

// SetReportUrl is ReportUrl Setter
// 报告地址
func (r *AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) SetReportUrl(_reportUrl string) error {
	r._reportUrl = _reportUrl
	r.Set("report_url", _reportUrl)
	return nil
}

// GetReportUrl ReportUrl Getter
func (r AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) GetReportUrl() string {
	return r._reportUrl
}

// SetQuestion is Question Setter
// 主诉问题
func (r *AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) SetQuestion(_question string) error {
	r._question = _question
	r.Set("question", _question)
	return nil
}

// GetQuestion Question Getter
func (r AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) GetQuestion() string {
	return r._question
}

// SetPatientName is PatientName Setter
// 咨询人名称
func (r *AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) SetPatientName(_patientName string) error {
	r._patientName = _patientName
	r.Set("patient_name", _patientName)
	return nil
}

// GetPatientName PatientName Getter
func (r AlibabaalihealthexaminationreportdiagnoseordersubmitAPIRequest) GetPatientName() string {
	return r._patientName
}
