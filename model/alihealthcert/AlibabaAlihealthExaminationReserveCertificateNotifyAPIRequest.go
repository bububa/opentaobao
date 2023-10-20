package alihealthcert

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest 健康证服务商预约结果通知阿里健康 API请求
// alibaba.alihealth.examination.reserve.certificate.notify
//
// 当ISV执行完健康证预约成功之后， 调用通知阿里健康
type AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest struct {
	model.Params
	// 服务商预约凭证
	_uniqReserveCode string
	// 健康预约凭证
	_reserveNumber string
	// 预约成功（ISSUED），预约失败(FAIL)
	_reportStatus string
	// 到检凭证，exam_done状态下，该字段必填
	_checkNo string
	// 预约状态原因，描述预约状态如FAIL时的原因
	_statusReason string
}

// NewAlibabaAlihealthExaminationReserveCertificateNotifyRequest 初始化AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest对象
func NewAlibabaAlihealthExaminationReserveCertificateNotifyRequest() *AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest {
	return &AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.reserve.certificate.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUniqReserveCode is UniqReserveCode Setter
// 服务商预约凭证
func (r *AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
	r._uniqReserveCode = _uniqReserveCode
	r.Set("uniq_reserve_code", _uniqReserveCode)
	return nil
}

// GetUniqReserveCode UniqReserveCode Getter
func (r AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest) GetUniqReserveCode() string {
	return r._uniqReserveCode
}

// SetReserveNumber is ReserveNumber Setter
// 健康预约凭证
func (r *AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// GetReserveNumber ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// SetReportStatus is ReportStatus Setter
// 预约成功（ISSUED），预约失败(FAIL)
func (r *AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest) SetReportStatus(_reportStatus string) error {
	r._reportStatus = _reportStatus
	r.Set("report_status", _reportStatus)
	return nil
}

// GetReportStatus ReportStatus Getter
func (r AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest) GetReportStatus() string {
	return r._reportStatus
}

// SetCheckNo is CheckNo Setter
// 到检凭证，exam_done状态下，该字段必填
func (r *AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest) SetCheckNo(_checkNo string) error {
	r._checkNo = _checkNo
	r.Set("check_no", _checkNo)
	return nil
}

// GetCheckNo CheckNo Getter
func (r AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest) GetCheckNo() string {
	return r._checkNo
}

// SetStatusReason is StatusReason Setter
// 预约状态原因，描述预约状态如FAIL时的原因
func (r *AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest) SetStatusReason(_statusReason string) error {
	r._statusReason = _statusReason
	r.Set("status_reason", _statusReason)
	return nil
}

// GetStatusReason StatusReason Getter
func (r AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest) GetStatusReason() string {
	return r._statusReason
}
