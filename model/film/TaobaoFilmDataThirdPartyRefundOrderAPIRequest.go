package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFilmDataThirdPartyRefundOrderAPIRequest
退票接口 API请求
taobao.film.data.third.party.refund.order

淘票票第三方退票接口 */
type TaobaoFilmDataThirdPartyRefundOrderAPIRequest struct {
	model.Params
	// 淘宝账号ID，此ID是一串数字。可自行百度查看如何获取或者咨询淘票票技术人员提供
	_userId int64
	// 淘票票分配的渠道码
	_platform int64
	// 退票身份ID，用于标识一个购票用户的身份，该参数需要跟锁座接口的ext_order_id参数一致，否则下单会失败。外部渠道需保证该参数的唯一性及准确性，下单出票接口会利用该参数做冥等性判断，如果由于外部渠道自身传入的参数有问题而导致的下单出票接口返回的结果有误，需要外部渠道自己承担损失
	_extUserId string
	// 退款时候需要传入第三方的订单号。外部渠道需保证该参数的唯一性和准确性
	_extOrderId string
	// 下单时返回的淘宝订单号参数
	_tbOrderId int64
	// 退款金额，以分为单位，为指定的退款订单的金额
	_refundAmount int64
	// 退款服务费，目前都为0
	_refundServiceFee int64
	// 目前可以暂时不填参数
	_params string
}

// New
