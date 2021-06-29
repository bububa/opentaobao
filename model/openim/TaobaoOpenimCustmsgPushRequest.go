package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推送自定义openim消息 API请求
taobao.openim.custmsg.push

isv通过该接口给openim用户推送自定义消息
*/
type TaobaoOpenimCustmsgPushRequest struct {
    model.Params
    // 自定义消息内容
    _custmsg   *CustMsg
}

// 初始化TaobaoOpenimCustmsgPushRequest对象
func NewTaobaoOpenimCustmsgPushRequest() *TaobaoOpenimCustmsgPushRequest{
    return &TaobaoOpenimCustmsgPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimCustmsgPushRequest) GetApiMethodName() string {
    return "taobao.openim.custmsg.push"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimCustmsgPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Custmsg Setter
// 自定义消息内容
func (r *TaobaoOpenimCustmsgPushRequest) SetCustmsg(_custmsg *CustMsg) error {
    r._custmsg = _custmsg
    r.Set("custmsg", _custmsg)
    return nil
}

// Custmsg Getter
func (r TaobaoOpenimCustmsgPushRequest) GetCustmsg() *CustMsg {
    return r._custmsg
}
