package drug

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest 根据订单id获取单条订单详情 API请求
// alibaba.alihealth.nr.trade.order.getorderdetail
//
// 阿里健康O2O，获取订单详情，修复组合商品价格精度问题
type AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest struct {
	model.Params
	// 淘宝订单ID
	_orderId int64
}

// NewAlibabaAlihealthNrTradeOrderGetorderdetailRequest 初始化AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest对象
func NewAlibabaAlihealthNrTradeOrderGetorderdetailRequest() *AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest {
	return &AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nr.trade.order.getorderdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 淘宝订单ID
func (r *AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolAlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthNrTradeOrderGetorderdetailRequest()
	},
}

// GetAlibabaAlihealthNrTradeOrderGetorderdetailRequest 从 sync.Pool 获取 AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest
func GetAlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest() *AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest {
	return poolAlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest.Get().(*AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest)
}

// ReleaseAlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest 将 AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest(v *AlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthNrTradeOrderGetorderdetailAPIRequest.Put(v)
}
