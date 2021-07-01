package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCaipiaoLotterySendbynickAPIRequest
卖家使用nick给买家送彩票 API请求
taobao.caipiao.lottery.sendbynick

卖家使用nick给买家送彩票，可以指定彩种和注数。赠送成功，返回true; 以下几种情况情况， 返回false: 注数超过100注、卖家未签署支付宝代扣协议、卖家或者买家信息不存在等。 */
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

// New
