package jipiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerRefundRefusereturnAPIRequest 【机票代理商】拒绝退票 API请求
// taobao.alitrip.seller.refund.refusereturn
//
// 拒绝退票
type TaobaoAlitripSellerRefundRefusereturnAPIRequest struct {
	model.Params
	// 拒绝理由
	_reason string
	// 申请单ID
	_applyId int64
}

// NewTaobaoAlitripSellerRefundRefusereturnRequest 初始化TaobaoAlitripSellerRefundRefusereturnAPIRequest对象
func NewTaobaoAlitripSellerRefundRefusereturnRequest() *TaobaoAlitripSellerRefundRefusereturnAPIRequest {
	return &TaobaoAlitripSellerRefundRefusereturnAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripSellerRefundRefusereturnAPIRequest) Reset() {
	r._reason = ""
	r._applyId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundRefusereturnAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refund.refusereturn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundRefusereturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripSellerRefundRefusereturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReason is Reason Setter
// 拒绝理由
func (r *TaobaoAlitripSellerRefundRefusereturnAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoAlitripSellerRefundRefusereturnAPIRequest) GetReason() string {
	return r._reason
}

// SetApplyId is ApplyId Setter
// 申请单ID
func (r *TaobaoAlitripSellerRefundRefusereturnAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r TaobaoAlitripSellerRefundRefusereturnAPIRequest) GetApplyId() int64 {
	return r._applyId
}

var poolTaobaoAlitripSellerRefundRefusereturnAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripSellerRefundRefusereturnRequest()
	},
}

// GetTaobaoAlitripSellerRefundRefusereturnRequest 从 sync.Pool 获取 TaobaoAlitripSellerRefundRefusereturnAPIRequest
func GetTaobaoAlitripSellerRefundRefusereturnAPIRequest() *TaobaoAlitripSellerRefundRefusereturnAPIRequest {
	return poolTaobaoAlitripSellerRefundRefusereturnAPIRequest.Get().(*TaobaoAlitripSellerRefundRefusereturnAPIRequest)
}

// ReleaseTaobaoAlitripSellerRefundRefusereturnAPIRequest 将 TaobaoAlitripSellerRefundRefusereturnAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripSellerRefundRefusereturnAPIRequest(v *TaobaoAlitripSellerRefundRefusereturnAPIRequest) {
	v.Reset()
	poolTaobaoAlitripSellerRefundRefusereturnAPIRequest.Put(v)
}
