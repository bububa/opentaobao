package degoperation

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通用用户抽奖次数限制 API请求
taobao.degoperation.get.by.eventkey

通用用户抽奖次数限制
*/
type TaobaoDegoperationGetByEventkeyRequest struct {
    model.Params
    // 活动后台配置appkey
    degAppKey   string
    // 活动后台配置eventkey
    eventKey   string
    // info
    degAccessToken   string
}

// 初始化TaobaoDegoperationGetByEventkeyRequest对象
func NewTaobaoDegoperationGetByEventkeyRequest() *TaobaoDegoperationGetByEventkeyRequest{
    return &TaobaoDegoperationGetByEventkeyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationGetByEventkeyRequest) GetApiMethodName() string {
    return "taobao.degoperation.get.by.eventkey"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationGetByEventkeyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DegAppKey Setter
// 活动后台配置appkey
func (r *TaobaoDegoperationGetByEventkeyRequest) SetDegAppKey(degAppKey string) error {
    r.degAppKey = degAppKey
    r.Set("deg_app_key", degAppKey)
    return nil
}

// DegAppKey Getter
func (r TaobaoDegoperationGetByEventkeyRequest) GetDegAppKey() string {
    return r.degAppKey
}
// EventKey Setter
// 活动后台配置eventkey
func (r *TaobaoDegoperationGetByEventkeyRequest) SetEventKey(eventKey string) error {
    r.eventKey = eventKey
    r.Set("event_key", eventKey)
    return nil
}

// EventKey Getter
func (r TaobaoDegoperationGetByEventkeyRequest) GetEventKey() string {
    return r.eventKey
}
// DegAccessToken Setter
// info
func (r *TaobaoDegoperationGetByEventkeyRequest) SetDegAccessToken(degAccessToken string) error {
    r.degAccessToken = degAccessToken
    r.Set("deg_access_token", degAccessToken)
    return nil
}

// DegAccessToken Getter
func (r TaobaoDegoperationGetByEventkeyRequest) GetDegAccessToken() string {
    return r.degAccessToken
}
