package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReserveStateNotifyAPIRequest 体检机构对接_体检状态主动通知 API请求
// alibaba.alihealth.examination.reserve.state.notify
//
// 到了体检当天后，服务商主动通知体检预约状态
type AlibabaAlihealthExaminationReserveStateNotifyAPIRequest struct {
	model.Params
	// 服务商预约凭证
	_uniqReserveCode string
	// 健康预约凭证
	_reserveNumber string
	// 体检状态：未到检(exam_not), 已到检(exam_done)； 上门服务中还需以下两种状态：预约确认中（reserve_confirming），预约拒绝（reserve_rejected）；
	_reportStatus string
	// 到检凭证，exam_done状态下，该字段必填
	_checkNo string
}

// NewAlibabaAlihealthExaminationReserveStateNotifyRequest 初始化AlibabaAlihealthExaminationReserveStateNotifyAPIRequest对象
func NewAlibabaAlihealthExaminationReserveStateNotifyRequest() *AlibabaAlihealthExaminationReserveStateNotifyAPIRequest {
	return &AlibabaAlihealthExaminationReserveStateNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveStateNotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.reserve.state.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveStateNotifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUniqReserveCode is UniqReserveCode Setter
// 服务商预约凭证
func (r *AlibabaAlihealthExaminationReserveStateNotifyAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
	r._uniqReserveCode = _uniqReserveCode
	r.Set("uniq_reserve_code", _uniqReserveCode)
	return nil
}

// GetUniqReserveCode UniqReserveCode Getter
func (r AlibabaAlihealthExaminationReserveStateNotifyAPIRequest) GetUniqReserveCode() string {
	return r._uniqReserveCode
}

// SetReserveNumber is ReserveNumber Setter
// 健康预约凭证
func (r *AlibabaAlihealthExaminationReserveStateNotifyAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// GetReserveNumber ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveStateNotifyAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// SetReportStatus is ReportStatus Setter
// 体检状态：未到检(exam_not), 已到检(exam_done)； 上门服务中还需以下两种状态：预约确认中（reserve_confirming），预约拒绝（reserve_rejected）；
func (r *AlibabaAlihealthExaminationReserveStateNotifyAPIRequest) SetReportStatus(_reportStatus string) error {
	r._reportStatus = _reportStatus
	r.Set("report_status", _reportStatus)
	return nil
}

// GetReportStatus ReportStatus Getter
func (r AlibabaAlihealthExaminationReserveStateNotifyAPIRequest) GetReportStatus() string {
	return r._reportStatus
}

// SetCheckNo is CheckNo Setter
// 到检凭证，exam_done状态下，该字段必填
func (r *AlibabaAlihealthExaminationReserveStateNotifyAPIRequest) SetCheckNo(_checkNo string) error {
	r._checkNo = _checkNo
	r.Set("check_no", _checkNo)
	return nil
}

// GetCheckNo CheckNo Getter
func (r AlibabaAlihealthExaminationReserveStateNotifyAPIRequest) GetCheckNo() string {
	return r._checkNo
}
