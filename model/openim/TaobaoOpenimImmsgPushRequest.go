package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openim标准消息发送 APIRequest
taobao.openim.immsg.push

服务端对openim用户发送标准消息，包括文字、语音、图片等。
*/
type TaobaoOpenimImmsgPushRequest struct {
    model.Params

    // openim消息结构体
    immsg   *ImMsg 

}

func NewTaobaoOpenimImmsgPushRequest() *TaobaoOpenimImmsgPushRequest{
    return &TaobaoOpenimImmsgPushRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimImmsgPushRequest) GetApiMethodName() string {
    return "taobao.openim.immsg.push"
}

func (r TaobaoOpenimImmsgPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimImmsgPushRequest) SetImmsg(immsg *ImMsg) error {
    r.immsg = immsg
    r.Set("immsg", immsg)
    return nil
}

func (r TaobaoOpenimImmsgPushRequest) GetImmsg() *ImMsg {
    return r.immsg
}

