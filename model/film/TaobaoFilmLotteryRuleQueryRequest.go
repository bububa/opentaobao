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
    accountType   string
    // 渠道来源
    channel   string
    // 账号ID
    accountNo   string
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
func (r *TaobaoFilmLotteryRuleQueryRequest) SetAccountType(accountType string) error {
    r.accountType = accountType
    r.Set("account_type", accountType)
    return nil
}

// AccountType Getter
func (r TaobaoFilmLotteryRuleQueryRequest) GetAccountType() string {
    return r.accountType
}
// Channel Setter
// 渠道来源
func (r *TaobaoFilmLotteryRuleQueryRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

// Channel Getter
func (r TaobaoFilmLotteryRuleQueryRequest) GetChannel() string {
    return r.channel
}
// AccountNo Setter
// 账号ID
func (r *TaobaoFilmLotteryRuleQueryRequest) SetAccountNo(accountNo string) error {
    r.accountNo = accountNo
    r.Set("account_no", accountNo)
    return nil
}

// AccountNo Getter
func (r TaobaoFilmLotteryRuleQueryRequest) GetAccountNo() string {
    return r.accountNo
}
