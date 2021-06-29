package gameact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取活动信息 API请求
taobao.de.activity.info.get

根据appKey和活动id获取活动
*/
type TaobaoDeActivityInfoGetRequest struct {
    model.Params
    // 事件唯一标识
    eventKey   string
}

// 初始化TaobaoDeActivityInfoGetRequest对象
func NewTaobaoDeActivityInfoGetRequest() *TaobaoDeActivityInfoGetRequest{
    return &TaobaoDeActivityInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDeActivityInfoGetRequest) GetApiMethodName() string {
    return "taobao.de.activity.info.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDeActivityInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EventKey Setter
// 事件唯一标识
func (r *TaobaoDeActivityInfoGetRequest) SetEventKey(eventKey string) error {
    r.eventKey = eventKey
    r.Set("event_key", eventKey)
    return nil
}

// EventKey Getter
func (r TaobaoDeActivityInfoGetRequest) GetEventKey() string {
    return r.eventKey
}
