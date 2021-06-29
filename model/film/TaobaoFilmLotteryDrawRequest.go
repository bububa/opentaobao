package film

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘票票抽奖发放权益API API请求
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

// 初始化TaobaoFilmLotteryDrawRequest对象
func NewTaobaoFilmLotteryDrawRequest() *TaobaoFilmLotteryDrawRequest{
    return &TaobaoFilmLotteryDrawRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFilmLotteryDrawRequest) GetApiMethodName() string {
    return "taobao.film.lottery.draw"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFilmLotteryDrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountNo Setter
// 账号ID
func (r *TaobaoFilmLotteryDrawRequest) SetAccountNo(accountNo string) error {
    r.accountNo = accountNo
    r.Set("account_no", accountNo)
    return nil
}

// AccountNo Getter
func (r TaobaoFilmLotteryDrawRequest) GetAccountNo() string {
    return r.accountNo
}
// AccountType Setter
// 账号类型（TAOBAO\ALIPAY\PHONE\TAOBAO_NAME\OPEN_ID）
func (r *TaobaoFilmLotteryDrawRequest) SetAccountType(accountType string) error {
    r.accountType = accountType
    r.Set("account_type", accountType)
    return nil
}

// AccountType Getter
func (r TaobaoFilmLotteryDrawRequest) GetAccountType() string {
    return r.accountType
}
// LotteryMixId Setter
// 活动ID
func (r *TaobaoFilmLotteryDrawRequest) SetLotteryMixId(lotteryMixId string) error {
    r.lotteryMixId = lotteryMixId
    r.Set("lottery_mix_id", lotteryMixId)
    return nil
}

// LotteryMixId Getter
func (r TaobaoFilmLotteryDrawRequest) GetLotteryMixId() string {
    return r.lotteryMixId
}
// Platform Setter
// 平台类型
func (r *TaobaoFilmLotteryDrawRequest) SetPlatform(platform int64) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

// Platform Getter
func (r TaobaoFilmLotteryDrawRequest) GetPlatform() int64 {
    return r.platform
}
// BizData Setter
// 扩展参数
func (r *TaobaoFilmLotteryDrawRequest) SetBizData(bizData string) error {
    r.bizData = bizData
    r.Set("biz_data", bizData)
    return nil
}

// BizData Getter
func (r TaobaoFilmLotteryDrawRequest) GetBizData() string {
    return r.bizData
}
