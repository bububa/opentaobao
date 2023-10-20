package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarFpcarRestpayReceiveAPIRequest 门店线下已收尾款 API请求
// tmall.car.fpcar.restpay.receive
//
// 提供给外部(大搜或其它合作方)的接口-门店线下已收尾款(不执行分佣)
type TmallCarFpcarRestpayReceiveAPIRequest struct {
	model.Params
	// 卖家id
	_sellerId int64
	// 订单id
	_orderId int64
	// 商品宝贝id
	_itemId int64
}

// NewTmallCarFpcarRestpayReceiveRequest 初始化TmallCarFpcarRestpayReceiveAPIRequest对象
func NewTmallCarFpcarRestpayReceiveRequest() *TmallCarFpcarRestpayReceiveAPIRequest {
	return &TmallCarFpcarRestpayReceiveAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarFpcarRestpayReceiveAPIRequest) Reset() {
	r._sellerId = 0
	r._orderId = 0
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarFpcarRestpayReceiveAPIRequest) GetApiMethodName() string {
	return "tmall.car.fpcar.restpay.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarFpcarRestpayReceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarFpcarRestpayReceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerId is SellerId Setter
// 卖家id
func (r *TmallCarFpcarRestpayReceiveAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TmallCarFpcarRestpayReceiveAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// SetOrderId is OrderId Setter
// 订单id
func (r *TmallCarFpcarRestpayReceiveAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallCarFpcarRestpayReceiveAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetItemId is ItemId Setter
// 商品宝贝id
func (r *TmallCarFpcarRestpayReceiveAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallCarFpcarRestpayReceiveAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTmallCarFpcarRestpayReceiveAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarFpcarRestpayReceiveRequest()
	},
}

// GetTmallCarFpcarRestpayReceiveRequest 从 sync.Pool 获取 TmallCarFpcarRestpayReceiveAPIRequest
func GetTmallCarFpcarRestpayReceiveAPIRequest() *TmallCarFpcarRestpayReceiveAPIRequest {
	return poolTmallCarFpcarRestpayReceiveAPIRequest.Get().(*TmallCarFpcarRestpayReceiveAPIRequest)
}

// ReleaseTmallCarFpcarRestpayReceiveAPIRequest 将 TmallCarFpcarRestpayReceiveAPIRequest 放入 sync.Pool
func ReleaseTmallCarFpcarRestpayReceiveAPIRequest(v *TmallCarFpcarRestpayReceiveAPIRequest) {
	v.Reset()
	poolTmallCarFpcarRestpayReceiveAPIRequest.Put(v)
}
