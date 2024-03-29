package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeDrugRefuseorderAPIRequest 阿里健康020拒单 API请求
// taobao.trade.drug.refuseorder
//
// 阿里健康020拒单
type TaobaoTradeDrugRefuseorderAPIRequest struct {
	model.Params
	// 拒单原因
	_refuseReason string
	// 订单ID
	_orderId int64
}

// NewTaobaoTradeDrugRefuseorderRequest 初始化TaobaoTradeDrugRefuseorderAPIRequest对象
func NewTaobaoTradeDrugRefuseorderRequest() *TaobaoTradeDrugRefuseorderAPIRequest {
	return &TaobaoTradeDrugRefuseorderAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTradeDrugRefuseorderAPIRequest) Reset() {
	r._refuseReason = ""
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeDrugRefuseorderAPIRequest) GetApiMethodName() string {
	return "taobao.trade.drug.refuseorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeDrugRefuseorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTradeDrugRefuseorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefuseReason is RefuseReason Setter
// 拒单原因
func (r *TaobaoTradeDrugRefuseorderAPIRequest) SetRefuseReason(_refuseReason string) error {
	r._refuseReason = _refuseReason
	r.Set("refuse_reason", _refuseReason)
	return nil
}

// GetRefuseReason RefuseReason Getter
func (r TaobaoTradeDrugRefuseorderAPIRequest) GetRefuseReason() string {
	return r._refuseReason
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *TaobaoTradeDrugRefuseorderAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoTradeDrugRefuseorderAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolTaobaoTradeDrugRefuseorderAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTradeDrugRefuseorderRequest()
	},
}

// GetTaobaoTradeDrugRefuseorderRequest 从 sync.Pool 获取 TaobaoTradeDrugRefuseorderAPIRequest
func GetTaobaoTradeDrugRefuseorderAPIRequest() *TaobaoTradeDrugRefuseorderAPIRequest {
	return poolTaobaoTradeDrugRefuseorderAPIRequest.Get().(*TaobaoTradeDrugRefuseorderAPIRequest)
}

// ReleaseTaobaoTradeDrugRefuseorderAPIRequest 将 TaobaoTradeDrugRefuseorderAPIRequest 放入 sync.Pool
func ReleaseTaobaoTradeDrugRefuseorderAPIRequest(v *TaobaoTradeDrugRefuseorderAPIRequest) {
	v.Reset()
	poolTaobaoTradeDrugRefuseorderAPIRequest.Put(v)
}
