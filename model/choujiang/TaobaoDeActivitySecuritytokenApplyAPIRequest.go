package choujiang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
安全token获取 API请求
taobao.de.activity.securitytoken.apply

新增接口，这个接口是用于在手机端进行抽奖时候的验证使用
*/
type TaobaoDeActivitySecuritytokenApplyAPIRequest struct {
    model.Params
    // 运营和cp约定的事件唯一标示
    _eventKey   string
}

// 初始化TaobaoDeActivitySecuritytokenApplyAPIRequest对象
func NewTaobaoDeActivitySecuritytokenApplyRequest() *TaobaoDeActivitySecuritytokenApplyAPIRequest{
    return &TaobaoDeActivitySecuritytokenApplyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDeActivitySecuritytokenApplyAPIRequest) GetApiMethodName() string {
    return "taobao.de.activity.securitytoken.apply"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDeActivitySecuritytokenApplyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EventKey Setter
// 运营和cp约定的事件唯一标示
func (r *TaobaoDeActivitySecuritytokenApplyAPIRequest) SetEventKey(_eventKey string) error {
    r._eventKey = _eventKey
    r.Set("event_key", _eventKey)
    return nil
}

// EventKey Getter
func (r TaobaoDeActivitySecuritytokenApplyAPIRequest) GetEventKey() string {
    return r._eventKey
}
