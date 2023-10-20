package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeDrugOrderGetAPIRequest 查看订单详情 API请求
// taobao.trade.drug.order.get
//
// 商家查看订单详情
type TaobaoTradeDrugOrderGetAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewTaobaoTradeDrugOrderGetRequest 初始化TaobaoTradeDrugOrderGetAPIRequest对象
func NewTaobaoTradeDrugOrderGetRequest() *TaobaoTradeDrugOrderGetAPIRequest {
	return &TaobaoTradeDrugOrderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTradeDrugOrderGetAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeDrugOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.drug.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeDrugOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTradeDrugOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *TaobaoTradeDrugOrderGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoTradeDrugOrderGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolTaobaoTradeDrugOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTradeDrugOrderGetRequest()
	},
}

// GetTaobaoTradeDrugOrderGetRequest 从 sync.Pool 获取 TaobaoTradeDrugOrderGetAPIRequest
func GetTaobaoTradeDrugOrderGetAPIRequest() *TaobaoTradeDrugOrderGetAPIRequest {
	return poolTaobaoTradeDrugOrderGetAPIRequest.Get().(*TaobaoTradeDrugOrderGetAPIRequest)
}

// ReleaseTaobaoTradeDrugOrderGetAPIRequest 将 TaobaoTradeDrugOrderGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTradeDrugOrderGetAPIRequest(v *TaobaoTradeDrugOrderGetAPIRequest) {
	v.Reset()
	poolTaobaoTradeDrugOrderGetAPIRequest.Put(v)
}
