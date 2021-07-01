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
type AlibabaMjMoscarnivalReceiveencryptAPIRequest struct {
    model.Params
    // 加密手机号码
    _mobile   string
    // 活动id
    _activityId   int64
}

// 初始化AlibabaMjMoscarnivalReceiveencryptAPIRequest对象
func NewAlibabaMjMoscarnivalReceiveencryptRequest() *AlibabaMjMoscarnivalReceiveencryptAPIRequest{
    return &AlibabaMjMoscarnivalReceiveencryptAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjMoscarnivalReceiveencryptAPIRequest) GetApiMethodName() string {
    return "alibaba.mj.moscarnival.receiveencrypt"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjMoscarnivalReceiveencryptAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mobile Setter
// 加密手机号码
func (r *AlibabaMjMoscarnivalReceiveencryptAPIRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r AlibabaMjMoscarnivalReceiveencryptAPIRequest) GetMobile() string {
    return r._mobile
}
// ActivityId Setter
// 活动id
func (r *AlibabaMjMoscarnivalReceiveencryptAPIRequest) SetActivityId(_activityId int64) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r AlibabaMjMoscarnivalReceiveencryptAPIRequest) GetActivityId() int64 {
    return r._activityId
}
