package gameact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取活动信息 APIRequest
taobao.de.activity.info.get

根据appKey和活动id获取活动
*/
type TaobaoDeActivityInfoGetRequest struct {
    model.Params

    // 事件唯一标识
    eventKey   string 

}

func NewTaobaoDeActivityInfoGetRequest() *TaobaoDeActivityInfoGetRequest{
    return &TaobaoDeActivityInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDeActivityInfoGetRequest) GetApiMethodName() string {
    return "taobao.de.activity.info.get"
}

func (r TaobaoDeActivityInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDeActivityInfoGetRequest) SetEventKey(eventKey string) error {
    r.eventKey = eventKey
    r.Set("event_key", eventKey)
    return nil
}

func (r TaobaoDeActivityInfoGetRequest) GetEventKey() string {
    return r.eventKey
}

