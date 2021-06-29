package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上门检测服务信息同步 API请求
alibaba.alihealth.examination.todoor.serviceinfo.sync

isv同步上门检测服务信息给健康
*/
type AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest struct {
    model.Params
    // 服务商预约凭证
    uniqReserveCode   string
    // 从业者信息
    medicalPractitionerInfo   *MedicalPractitionerInfo
    // 健康预约凭证
    reserveNumber   string
    // 事件(ASSIGNED_PRACTITONER:已分配医护人员、PRACTITONER_GO_OUT:医护人员已出发、PRACTITONER_HOME:医护人员已到家、PRACTITONER_CHECKED:医护人员检查完成)、CHANGE_PRACTITONER(变更医护人员)
    event   string
    // 事件发生时间
    eventOccurTime   string
}

// 初始化AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest对象
func NewAlibabaAlihealthExaminationTodoorServiceinfoSyncRequest() *AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest{
    return &AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.todoor.serviceinfo.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UniqReserveCode Setter
// 服务商预约凭证
func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) SetUniqReserveCode(uniqReserveCode string) error {
    r.uniqReserveCode = uniqReserveCode
    r.Set("uniq_reserve_code", uniqReserveCode)
    return nil
}

// UniqReserveCode Getter
func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) GetUniqReserveCode() string {
    return r.uniqReserveCode
}
// MedicalPractitionerInfo Setter
// 从业者信息
func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) SetMedicalPractitionerInfo(medicalPractitionerInfo *MedicalPractitionerInfo) error {
    r.medicalPractitionerInfo = medicalPractitionerInfo
    r.Set("medical_practitioner_info", medicalPractitionerInfo)
    return nil
}

// MedicalPractitionerInfo Getter
func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) GetMedicalPractitionerInfo() *MedicalPractitionerInfo {
    return r.medicalPractitionerInfo
}
// ReserveNumber Setter
// 健康预约凭证
func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) SetReserveNumber(reserveNumber string) error {
    r.reserveNumber = reserveNumber
    r.Set("reserve_number", reserveNumber)
    return nil
}

// ReserveNumber Getter
func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) GetReserveNumber() string {
    return r.reserveNumber
}
// Event Setter
// 事件(ASSIGNED_PRACTITONER:已分配医护人员、PRACTITONER_GO_OUT:医护人员已出发、PRACTITONER_HOME:医护人员已到家、PRACTITONER_CHECKED:医护人员检查完成)、CHANGE_PRACTITONER(变更医护人员)
func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) SetEvent(event string) error {
    r.event = event
    r.Set("event", event)
    return nil
}

// Event Getter
func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) GetEvent() string {
    return r.event
}
// EventOccurTime Setter
// 事件发生时间
func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) SetEventOccurTime(eventOccurTime string) error {
    r.eventOccurTime = eventOccurTime
    r.Set("event_occur_time", eventOccurTime)
    return nil
}

// EventOccurTime Getter
func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) GetEventOccurTime() string {
    return r.eventOccurTime
}
