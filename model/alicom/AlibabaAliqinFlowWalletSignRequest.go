package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
流量平台用户签约情况查询 APIRequest
alibaba.aliqin.flow.wallet.sign

流量平台用户签约情况查询
*/
type AlibabaAliqinFlowWalletSignRequest struct {
    model.Params

    // 用户昵称
    userNick   string 

}

func NewAlibabaAliqinFlowWalletSignRequest() *AlibabaAliqinFlowWalletSignRequest{
    return &AlibabaAliqinFlowWalletSignRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFlowWalletSignRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.sign"
}

func (r AlibabaAliqinFlowWalletSignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFlowWalletSignRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r AlibabaAliqinFlowWalletSignRequest) GetUserNick() string {
    return r.userNick
}

