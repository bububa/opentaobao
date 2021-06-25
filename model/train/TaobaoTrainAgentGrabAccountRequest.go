package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代购抢代理商回传12306账号 APIRequest
taobao.train.agent.grab.account

火车票业务代购抢功能，代理商回传12306账号，用于自营抢票链路出票
*/
type TaobaoTrainAgentGrabAccountRequest struct {
    model.Params

    // 12306账户信息
    accountParam   *AccountParam 

}

func NewTaobaoTrainAgentGrabAccountRequest() *TaobaoTrainAgentGrabAccountRequest{
    return &TaobaoTrainAgentGrabAccountRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentGrabAccountRequest) GetApiMethodName() string {
    return "taobao.train.agent.grab.account"
}

func (r TaobaoTrainAgentGrabAccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentGrabAccountRequest) SetAccountParam(accountParam *AccountParam) error {
    r.accountParam = accountParam
    r.Set("account_param", accountParam)
    return nil
}

func (r TaobaoTrainAgentGrabAccountRequest) GetAccountParam() *AccountParam {
    return r.accountParam
}

