package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemSingleQueryAPIRequest 【API3.0】度假单个商品查询接口 API请求
// taobao.alitrip.travel.item.single.query
//
// 旅行度假新商品查询接口（单个商品查询） 第三版
type TaobaoAlitripTravelItemSingleQueryAPIRequest struct {
	model.Params
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
}

// NewTaobaoAlitripTravelItemSingleQueryRequest 初始化TaobaoAlitripTravelItemSingleQueryAPIRequest对象
func NewTaobaoAlitripTravelItemSingleQueryRequest() *TaobaoAlitripTravelItemSingleQueryAPIRequest {
	return &TaobaoAlitripTravelItemSingleQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelItemSingleQueryAPIRequest) Reset() {
	r._outProductId = ""
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemSingleQueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.item.single.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemSingleQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelItemSingleQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutProductId is OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSingleQueryAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r TaobaoAlitripTravelItemSingleQueryAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetItemId is ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSingleQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoAlitripTravelItemSingleQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoAlitripTravelItemSingleQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelItemSingleQueryRequest()
	},
}

// GetTaobaoAlitripTravelItemSingleQueryRequest 从 sync.Pool 获取 TaobaoAlitripTravelItemSingleQueryAPIRequest
func GetTaobaoAlitripTravelItemSingleQueryAPIRequest() *TaobaoAlitripTravelItemSingleQueryAPIRequest {
	return poolTaobaoAlitripTravelItemSingleQueryAPIRequest.Get().(*TaobaoAlitripTravelItemSingleQueryAPIRequest)
}

// ReleaseTaobaoAlitripTravelItemSingleQueryAPIRequest 将 TaobaoAlitripTravelItemSingleQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelItemSingleQueryAPIRequest(v *TaobaoAlitripTravelItemSingleQueryAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelItemSingleQueryAPIRequest.Put(v)
}
