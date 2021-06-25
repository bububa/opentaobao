package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推送自定义openim消息 APIRequest
taobao.openim.custmsg.push

isv通过该接口给openim用户推送自定义消息
*/
type TaobaoOpenimCustmsgPushRequest struct {
    model.Params

    // 自定义消息内容
    custmsg   *CustMsg 

}

func NewTaobaoOpenimCustmsgPushRequest() *TaobaoOpenimCustmsgPushRequest{
    return &TaobaoOpenimCustmsgPushRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimCustmsgPushRequest) GetApiMethodName() string {
    return "taobao.openim.custmsg.push"
}

func (r TaobaoOpenimCustmsgPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimCustmsgPushRequest) SetCustmsg(custmsg *CustMsg) error {
    r.custmsg = custmsg
    r.Set("custmsg", custmsg)
    return nil
}

func (r TaobaoOpenimCustmsgPushRequest) GetCustmsg() *CustMsg {
    return r.custmsg
}

