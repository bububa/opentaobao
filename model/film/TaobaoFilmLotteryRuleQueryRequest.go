package film

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘票票抽奖活动查询API(渠道) APIRequest
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

func NewTaobaoFilmLotteryRuleQueryRequest() *TaobaoFilmLotteryRuleQueryRequest{
    return &TaobaoFilmLotteryRuleQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFilmLotteryRuleQueryRequest) GetApiMethodName() string {
    return "taobao.film.lottery.rule.query"
}

func (r TaobaoFilmLotteryRuleQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFilmLotteryRuleQueryRequest) SetAccountType(accountType string) error {
    r.accountType = accountType
    r.Set("account_type", accountType)
    return nil
}

func (r TaobaoFilmLotteryRuleQueryRequest) GetAccountType() string {
    return r.accountType
}

func (r *TaobaoFilmLotteryRuleQueryRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r TaobaoFilmLotteryRuleQueryRequest) GetChannel() string {
    return r.channel
}

func (r *TaobaoFilmLotteryRuleQueryRequest) SetAccountNo(accountNo string) error {
    r.accountNo = accountNo
    r.Set("account_no", accountNo)
    return nil
}

func (r TaobaoFilmLotteryRuleQueryRequest) GetAccountNo() string {
    return r.accountNo
}

