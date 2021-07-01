package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripSellerRefundmoneyConfirmAPIRequest
【机票代理商订单】确认退款 API请求
taobao.alitrip.seller.refundmoney.confirm

代理人确认退票申请单的退款 */
type TaobaoAlitripSellerRefundmoneyConfirmAPIRequest struct {
	model.Params
	// 申请单id
	_applyId int64
}

// NewTaobaoAlitripSellerRefundmoneyConfirmRequest 初始化TaobaoAlitripSellerRefundmoneyConfirmAPIRequest对象
func NewTaobaoAlitripSellerRefundmoneyConfirmRequest() *TaobaoAlitripSellerRefundmoneyConfirmAPIRequest {
	return &TaobaoAlitripSellerRefundmoneyConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundmoneyConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refundmoney.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundmoneyConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ApplyId Setter
// 申请单id
func (r *TaobaoAlitripSellerRefundmoneyConfirmAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// Get ApplyId Getter
func (r TaobaoAlitripSellerRefundmoneyConfirmAPIRequest) GetApplyId() int64 {
	return r._applyId
}
