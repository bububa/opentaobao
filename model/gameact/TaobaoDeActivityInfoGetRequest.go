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
type TaobaoDeActivityInfoGetAPIRequest struct {
    model.Params
    // 事件唯一标识
    _eventKey   string
}

// 初始化TaobaoDeActivityInfoGetAPIRequest对象
func NewTaobaoDeActivityInfoGetRequest() *TaobaoDeActivityInfoGetAPIRequest{
    return &TaobaoDeActivityInfoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDeActivityInfoGetAPIRequest) GetApiMethodName() string {
    return "taobao.de.activity.info.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDeActivityInfoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EventKey Setter
// 事件唯一标识
func (r *TaobaoDeActivityInfoGetAPIRequest) SetEventKey(_eventKey string) error {
    r._eventKey = _eventKey
    r.Set("event_key", _eventKey)
    return nil
}

// EventKey Getter
func (r TaobaoDeActivityInfoGetAPIRequest) GetEventKey() string {
    return r._eventKey
}
