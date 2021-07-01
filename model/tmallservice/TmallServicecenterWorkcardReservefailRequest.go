package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预约失败 API请求
tmall.servicecenter.workcard.reservefail

服务商调用该接口回传工单预约失败
*/
type TmallServicecenterWorkcardReservefailAPIRequest struct {
    model.Params
    // 核销单外部id
    _identifyTaskId   int64
    // 下次联系时间
    _gmtNextContact   string
    // 预约失败原因码
    _failCode   int64
    // 预约失败原因描述
    _failDesc   string
    // 工单id
    _workcardId   int64
}

// 初始化TmallServicecenterWorkcardReservefailAPIRequest对象
func NewTmallServicecenterWorkcardReservefailRequest() *TmallServicecenterWorkcardReservefailAPIRequest{
    return &TmallServicecenterWorkcardReservefailAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.reservefail"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IdentifyTaskId Setter
// 核销单外部id
func (r *TmallServicecenterWorkcardReservefailAPIRequest) SetIdentifyTaskId(_identifyTaskId int64) error {
    r._identifyTaskId = _identifyTaskId
    r.Set("identify_task_id", _identifyTaskId)
    return nil
}

// IdentifyTaskId Getter
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetIdentifyTaskId() int64 {
    return r._identifyTaskId
}
// GmtNextContact Setter
// 下次联系时间
func (r *TmallServicecenterWorkcardReservefailAPIRequest) SetGmtNextContact(_gmtNextContact string) error {
    r._gmtNextContact = _gmtNextContact
    r.Set("gmt_next_contact", _gmtNextContact)
    return nil
}

// GmtNextContact Getter
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetGmtNextContact() string {
    return r._gmtNextContact
}
// FailCode Setter
// 预约失败原因码
func (r *TmallServicecenterWorkcardReservefailAPIRequest) SetFailCode(_failCode int64) error {
    r._failCode = _failCode
    r.Set("fail_code", _failCode)
    return nil
}

// FailCode Getter
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetFailCode() int64 {
    return r._failCode
}
// FailDesc Setter
// 预约失败原因描述
func (r *TmallServicecenterWorkcardReservefailAPIRequest) SetFailDesc(_failDesc string) error {
    r._failDesc = _failDesc
    r.Set("fail_desc", _failDesc)
    return nil
}

// FailDesc Getter
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetFailDesc() string {
    return r._failDesc
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardReservefailAPIRequest) SetWorkcardId(_workcardId int64) error {
    r._workcardId = _workcardId
    r.Set("workcard_id", _workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardReservefailAPIRequest) GetWorkcardId() int64 {
    return r._workcardId
}
