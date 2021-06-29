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
    _degAppKey   string
    // 活动后台配置eventkey
    _eventKey   string
    // info
    _degAccessToken   string
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
func (r *TaobaoDegoperationGetByEventkeyRequest) SetDegAppKey(_degAppKey string) error {
    r._degAppKey = _degAppKey
    r.Set("deg_app_key", _degAppKey)
    return nil
}

// DegAppKey Getter
func (r TaobaoDegoperationGetByEventkeyRequest) GetDegAppKey() string {
    return r._degAppKey
}
// EventKey Setter
// 活动后台配置eventkey
func (r *TaobaoDegoperationGetByEventkeyRequest) SetEventKey(_eventKey string) error {
    r._eventKey = _eventKey
    r.Set("event_key", _eventKey)
    return nil
}

// EventKey Getter
func (r TaobaoDegoperationGetByEventkeyRequest) GetEventKey() string {
    return r._eventKey
}
// DegAccessToken Setter
// info
func (r *TaobaoDegoperationGetByEventkeyRequest) SetDegAccessToken(_degAccessToken string) error {
    r._degAccessToken = _degAccessToken
    r.Set("deg_access_token", _degAccessToken)
    return nil
}

// DegAccessToken Getter
func (r TaobaoDegoperationGetByEventkeyRequest) GetDegAccessToken() string {
    return r._degAccessToken
}
