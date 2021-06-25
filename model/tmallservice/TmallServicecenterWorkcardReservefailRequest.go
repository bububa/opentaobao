package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预约失败 APIRequest
tmall.servicecenter.workcard.reservefail

服务商调用该接口回传工单预约失败
*/
type TmallServicecenterWorkcardReservefailRequest struct {
    model.Params

    // 核销单外部id
    identifyTaskId   int64 

    // 下次联系时间
    gmtNextContact   string 

    // 预约失败原因码
    failCode   int64 

    // 预约失败原因描述
    failDesc   string 

    // 工单id
    workcardId   int64 

}

func NewTmallServicecenterWorkcardReservefailRequest() *TmallServicecenterWorkcardReservefailRequest{
    return &TmallServicecenterWorkcardReservefailRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardReservefailRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.reservefail"
}

func (r TmallServicecenterWorkcardReservefailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardReservefailRequest) SetIdentifyTaskId(identifyTaskId int64) error {
    r.identifyTaskId = identifyTaskId
    r.Set("identify_task_id", identifyTaskId)
    return nil
}

func (r TmallServicecenterWorkcardReservefailRequest) GetIdentifyTaskId() int64 {
    return r.identifyTaskId
}

func (r *TmallServicecenterWorkcardReservefailRequest) SetGmtNextContact(gmtNextContact string) error {
    r.gmtNextContact = gmtNextContact
    r.Set("gmt_next_contact", gmtNextContact)
    return nil
}

func (r TmallServicecenterWorkcardReservefailRequest) GetGmtNextContact() string {
    return r.gmtNextContact
}

func (r *TmallServicecenterWorkcardReservefailRequest) SetFailCode(failCode int64) error {
    r.failCode = failCode
    r.Set("fail_code", failCode)
    return nil
}

func (r TmallServicecenterWorkcardReservefailRequest) GetFailCode() int64 {
    return r.failCode
}

func (r *TmallServicecenterWorkcardReservefailRequest) SetFailDesc(failDesc string) error {
    r.failDesc = failDesc
    r.Set("fail_desc", failDesc)
    return nil
}

func (r TmallServicecenterWorkcardReservefailRequest) GetFailDesc() string {
    return r.failDesc
}

func (r *TmallServicecenterWorkcardReservefailRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

func (r TmallServicecenterWorkcardReservefailRequest) GetWorkcardId() int64 {
    return r.workcardId
}

