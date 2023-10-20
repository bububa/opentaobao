package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarfpcargetcarnotifyAPIRequest 门店通知用户提车 API请求
// tmall.car.fpcar.getcar.notify
//
// 提供给外部(大搜或其它合作方)的接口-门店通知用户提车
type TmallcarfpcargetcarnotifyAPIRequest struct {
	model.Params
	// 商品宝贝id
	_itemId int64
	// 订单id
	_orderId int64
	// 卖家id
	_sellerId int64
}

// NewTmallcarfpcargetcarnotifyRequest 初始化TmallcarfpcargetcarnotifyAPIRequest对象
func NewTmallcarfpcargetcarnotifyRequest() *TmallcarfpcargetcarnotifyAPIRequest {
	return &TmallcarfpcargetcarnotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarfpcargetcarnotifyAPIRequest) GetApiMethodName() string {
	return "tmall.car.fpcar.getcar.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarfpcargetcarnotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarfpcargetcarnotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品宝贝id
func (r *TmallcarfpcargetcarnotifyAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallcarfpcargetcarnotifyAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetOrderId is OrderId Setter
// 订单id
func (r *TmallcarfpcargetcarnotifyAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallcarfpcargetcarnotifyAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetSellerId is SellerId Setter
// 卖家id
func (r *TmallcarfpcargetcarnotifyAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TmallcarfpcargetcarnotifyAPIRequest) GetSellerId() int64 {
	return r._sellerId
}
