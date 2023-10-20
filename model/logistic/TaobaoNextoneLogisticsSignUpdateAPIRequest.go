package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoNextoneLogisticsSignUpdateAPIRequest AG物流签收状态写接口 API请求
// taobao.nextone.logistics.sign.update
//
// 商家上传退货的签收状态给AG
type TaobaoNextoneLogisticsSignUpdateAPIRequest struct {
	model.Params
	// 退款编号
	_refundId int64
	// 货物签收状态
	_signStatus int64
}

// NewTaobaoNextoneLogisticsSignUpdateRequest 初始化TaobaoNextoneLogisticsSignUpdateAPIRequest对象
func NewTaobaoNextoneLogisticsSignUpdateRequest() *TaobaoNextoneLogisticsSignUpdateAPIRequest {
	return &TaobaoNextoneLogisticsSignUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoNextoneLogisticsSignUpdateAPIRequest) Reset() {
	r._refundId = 0
	r._signStatus = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoNextoneLogisticsSignUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.nextone.logistics.sign.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoNextoneLogisticsSignUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoNextoneLogisticsSignUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaoNextoneLogisticsSignUpdateAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoNextoneLogisticsSignUpdateAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetSignStatus is SignStatus Setter
// 货物签收状态
func (r *TaobaoNextoneLogisticsSignUpdateAPIRequest) SetSignStatus(_signStatus int64) error {
	r._signStatus = _signStatus
	r.Set("sign_status", _signStatus)
	return nil
}

// GetSignStatus SignStatus Getter
func (r TaobaoNextoneLogisticsSignUpdateAPIRequest) GetSignStatus() int64 {
	return r._signStatus
}

var poolTaobaoNextoneLogisticsSignUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoNextoneLogisticsSignUpdateRequest()
	},
}

// GetTaobaoNextoneLogisticsSignUpdateRequest 从 sync.Pool 获取 TaobaoNextoneLogisticsSignUpdateAPIRequest
func GetTaobaoNextoneLogisticsSignUpdateAPIRequest() *TaobaoNextoneLogisticsSignUpdateAPIRequest {
	return poolTaobaoNextoneLogisticsSignUpdateAPIRequest.Get().(*TaobaoNextoneLogisticsSignUpdateAPIRequest)
}

// ReleaseTaobaoNextoneLogisticsSignUpdateAPIRequest 将 TaobaoNextoneLogisticsSignUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoNextoneLogisticsSignUpdateAPIRequest(v *TaobaoNextoneLogisticsSignUpdateAPIRequest) {
	v.Reset()
	poolTaobaoNextoneLogisticsSignUpdateAPIRequest.Put(v)
}
