package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据手机号码领券 API请求
alibaba.mj.moscarnival.receivecoupon

根据手机号码领券
*/
type AlibabaMjMoscarnivalReceivecouponRequest struct {
    model.Params
    // 手机号码
    mobile   string
    // 活动id
    activityId   int64
}

// 初始化AlibabaMjMoscarnivalReceivecouponRequest对象
func NewAlibabaMjMoscarnivalReceivecouponRequest() *AlibabaMjMoscarnivalReceivecouponRequest{
    return &AlibabaMjMoscarnivalReceivecouponRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjMoscarnivalReceivecouponRequest) GetApiMethodName() string {
    return "alibaba.mj.moscarnival.receivecoupon"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjMoscarnivalReceivecouponRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mobile Setter
// 手机号码
func (r *AlibabaMjMoscarnivalReceivecouponRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r AlibabaMjMoscarnivalReceivecouponRequest) GetMobile() string {
    return r.mobile
}
// ActivityId Setter
// 活动id
func (r *AlibabaMjMoscarnivalReceivecouponRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r AlibabaMjMoscarnivalReceivecouponRequest) GetActivityId() int64 {
    return r.activityId
}
