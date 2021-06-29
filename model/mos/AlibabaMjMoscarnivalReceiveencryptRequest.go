package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据加密手机号领券 API请求
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

// 初始化AlibabaMjMoscarnivalReceiveencryptRequest对象
func NewAlibabaMjMoscarnivalReceiveencryptRequest() *AlibabaMjMoscarnivalReceiveencryptRequest{
    return &AlibabaMjMoscarnivalReceiveencryptRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjMoscarnivalReceiveencryptRequest) GetApiMethodName() string {
    return "alibaba.mj.moscarnival.receiveencrypt"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjMoscarnivalReceiveencryptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mobile Setter
// 加密手机号码
func (r *AlibabaMjMoscarnivalReceiveencryptRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r AlibabaMjMoscarnivalReceiveencryptRequest) GetMobile() string {
    return r.mobile
}
// ActivityId Setter
// 活动id
func (r *AlibabaMjMoscarnivalReceiveencryptRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r AlibabaMjMoscarnivalReceiveencryptRequest) GetActivityId() int64 {
    return r.activityId
}
