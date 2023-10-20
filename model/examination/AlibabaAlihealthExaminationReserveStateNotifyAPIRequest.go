package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreservestatenotifyAPIRequest 体检机构对接_体检状态主动通知 API请求
// alibaba.alihealth.examination.reserve.state.notify
//
// 到了体检当天后，服务商主动通知体检预约状态
type AlibabaalihealthexaminationreservestatenotifyAPIRequest struct {
	model.Params
	// 服务商预约凭证
	_uniqReserveCode string
	// 健康预约凭证
	_reserveNumber string
	// 体检状态：已到检(exam_done)、确认预约(exam_not)、取消预约(exam_cancel)；
	_reportStatus string
	// 到检凭证，exam_done状态下，该字段必填
	_checkNo string
}

// NewAlibabaalihealthexaminationreservestatenotifyRequest 初始化AlibabaalihealthexaminationreservestatenotifyAPIRequest对象
func NewAlibabaalihealthexaminationreservestatenotifyRequest() *AlibabaalihealthexaminationreservestatenotifyAPIRequest {
	return &AlibabaalihealthexaminationreservestatenotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationreservestatenotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.reserve.state.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationreservestatenotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationreservestatenotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUniqReserveCode is UniqReserveCode Setter
// 服务商预约凭证
func (r *AlibabaalihealthexaminationreservestatenotifyAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
	r._uniqReserveCode = _uniqReserveCode
	r.Set("uniq_reserve_code", _uniqReserveCode)
	return nil
}

// GetUniqReserveCode UniqReserveCode Getter
func (r AlibabaalihealthexaminationreservestatenotifyAPIRequest) GetUniqReserveCode() string {
	return r._uniqReserveCode
}

// SetReserveNumber is ReserveNumber Setter
// 健康预约凭证
func (r *AlibabaalihealthexaminationreservestatenotifyAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// GetReserveNumber ReserveNumber Getter
func (r AlibabaalihealthexaminationreservestatenotifyAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// SetReportStatus is ReportStatus Setter
// 体检状态：已到检(exam_done)、确认预约(exam_not)、取消预约(exam_cancel)；
func (r *AlibabaalihealthexaminationreservestatenotifyAPIRequest) SetReportStatus(_reportStatus string) error {
	r._reportStatus = _reportStatus
	r.Set("report_status", _reportStatus)
	return nil
}

// GetReportStatus ReportStatus Getter
func (r AlibabaalihealthexaminationreservestatenotifyAPIRequest) GetReportStatus() string {
	return r._reportStatus
}

// SetCheckNo is CheckNo Setter
// 到检凭证，exam_done状态下，该字段必填
func (r *AlibabaalihealthexaminationreservestatenotifyAPIRequest) SetCheckNo(_checkNo string) error {
	r._checkNo = _checkNo
	r.Set("check_no", _checkNo)
	return nil
}

// GetCheckNo CheckNo Getter
func (r AlibabaalihealthexaminationreservestatenotifyAPIRequest) GetCheckNo() string {
	return r._checkNo
}
