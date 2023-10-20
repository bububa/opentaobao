package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest 上门检测服务信息同步 API请求
// alibaba.alihealth.examination.todoor.serviceinfo.sync
//
// isv同步上门检测服务信息给健康
type AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest struct {
	model.Params
	// 服务商预约凭证
	_uniqReserveCode string
	// 事件发生时间
	_eventOccurTime string
	// 事件(ASSIGNED_PRACTITONER:已分配医护人员、PRACTITONER_GO_OUT:医护人员已出发、PRACTITONER_HOME:医护人员已到家、PRACTITONER_CHECKED:医护人员检查完成)、CHANGE_PRACTITONER(变更医护人员)
	_event string
	// 健康预约凭证
	_reserveNumber string
	// 从业者信息
	_medicalPractitionerInfo *MedicalPractitionerInfo
}

// NewAlibabaalihealthexaminationtodoorserviceinfosyncRequest 初始化AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest对象
func NewAlibabaalihealthexaminationtodoorserviceinfosyncRequest() *AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest {
	return &AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.todoor.serviceinfo.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUniqReserveCode is UniqReserveCode Setter
// 服务商预约凭证
func (r *AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
	r._uniqReserveCode = _uniqReserveCode
	r.Set("uniq_reserve_code", _uniqReserveCode)
	return nil
}

// GetUniqReserveCode UniqReserveCode Getter
func (r AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest) GetUniqReserveCode() string {
	return r._uniqReserveCode
}

// SetEventOccurTime is EventOccurTime Setter
// 事件发生时间
func (r *AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest) SetEventOccurTime(_eventOccurTime string) error {
	r._eventOccurTime = _eventOccurTime
	r.Set("event_occur_time", _eventOccurTime)
	return nil
}

// GetEventOccurTime EventOccurTime Getter
func (r AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest) GetEventOccurTime() string {
	return r._eventOccurTime
}

// SetEvent is Event Setter
// 事件(ASSIGNED_PRACTITONER:已分配医护人员、PRACTITONER_GO_OUT:医护人员已出发、PRACTITONER_HOME:医护人员已到家、PRACTITONER_CHECKED:医护人员检查完成)、CHANGE_PRACTITONER(变更医护人员)
func (r *AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest) SetEvent(_event string) error {
	r._event = _event
	r.Set("event", _event)
	return nil
}

// GetEvent Event Getter
func (r AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest) GetEvent() string {
	return r._event
}

// SetReserveNumber is ReserveNumber Setter
// 健康预约凭证
func (r *AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// GetReserveNumber ReserveNumber Getter
func (r AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// SetMedicalPractitionerInfo is MedicalPractitionerInfo Setter
// 从业者信息
func (r *AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest) SetMedicalPractitionerInfo(_medicalPractitionerInfo *MedicalPractitionerInfo) error {
	r._medicalPractitionerInfo = _medicalPractitionerInfo
	r.Set("medical_practitioner_info", _medicalPractitionerInfo)
	return nil
}

// GetMedicalPractitionerInfo MedicalPractitionerInfo Getter
func (r AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest) GetMedicalPractitionerInfo() *MedicalPractitionerInfo {
	return r._medicalPractitionerInfo
}
