package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否已经登录 APIRequest
alibaba.interact.user.islogin

API的功能是校验用户是否登录，ISV调用接口的时候，通过此接口映射到mtop.interact.user.islogin接口上，因此接口只是做一个给ISV注册调用api的入口，没有发生真实的RPC
*/
type AlibabaInteractUserIsloginRequest struct {
    model.Params

    // 用户nick
    buyerNick   string 

}

func NewAlibabaInteractUserIsloginRequest() *AlibabaInteractUserIsloginRequest{
    return &AlibabaInteractUserIsloginRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractUserIsloginRequest) GetApiMethodName() string {
    return "alibaba.interact.user.islogin"
}

func (r AlibabaInteractUserIsloginRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractUserIsloginRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r AlibabaInteractUserIsloginRequest) GetBuyerNick() string {
    return r.buyerNick
}

