package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单挂起 API请求
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

// 初始化TmallServicecenterWorkcardSuspendRequest对象
func NewTmallServicecenterWorkcardSuspendRequest() *TmallServicecenterWorkcardSuspendRequest{
    return &TmallServicecenterWorkcardSuspendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardSuspendRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.suspend"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardSuspendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReserveServiceDate Setter
// 预约时间
func (r *TmallServicecenterWorkcardSuspendRequest) SetReserveServiceDate(reserveServiceDate string) error {
    r.reserveServiceDate = reserveServiceDate
    r.Set("reserve_service_date", reserveServiceDate)
    return nil
}

// ReserveServiceDate Getter
func (r TmallServicecenterWorkcardSuspendRequest) GetReserveServiceDate() string {
    return r.reserveServiceDate
}
// GmtNextContact Setter
// 下次联系时间
func (r *TmallServicecenterWorkcardSuspendRequest) SetGmtNextContact(gmtNextContact string) error {
    r.gmtNextContact = gmtNextContact
    r.Set("gmt_next_contact", gmtNextContact)
    return nil
}

// GmtNextContact Getter
func (r TmallServicecenterWorkcardSuspendRequest) GetGmtNextContact() string {
    return r.gmtNextContact
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardSuspendRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardSuspendRequest) GetWorkcardId() int64 {
    return r.workcardId
}
// FailCode Setter
// 挂起原因类型code
func (r *TmallServicecenterWorkcardSuspendRequest) SetFailCode(failCode int64) error {
    r.failCode = failCode
    r.Set("fail_code", failCode)
    return nil
}

// FailCode Getter
func (r TmallServicecenterWorkcardSuspendRequest) GetFailCode() int64 {
    return r.failCode
}
// FailDesc Setter
// 挂起原因描述
func (r *TmallServicecenterWorkcardSuspendRequest) SetFailDesc(failDesc string) error {
    r.failDesc = failDesc
    r.Set("fail_desc", failDesc)
    return nil
}

// FailDesc Getter
func (r TmallServicecenterWorkcardSuspendRequest) GetFailDesc() string {
    return r.failDesc
}
