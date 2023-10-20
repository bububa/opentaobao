package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpPromotionGetAPIRequest 商品优惠详情查询 API请求
// taobao.ump.promotion.get
//
// 商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
type TaobaoUmpPromotionGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// NewTaobaoUmpPromotionGetRequest 初始化TaobaoUmpPromotionGetAPIRequest对象
func NewTaobaoUmpPromotionGetRequest() *TaobaoUmpPromotionGetAPIRequest {
	return &TaobaoUmpPromotionGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpPromotionGetAPIRequest) Reset() {
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpPromotionGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.promotion.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpPromotionGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpPromotionGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoUmpPromotionGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoUmpPromotionGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoUmpPromotionGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpPromotionGetRequest()
	},
}

// GetTaobaoUmpPromotionGetRequest 从 sync.Pool 获取 TaobaoUmpPromotionGetAPIRequest
func GetTaobaoUmpPromotionGetAPIRequest() *TaobaoUmpPromotionGetAPIRequest {
	return poolTaobaoUmpPromotionGetAPIRequest.Get().(*TaobaoUmpPromotionGetAPIRequest)
}

// ReleaseTaobaoUmpPromotionGetAPIRequest 将 TaobaoUmpPromotionGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpPromotionGetAPIRequest(v *TaobaoUmpPromotionGetAPIRequest) {
	v.Reset()
	poolTaobaoUmpPromotionGetAPIRequest.Put(v)
}
