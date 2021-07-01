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
type TaobaoFilmLotteryDrawAPIRequest struct {
    model.Params
    // 账号ID
    _accountNo   string
    // 账号类型（TAOBAO\ALIPAY\PHONE\TAOBAO_NAME\OPEN_ID）
    _accountType   string
    // 活动ID
    _lotteryMixId   string
    // 平台类型
    _platform   int64
    // 扩展参数
    _bizData   string
}

// 初始化TaobaoFilmLotteryDrawAPIRequest对象
func NewTaobaoFilmLotteryDrawRequest() *TaobaoFilmLotteryDrawAPIRequest{
    return &TaobaoFilmLotteryDrawAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFilmLotteryDrawAPIRequest) GetApiMethodName() string {
    return "taobao.film.lottery.draw"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFilmLotteryDrawAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountNo Setter
// 账号ID
func (r *TaobaoFilmLotteryDrawAPIRequest) SetAccountNo(_accountNo string) error {
    r._accountNo = _accountNo
    r.Set("account_no", _accountNo)
    return nil
}

// AccountNo Getter
func (r TaobaoFilmLotteryDrawAPIRequest) GetAccountNo() string {
    return r._accountNo
}
// AccountType Setter
// 账号类型（TAOBAO\ALIPAY\PHONE\TAOBAO_NAME\OPEN_ID）
func (r *TaobaoFilmLotteryDrawAPIRequest) SetAccountType(_accountType string) error {
    r._accountType = _accountType
    r.Set("account_type", _accountType)
    return nil
}

// AccountType Getter
func (r TaobaoFilmLotteryDrawAPIRequest) GetAccountType() string {
    return r._accountType
}
// LotteryMixId Setter
// 活动ID
func (r *TaobaoFilmLotteryDrawAPIRequest) SetLotteryMixId(_lotteryMixId string) error {
    r._lotteryMixId = _lotteryMixId
    r.Set("lottery_mix_id", _lotteryMixId)
    return nil
}

// LotteryMixId Getter
func (r TaobaoFilmLotteryDrawAPIRequest) GetLotteryMixId() string {
    return r._lotteryMixId
}
// Platform Setter
// 平台类型
func (r *TaobaoFilmLotteryDrawAPIRequest) SetPlatform(_platform int64) error {
    r._platform = _platform
    r.Set("platform", _platform)
    return nil
}

// Platform Getter
func (r TaobaoFilmLotteryDrawAPIRequest) GetPlatform() int64 {
    return r._platform
}
// BizData Setter
// 扩展参数
func (r *TaobaoFilmLotteryDrawAPIRequest) SetBizData(_bizData string) error {
    r._bizData = _bizData
    r.Set("biz_data", _bizData)
    return nil
}

// BizData Getter
func (r TaobaoFilmLotteryDrawAPIRequest) GetBizData() string {
    return r._bizData
}
