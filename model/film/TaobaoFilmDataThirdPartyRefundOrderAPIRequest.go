package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmDataThirdPartyRefundOrderAPIRequest 退票接口 API请求
// taobao.film.data.third.party.refund.order
//
// 淘票票第三方退票接口
type TaobaoFilmDataThirdPartyRefundOrderAPIRequest struct {
	model.Params
	// 退票身份ID，用于标识一个购票用户的身份，该参数需要跟锁座接口的ext_order_id参数一致，否则下单会失败。外部渠道需保证该参数的唯一性及准确性，下单出票接口会利用该参数做冥等性判断，如果由于外部渠道自身传入的参数有问题而导致的下单出票接口返回的结果有误，需要外部渠道自己承担损失
	_extUserId string
	// 退款时候需要传入第三方的订单号。外部渠道需保证该参数的唯一性和准确性
	_extOrderId string
	// 目前可以暂时不填参数
	_params string
	// 淘宝账号ID，此ID是一串数字。可自行百度查看如何获取或者咨询淘票票技术人员提供
	_userId int64
	// 淘票票分配的渠道码
	_platform int64
	// 下单时返回的淘宝订单号参数
	_tbOrderId int64
	// 退款金额，以分为单位，为指定的退款订单的金额
	_refundAmount int64
	// 退款服务费，目前都为0
	_refundServiceFee int64
}

// NewTaobaoFilmDataThirdPartyRefundOrderRequest 初始化TaobaoFilmDataThirdPartyRefundOrderAPIRequest对象
func NewTaobaoFilmDataThirdPartyRefundOrderRequest() *TaobaoFilmDataThirdPartyRefundOrderAPIRequest {
	return &TaobaoFilmDataThirdPartyRefundOrderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFilmDataThirdPartyRefundOrderAPIRequest) GetApiMethodName() string {
	return "taobao.film.data.third.party.refund.order"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFilmDataThirdPartyRefundOrderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetExtUserId is ExtUserId Setter
// 退票身份ID，用于标识一个购票用户的身份，该参数需要跟锁座接口的ext_order_id参数一致，否则下单会失败。外部渠道需保证该参数的唯一性及准确性，下单出票接口会利用该参数做冥等性判断，如果由于外部渠道自身传入的参数有问题而导致的下单出票接口返回的结果有误，需要外部渠道自己承担损失
func (r *TaobaoFilmDataThirdPartyRefundOrderAPIRequest) SetExtUserId(_extUserId string) error {
	r._extUserId = _extUserId
	r.Set("ext_user_id", _extUserId)
	return nil
}

// GetExtUserId ExtUserId Getter
func (r TaobaoFilmDataThirdPartyRefundOrderAPIRequest) GetExtUserId() string {
	return r._extUserId
}

// SetExtOrderId is ExtOrderId Setter
// 退款时候需要传入第三方的订单号。外部渠道需保证该参数的唯一性和准确性
func (r *TaobaoFilmDataThirdPartyRefundOrderAPIRequest) SetExtOrderId(_extOrderId string) error {
	r._extOrderId = _extOrderId
	r.Set("ext_order_id", _extOrderId)
	return nil
}

// GetExtOrderId ExtOrderId Getter
func (r TaobaoFilmDataThirdPartyRefundOrderAPIRequest) GetExtOrderId() string {
	return r._extOrderId
}

// SetParams is Params Setter
// 目前可以暂时不填参数
func (r *TaobaoFilmDataThirdPartyRefundOrderAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r TaobaoFilmDataThirdPartyRefundOrderAPIRequest) GetParams() string {
	return r._params
}

// SetUserId is UserId Setter
// 淘宝账号ID，此ID是一串数字。可自行百度查看如何获取或者咨询淘票票技术人员提供
func (r *TaobaoFilmDataThirdPartyRefundOrderAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoFilmDataThirdPartyRefundOrderAPIRequest) GetUserId() int64 {
	return r._userId
}

// SetPlatform is Platform Setter
// 淘票票分配的渠道码
func (r *TaobaoFilmDataThirdPartyRefundOrderAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaoFilmDataThirdPartyRefundOrderAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetTbOrderId is TbOrderId Setter
// 下单时返回的淘宝订单号参数
func (r *TaobaoFilmDataThirdPartyRefundOrderAPIRequest) SetTbOrderId(_tbOrderId int64) error {
	r._tbOrderId = _tbOrderId
	r.Set("tb_order_id", _tbOrderId)
	return nil
}

// GetTbOrderId TbOrderId Getter
func (r TaobaoFilmDataThirdPartyRefundOrderAPIRequest) GetTbOrderId() int64 {
	return r._tbOrderId
}

// SetRefundAmount is RefundAmount Setter
// 退款金额，以分为单位，为指定的退款订单的金额
func (r *TaobaoFilmDataThirdPartyRefundOrderAPIRequest) SetRefundAmount(_refundAmount int64) error {
	r._refundAmount = _refundAmount
	r.Set("refund_amount", _refundAmount)
	return nil
}

// GetRefundAmount RefundAmount Getter
func (r TaobaoFilmDataThirdPartyRefundOrderAPIRequest) GetRefundAmount() int64 {
	return r._refundAmount
}

// SetRefundServiceFee is RefundServiceFee Setter
// 退款服务费，目前都为0
func (r *TaobaoFilmDataThirdPartyRefundOrderAPIRequest) SetRefundServiceFee(_refundServiceFee int64) error {
	r._refundServiceFee = _refundServiceFee
	r.Set("refund_service_fee", _refundServiceFee)
	return nil
}

// GetRefundServiceFee RefundServiceFee Getter
func (r TaobaoFilmDataThirdPartyRefundOrderAPIRequest) GetRefundServiceFee() int64 {
	return r._refundServiceFee
}
