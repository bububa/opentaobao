package degoperation

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通用用户抽奖次数限制 APIRequest
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

func NewTaobaoDegoperationGetByEventkeyRequest() *TaobaoDegoperationGetByEventkeyRequest{
    return &TaobaoDegoperationGetByEventkeyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDegoperationGetByEventkeyRequest) GetApiMethodName() string {
    return "taobao.degoperation.get.by.eventkey"
}

func (r TaobaoDegoperationGetByEventkeyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDegoperationGetByEventkeyRequest) SetDegAppKey(degAppKey string) error {
    r.degAppKey = degAppKey
    r.Set("deg_app_key", degAppKey)
    return nil
}

func (r TaobaoDegoperationGetByEventkeyRequest) GetDegAppKey() string {
    return r.degAppKey
}

func (r *TaobaoDegoperationGetByEventkeyRequest) SetEventKey(eventKey string) error {
    r.eventKey = eventKey
    r.Set("event_key", eventKey)
    return nil
}

func (r TaobaoDegoperationGetByEventkeyRequest) GetEventKey() string {
    return r.eventKey
}

func (r *TaobaoDegoperationGetByEventkeyRequest) SetDegAccessToken(degAccessToken string) error {
    r.degAccessToken = degAccessToken
    r.Set("deg_access_token", degAccessToken)
    return nil
}

func (r TaobaoDegoperationGetByEventkeyRequest) GetDegAccessToken() string {
    return r.degAccessToken
}

