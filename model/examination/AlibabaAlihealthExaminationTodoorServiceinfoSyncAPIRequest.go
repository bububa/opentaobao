package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest 上门检测服务信息同步 API请求
// alibaba.alihealth.examination.todoor.serviceinfo.sync
//
// isv同步上门检测服务信息给健康
type AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest struct {
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

// NewAlibabaAlihealthExaminationTodoorServiceinfoSyncRequest 初始化AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest对象
func NewAlibabaAlihealthExaminationTodoorServiceinfoSyncRequest() *AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest {
	return &AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.todoor.serviceinfo.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUniqReserveCode is UniqReserveCode Setter
// 服务商预约凭证
func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
	r._uniqReserveCode = _uniqReserveCode
	r.Set("uniq_reserve_code", _uniqReserveCode)
	return nil
}

// GetUniqReserveCode UniqReserveCode Getter
func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest) GetUniqReserveCode() string {
	return r._uniqReserveCode
}

// SetEventOccurTime is EventOccurTime Setter
// 事件发生时间
func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest) SetEventOccurTime(_eventOccurTime string) error {
	r._eventOccurTime = _eventOccurTime
	r.Set("event_occur_time", _eventOccurTime)
	return nil
}

// GetEventOccurTime EventOccurTime Getter
func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest) GetEventOccurTime() string {
	return r._eventOccurTime
}

// SetEvent is Event Setter
// 事件(ASSIGNED_PRACTITONER:已分配医护人员、PRACTITONER_GO_OUT:医护人员已出发、PRACTITONER_HOME:医护人员已到家、PRACTITONER_CHECKED:医护人员检查完成)、CHANGE_PRACTITONER(变更医护人员)
func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest) SetEvent(_event string) error {
	r._event = _event
	r.Set("event", _event)
	return nil
}

// GetEvent Event Getter
func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest) GetEvent() string {
	return r._event
}

// SetReserveNumber is ReserveNumber Setter
// 健康预约凭证
func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// GetReserveNumber ReserveNumber Getter
func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// SetMedicalPractitionerInfo is MedicalPractitionerInfo Setter
// 从业者信息
func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest) SetMedicalPractitionerInfo(_medicalPractitionerInfo *MedicalPractitionerInfo) error {
	r._medicalPractitionerInfo = _medicalPractitionerInfo
	r.Set("medical_practitioner_info", _medicalPractitionerInfo)
	return nil
}

// GetMedicalPractitionerInfo MedicalPractitionerInfo Getter
func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest) GetMedicalPractitionerInfo() *MedicalPractitionerInfo {
	return r._medicalPractitionerInfo
}
