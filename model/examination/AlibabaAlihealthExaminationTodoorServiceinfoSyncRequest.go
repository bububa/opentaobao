package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上门检测服务信息同步 APIRequest
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

func NewAlibabaAlihealthExaminationTodoorServiceinfoSyncRequest() *AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest{
    return &AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.todoor.serviceinfo.sync"
}

func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) SetUniqReserveCode(uniqReserveCode string) error {
    r.uniqReserveCode = uniqReserveCode
    r.Set("uniq_reserve_code", uniqReserveCode)
    return nil
}

func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) GetUniqReserveCode() string {
    return r.uniqReserveCode
}

func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) SetMedicalPractitionerInfo(medicalPractitionerInfo *MedicalPractitionerInfo) error {
    r.medicalPractitionerInfo = medicalPractitionerInfo
    r.Set("medical_practitioner_info", medicalPractitionerInfo)
    return nil
}

func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) GetMedicalPractitionerInfo() *MedicalPractitionerInfo {
    return r.medicalPractitionerInfo
}

func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) SetReserveNumber(reserveNumber string) error {
    r.reserveNumber = reserveNumber
    r.Set("reserve_number", reserveNumber)
    return nil
}

func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) GetReserveNumber() string {
    return r.reserveNumber
}

func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) SetEvent(event string) error {
    r.event = event
    r.Set("event", event)
    return nil
}

func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) GetEvent() string {
    return r.event
}

func (r *AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) SetEventOccurTime(eventOccurTime string) error {
    r.eventOccurTime = eventOccurTime
    r.Set("event_occur_time", eventOccurTime)
    return nil
}

func (r AlibabaAlihealthExaminationTodoorServiceinfoSyncRequest) GetEventOccurTime() string {
    return r.eventOccurTime
}

