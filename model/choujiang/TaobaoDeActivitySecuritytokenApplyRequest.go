package choujiang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
安全token获取 APIRequest
taobao.de.activity.securitytoken.apply

新增接口，这个接口是用于在手机端进行抽奖时候的验证使用
*/
type TaobaoDeActivitySecuritytokenApplyRequest struct {
    model.Params

    // 运营和cp约定的事件唯一标示
    eventKey   string 

}

func NewTaobaoDeActivitySecuritytokenApplyRequest() *TaobaoDeActivitySecuritytokenApplyRequest{
    return &TaobaoDeActivitySecuritytokenApplyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDeActivitySecuritytokenApplyRequest) GetApiMethodName() string {
    return "taobao.de.activity.securitytoken.apply"
}

func (r TaobaoDeActivitySecuritytokenApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDeActivitySecuritytokenApplyRequest) SetEventKey(eventKey string) error {
    r.eventKey = eventKey
    r.Set("event_key", eventKey)
    return nil
}

func (r TaobaoDeActivitySecuritytokenApplyRequest) GetEventKey() string {
    return r.eventKey
}

