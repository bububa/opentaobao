package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代购抢代理商回传12306账号 API请求
taobao.train.agent.grab.account

火车票业务代购抢功能，代理商回传12306账号，用于自营抢票链路出票
*/
type TaobaoTrainAgentGrabAccountRequest struct {
    model.Params
    // 12306账户信息
    _accountParam   *AccountParam
}

// 初始化TaobaoTrainAgentGrabAccountRequest对象
func NewTaobaoTrainAgentGrabAccountRequest() *TaobaoTrainAgentGrabAccountRequest{
    return &TaobaoTrainAgentGrabAccountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentGrabAccountRequest) GetApiMethodName() string {
    return "taobao.train.agent.grab.account"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentGrabAccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountParam Setter
// 12306账户信息
func (r *TaobaoTrainAgentGrabAccountRequest) SetAccountParam(_accountParam *AccountParam) error {
    r._accountParam = _accountParam
    r.Set("account_param", _accountParam)
    return nil
}

// AccountParam Getter
func (r TaobaoTrainAgentGrabAccountRequest) GetAccountParam() *AccountParam {
    return r._accountParam
}
