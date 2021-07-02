package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCaipiaoLotterySendbynickAPIRequest 卖家使用nick给买家送彩票 API请求
// taobao.caipiao.lottery.sendbynick
//
// 卖家使用nick给买家送彩票，可以指定彩种和注数。赠送成功，返回true; 以下几种情况情况， 返回false: 注数超过100注、卖家未签署支付宝代扣协议、卖家或者买家信息不存在等。
type TaobaoCaipiaoLotterySendbynickAPIRequest struct {
	model.Params
	// 彩票接收方nick， 不可为空、""。
	_buyerNick string
	// 彩种ID，此彩种ID为彩票系统中的序号。
	_lotteryTypeId int64
	// 彩票注数，不可为空、0和负数，最大值为100。
	_stakeCount int64
	// 送彩票给接收方的赠言。 不能超过20个字符，1个中文字符、1个英文字母及1个数字等均当作一个字符，如果超过，则会截取。
	_sweetyWords string
}

// NewTaobaoCaipiaoLotterySendbynickRequest 初始化TaobaoCaipiaoLotterySendbynickAPIRequest对象
func NewTaobaoCaipiaoLotterySendbynickRequest() *TaobaoCaipiaoLotterySendbynickAPIRequest {
	return &TaobaoCaipiaoLotterySendbynickAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoLotterySendbynickAPIRequest) GetApiMethodName() string {
	return "taobao.caipiao.lottery.sendbynick"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoLotterySendbynickAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBuyerNick is BuyerNick Setter
// 彩票接收方nick， 不可为空、""。
func (r *TaobaoCaipiaoLotterySendbynickAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoCaipiaoLotterySendbynickAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetLotteryTypeId is LotteryTypeId Setter
// 彩种ID，此彩种ID为彩票系统中的序号。
func (r *TaobaoCaipiaoLotterySendbynickAPIRequest) SetLotteryTypeId(_lotteryTypeId int64) error {
	r._lotteryTypeId = _lotteryTypeId
	r.Set("lottery_type_id", _lotteryTypeId)
	return nil
}

// GetLotteryTypeId LotteryTypeId Getter
func (r TaobaoCaipiaoLotterySendbynickAPIRequest) GetLotteryTypeId() int64 {
	return r._lotteryTypeId
}

// SetStakeCount is StakeCount Setter
// 彩票注数，不可为空、0和负数，最大值为100。
func (r *TaobaoCaipiaoLotterySendbynickAPIRequest) SetStakeCount(_stakeCount int64) error {
	r._stakeCount = _stakeCount
	r.Set("stake_count", _stakeCount)
	return nil
}

// GetStakeCount StakeCount Getter
func (r TaobaoCaipiaoLotterySendbynickAPIRequest) GetStakeCount() int64 {
	return r._stakeCount
}

// SetSweetyWords is SweetyWords Setter
// 送彩票给接收方的赠言。 不能超过20个字符，1个中文字符、1个英文字母及1个数字等均当作一个字符，如果超过，则会截取。
func (r *TaobaoCaipiaoLotterySendbynickAPIRequest) SetSweetyWords(_sweetyWords string) error {
	r._sweetyWords = _sweetyWords
	r.Set("sweety_words", _sweetyWords)
	return nil
}

// GetSweetyWords SweetyWords Getter
func (r TaobaoCaipiaoLotterySendbynickAPIRequest) GetSweetyWords() string {
	return r._sweetyWords
}
