package film

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘票票抽奖发放权益API APIRequest
taobao.film.lottery.draw

对外第三方合作渠道通过抽奖形式发码
*/
type TaobaoFilmLotteryDrawRequest struct {
    model.Params

    // 账号ID
    accountNo   string 

    // 账号类型（TAOBAO\ALIPAY\PHONE\TAOBAO_NAME\OPEN_ID）
    accountType   string 

    // 活动ID
    lotteryMixId   string 

    // 平台类型
    platform   int64 

    // 扩展参数
    bizData   string 

}

func NewTaobaoFilmLotteryDrawRequest() *TaobaoFilmLotteryDrawRequest{
    return &TaobaoFilmLotteryDrawRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFilmLotteryDrawRequest) GetApiMethodName() string {
    return "taobao.film.lottery.draw"
}

func (r TaobaoFilmLotteryDrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFilmLotteryDrawRequest) SetAccountNo(accountNo string) error {
    r.accountNo = accountNo
    r.Set("account_no", accountNo)
    return nil
}

func (r TaobaoFilmLotteryDrawRequest) GetAccountNo() string {
    return r.accountNo
}

func (r *TaobaoFilmLotteryDrawRequest) SetAccountType(accountType string) error {
    r.accountType = accountType
    r.Set("account_type", accountType)
    return nil
}

func (r TaobaoFilmLotteryDrawRequest) GetAccountType() string {
    return r.accountType
}

func (r *TaobaoFilmLotteryDrawRequest) SetLotteryMixId(lotteryMixId string) error {
    r.lotteryMixId = lotteryMixId
    r.Set("lottery_mix_id", lotteryMixId)
    return nil
}

func (r TaobaoFilmLotteryDrawRequest) GetLotteryMixId() string {
    return r.lotteryMixId
}

func (r *TaobaoFilmLotteryDrawRequest) SetPlatform(platform int64) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

func (r TaobaoFilmLotteryDrawRequest) GetPlatform() int64 {
    return r.platform
}

func (r *TaobaoFilmLotteryDrawRequest) SetBizData(bizData string) error {
    r.bizData = bizData
    r.Set("biz_data", bizData)
    return nil
}

func (r TaobaoFilmLotteryDrawRequest) GetBizData() string {
    return r.bizData
}

