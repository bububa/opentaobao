package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据加密手机号领券 APIRequest
alibaba.mj.moscarnival.receiveencrypt

根据加密手机号领券
*/
type AlibabaMjMoscarnivalReceiveencryptRequest struct {
    model.Params

    // 加密手机号码
    mobile   string 

    // 活动id
    activityId   int64 

}

func NewAlibabaMjMoscarnivalReceiveencryptRequest() *AlibabaMjMoscarnivalReceiveencryptRequest{
    return &AlibabaMjMoscarnivalReceiveencryptRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjMoscarnivalReceiveencryptRequest) GetApiMethodName() string {
    return "alibaba.mj.moscarnival.receiveencrypt"
}

func (r AlibabaMjMoscarnivalReceiveencryptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjMoscarnivalReceiveencryptRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r AlibabaMjMoscarnivalReceiveencryptRequest) GetMobile() string {
    return r.mobile
}

func (r *AlibabaMjMoscarnivalReceiveencryptRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r AlibabaMjMoscarnivalReceiveencryptRequest) GetActivityId() int64 {
    return r.activityId
}

