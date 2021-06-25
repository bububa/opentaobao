package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单挂起 APIRequest
tmall.servicecenter.workcard.suspend

工单挂起
*/
type TmallServicecenterWorkcardSuspendRequest struct {
    model.Params

    // 预约时间
    reserveServiceDate   string 

    // 下次联系时间
    gmtNextContact   string 

    // 工单id
    workcardId   int64 

    // 挂起原因类型code
    failCode   int64 

    // 挂起原因描述
    failDesc   string 

}

func NewTmallServicecenterWorkcardSuspendRequest() *TmallServicecenterWorkcardSuspendRequest{
    return &TmallServicecenterWorkcardSuspendRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardSuspendRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.suspend"
}

func (r TmallServicecenterWorkcardSuspendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardSuspendRequest) SetReserveServiceDate(reserveServiceDate string) error {
    r.reserveServiceDate = reserveServiceDate
    r.Set("reserve_service_date", reserveServiceDate)
    return nil
}

func (r TmallServicecenterWorkcardSuspendRequest) GetReserveServiceDate() string {
    return r.reserveServiceDate
}

func (r *TmallServicecenterWorkcardSuspendRequest) SetGmtNextContact(gmtNextContact string) error {
    r.gmtNextContact = gmtNextContact
    r.Set("gmt_next_contact", gmtNextContact)
    return nil
}

func (r TmallServicecenterWorkcardSuspendRequest) GetGmtNextContact() string {
    return r.gmtNextContact
}

func (r *TmallServicecenterWorkcardSuspendRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

func (r TmallServicecenterWorkcardSuspendRequest) GetWorkcardId() int64 {
    return r.workcardId
}

func (r *TmallServicecenterWorkcardSuspendRequest) SetFailCode(failCode int64) error {
    r.failCode = failCode
    r.Set("fail_code", failCode)
    return nil
}

func (r TmallServicecenterWorkcardSuspendRequest) GetFailCode() int64 {
    return r.failCode
}

func (r *TmallServicecenterWorkcardSuspendRequest) SetFailDesc(failDesc string) error {
    r.failDesc = failDesc
    r.Set("fail_desc", failDesc)
    return nil
}

func (r TmallServicecenterWorkcardSuspendRequest) GetFailDesc() string {
    return r.failDesc
}

