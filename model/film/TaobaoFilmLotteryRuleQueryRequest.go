package film

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘票票抽奖活动查询API(渠道) API请求
taobao.film.lottery.rule.query

淘票票抽奖活动查询API，渠道维度查询
*/
type TaobaoFilmLotteryRuleQueryRequest struct {
    model.Params
    // 账号类型（TAOBAO\ALIPAY\PHONE）
    _accountType   string
    // 渠道来源
    _channel   string
    // 账号ID
    _accountNo   string
}

// 初始化TaobaoFilmLotteryRuleQueryRequest对象
func NewTaobaoFilmLotteryRuleQueryRequest() *TaobaoFilmLotteryRuleQueryRequest{
    return &TaobaoFilmLotteryRuleQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFilmLotteryRuleQueryRequest) GetApiMethodName() string {
    return "taobao.film.lottery.rule.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFilmLotteryRuleQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountType Setter
// 账号类型（TAOBAO\ALIPAY\PHONE）
func (r *TaobaoFilmLotteryRuleQueryRequest) SetAccountType(_accountType string) error {
    r._accountType = _accountType
    r.Set("account_type", _accountType)
    return nil
}

// AccountType Getter
func (r TaobaoFilmLotteryRuleQueryRequest) GetAccountType() string {
    return r._accountType
}
// Channel Setter
// 渠道来源
func (r *TaobaoFilmLotteryRuleQueryRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TaobaoFilmLotteryRuleQueryRequest) GetChannel() string {
    return r._channel
}
// AccountNo Setter
// 账号ID
func (r *TaobaoFilmLotteryRuleQueryRequest) SetAccountNo(_accountNo string) error {
    r._accountNo = _accountNo
    r.Set("account_no", _accountNo)
    return nil
}

// AccountNo Getter
func (r TaobaoFilmLotteryRuleQueryRequest) GetAccountNo() string {
    return r._accountNo
}
