package jipiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerRefundmoneyConfirmAPIRequest 【机票代理商订单】确认退款 API请求
// taobao.alitrip.seller.refundmoney.confirm
//
// 代理人确认退票申请单的退款
type TaobaoAlitripSellerRefundmoneyConfirmAPIRequest struct {
	model.Params
	// 申请单id
	_applyId int64
}

// NewTaobaoAlitripSellerRefundmoneyConfirmRequest 初始化TaobaoAlitripSellerRefundmoneyConfirmAPIRequest对象
func NewTaobaoAlitripSellerRefundmoneyConfirmRequest() *TaobaoAlitripSellerRefundmoneyConfirmAPIRequest {
	return &TaobaoAlitripSellerRefundmoneyConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripSellerRefundmoneyConfirmAPIRequest) Reset() {
	r._applyId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundmoneyConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refundmoney.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundmoneyConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripSellerRefundmoneyConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 申请单id
func (r *TaobaoAlitripSellerRefundmoneyConfirmAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoAlitripSellerRefundmoneyConfirmAPIRequest) GetApplyId() int64 {
	return r._applyId
}

var poolTaobaoAlitripSellerRefundmoneyConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripSellerRefundmoneyConfirmRequest()
	},
}

// GetTaobaoAlitripSellerRefundmoneyConfirmRequest 从 sync.Pool 获取 TaobaoAlitripSellerRefundmoneyConfirmAPIRequest
func GetTaobaoAlitripSellerRefundmoneyConfirmAPIRequest() *TaobaoAlitripSellerRefundmoneyConfirmAPIRequest {
	return poolTaobaoAlitripSellerRefundmoneyConfirmAPIRequest.Get().(*TaobaoAlitripSellerRefundmoneyConfirmAPIRequest)
}

// ReleaseTaobaoAlitripSellerRefundmoneyConfirmAPIRequest 将 TaobaoAlitripSellerRefundmoneyConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripSellerRefundmoneyConfirmAPIRequest(v *TaobaoAlitripSellerRefundmoneyConfirmAPIRequest) {
	v.Reset()
	poolTaobaoAlitripSellerRefundmoneyConfirmAPIRequest.Put(v)
}
