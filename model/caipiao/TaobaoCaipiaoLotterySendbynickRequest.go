package caipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家使用nick给买家送彩票 API请求
taobao.caipiao.lottery.sendbynick

卖家使用nick给买家送彩票，可以指定彩种和注数。赠送成功，返回true; 以下几种情况情况， 返回false: 注数超过100注、卖家未签署支付宝代扣协议、卖家或者买家信息不存在等。
*/
type TaobaoCaipiaoLotterySendbynickRequest struct {
    model.Params
    // 彩票接收方nick， 不可为空、""。
    buyerNick   string
    // 彩种ID，此彩种ID为彩票系统中的序号。
    lotteryTypeId   int64
    // 彩票注数，不可为空、0和负数，最大值为100。
    stakeCount   int64
    // 送彩票给接收方的赠言。 不能超过20个字符，1个中文字符、1个英文字母及1个数字等均当作一个字符，如果超过，则会截取。
    sweetyWords   string
}

// 初始化TaobaoCaipiaoLotterySendbynickRequest对象
func NewTaobaoCaipiaoLotterySendbynickRequest() *TaobaoCaipiaoLotterySendbynickRequest{
    return &TaobaoCaipiaoLotterySendbynickRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoLotterySendbynickRequest) GetApiMethodName() string {
    return "taobao.caipiao.lottery.sendbynick"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoLotterySendbynickRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BuyerNick Setter
// 彩票接收方nick， 不可为空、""。
func (r *TaobaoCaipiaoLotterySendbynickRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoCaipiaoLotterySendbynickRequest) GetBuyerNick() string {
    return r.buyerNick
}
// LotteryTypeId Setter
// 彩种ID，此彩种ID为彩票系统中的序号。
func (r *TaobaoCaipiaoLotterySendbynickRequest) SetLotteryTypeId(lotteryTypeId int64) error {
    r.lotteryTypeId = lotteryTypeId
    r.Set("lottery_type_id", lotteryTypeId)
    return nil
}

// LotteryTypeId Getter
func (r TaobaoCaipiaoLotterySendbynickRequest) GetLotteryTypeId() int64 {
    return r.lotteryTypeId
}
// StakeCount Setter
// 彩票注数，不可为空、0和负数，最大值为100。
func (r *TaobaoCaipiaoLotterySendbynickRequest) SetStakeCount(stakeCount int64) error {
    r.stakeCount = stakeCount
    r.Set("stake_count", stakeCount)
    return nil
}

// StakeCount Getter
func (r TaobaoCaipiaoLotterySendbynickRequest) GetStakeCount() int64 {
    return r.stakeCount
}
// SweetyWords Setter
// 送彩票给接收方的赠言。 不能超过20个字符，1个中文字符、1个英文字母及1个数字等均当作一个字符，如果超过，则会截取。
func (r *TaobaoCaipiaoLotterySendbynickRequest) SetSweetyWords(sweetyWords string) error {
    r.sweetyWords = sweetyWords
    r.Set("sweety_words", sweetyWords)
    return nil
}

// SweetyWords Getter
func (r TaobaoCaipiaoLotterySendbynickRequest) GetSweetyWords() string {
    return r.sweetyWords
}
