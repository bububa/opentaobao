package jipiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerRefundConfirmreturnAPIRequest 【机票代理商】确认退票 API请求
// taobao.alitrip.seller.refund.confirmreturn
//
// 确认退票
type TaobaoAlitripSellerRefundConfirmreturnAPIRequest struct {
	model.Params
	// 退票申请单
	_applyId int64
}

// NewTaobaoAlitripSellerRefundConfirmreturnRequest 初始化TaobaoAlitripSellerRefundConfirmreturnAPIRequest对象
func NewTaobaoAlitripSellerRefundConfirmreturnRequest() *TaobaoAlitripSellerRefundConfirmreturnAPIRequest {
	return &TaobaoAlitripSellerRefundConfirmreturnAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripSellerRefundConfirmreturnAPIRequest) Reset() {
	r._applyId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundConfirmreturnAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refund.confirmreturn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundConfirmreturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripSellerRefundConfirmreturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyId is ApplyId Setter
// 退票申请单
func (r *TaobaoAlitripSellerRefundConfirmreturnAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoAlitripSellerRefundConfirmreturnAPIRequest) GetApplyId() int64 {
	return r._applyId
}

var poolTaobaoAlitripSellerRefundConfirmreturnAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripSellerRefundConfirmreturnRequest()
	},
}

// GetTaobaoAlitripSellerRefundConfirmreturnRequest 从 sync.Pool 获取 TaobaoAlitripSellerRefundConfirmreturnAPIRequest
func GetTaobaoAlitripSellerRefundConfirmreturnAPIRequest() *TaobaoAlitripSellerRefundConfirmreturnAPIRequest {
	return poolTaobaoAlitripSellerRefundConfirmreturnAPIRequest.Get().(*TaobaoAlitripSellerRefundConfirmreturnAPIRequest)
}

// ReleaseTaobaoAlitripSellerRefundConfirmreturnAPIRequest 将 TaobaoAlitripSellerRefundConfirmreturnAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripSellerRefundConfirmreturnAPIRequest(v *TaobaoAlitripSellerRefundConfirmreturnAPIRequest) {
	v.Reset()
	poolTaobaoAlitripSellerRefundConfirmreturnAPIRequest.Put(v)
}
