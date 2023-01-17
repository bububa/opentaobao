package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinAppointmentResultNotifyAPIRequest 通知预约结果 API请求
// alibaba.health.vaccin.appointment.result.notify
//
// 和ISV合作，需ISV回传预约结果。
type AlibabaHealthVaccinAppointmentResultNotifyAPIRequest struct {
	model.Params
	// 预约单id
	_orderId string
	// 鹿苗预约单外部id
	_outId string
	// 预约失败原因
	_failReason string
	// 预约成功码
	_successCode string
	// 时间段内序号
	_periodSeqNo int64
	// 预约结果
	_appointResult bool
}

// NewAlibabaHealthVaccinAppointmentResultNotifyRequest 初始化AlibabaHealthVaccinAppointmentResultNotifyAPIRequest对象
func NewAlibabaHealthVaccinAppointmentResultNotifyRequest() *AlibabaHealthVaccinAppointmentResultNotifyAPIRequest {
	return &AlibabaHealthVaccinAppointmentResultNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.appointment.result.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 预约单id
func (r *AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetOutId is OutId Setter
// 鹿苗预约单外部id
func (r *AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// GetOutId OutId Getter
func (r AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) GetOutId() string {
	return r._outId
}

// SetFailReason is FailReason Setter
// 预约失败原因
func (r *AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) SetFailReason(_failReason string) error {
	r._failReason = _failReason
	r.Set("fail_reason", _failReason)
	return nil
}

// GetFailReason FailReason Getter
func (r AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) GetFailReason() string {
	return r._failReason
}

// SetSuccessCode is SuccessCode Setter
// 预约成功码
func (r *AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) SetSuccessCode(_successCode string) error {
	r._successCode = _successCode
	r.Set("success_code", _successCode)
	return nil
}

// GetSuccessCode SuccessCode Getter
func (r AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) GetSuccessCode() string {
	return r._successCode
}

// SetPeriodSeqNo is PeriodSeqNo Setter
// 时间段内序号
func (r *AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) SetPeriodSeqNo(_periodSeqNo int64) error {
	r._periodSeqNo = _periodSeqNo
	r.Set("period_seq_no", _periodSeqNo)
	return nil
}

// GetPeriodSeqNo PeriodSeqNo Getter
func (r AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) GetPeriodSeqNo() int64 {
	return r._periodSeqNo
}

// SetAppointResult is AppointResult Setter
// 预约结果
func (r *AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) SetAppointResult(_appointResult bool) error {
	r._appointResult = _appointResult
	r.Set("appoint_result", _appointResult)
	return nil
}

// GetAppointResult AppointResult Getter
func (r AlibabaHealthVaccinAppointmentResultNotifyAPIRequest) GetAppointResult() bool {
	return r._appointResult
}
