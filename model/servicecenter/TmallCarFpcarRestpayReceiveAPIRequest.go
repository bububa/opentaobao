package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarfpcarrestpayreceiveAPIRequest 门店线下已收尾款 API请求
// tmall.car.fpcar.restpay.receive
//
// 提供给外部(大搜或其它合作方)的接口-门店线下已收尾款(不执行分佣)
type TmallcarfpcarrestpayreceiveAPIRequest struct {
	model.Params
	// 卖家id
	_sellerId int64
	// 订单id
	_orderId int64
	// 商品宝贝id
	_itemId int64
}

// NewTmallcarfpcarrestpayreceiveRequest 初始化TmallcarfpcarrestpayreceiveAPIRequest对象
func NewTmallcarfpcarrestpayreceiveRequest() *TmallcarfpcarrestpayreceiveAPIRequest {
	return &TmallcarfpcarrestpayreceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarfpcarrestpayreceiveAPIRequest) GetApiMethodName() string {
	return "tmall.car.fpcar.restpay.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarfpcarrestpayreceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarfpcarrestpayreceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerId is SellerId Setter
// 卖家id
func (r *TmallcarfpcarrestpayreceiveAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TmallcarfpcarrestpayreceiveAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// SetOrderId is OrderId Setter
// 订单id
func (r *TmallcarfpcarrestpayreceiveAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallcarfpcarrestpayreceiveAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetItemId is ItemId Setter
// 商品宝贝id
func (r *TmallcarfpcarrestpayreceiveAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallcarfpcarrestpayreceiveAPIRequest) GetItemId() int64 {
	return r._itemId
}
