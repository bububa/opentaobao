package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemSkuGetAPIRequest 获取全渠道门店商品sku API请求
// taobao.omniitem.sku.get
//
// 通过skuId或者skuOutId查询全渠道门店商品sku信息
type TaobaoOmniitemSkuGetAPIRequest struct {
	model.Params
	// sku商家编码
	_skuOuterId string
	// 商品id
	_itemId int64
	// skuId
	_skuId int64
}

// NewTaobaoOmniitemSkuGetRequest 初始化TaobaoOmniitemSkuGetAPIRequest对象
func NewTaobaoOmniitemSkuGetRequest() *TaobaoOmniitemSkuGetAPIRequest {
	return &TaobaoOmniitemSkuGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniitemSkuGetAPIRequest) Reset() {
	r._skuOuterId = ""
	r._itemId = 0
	r._skuId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemSkuGetAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.sku.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemSkuGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniitemSkuGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuOuterId is SkuOuterId Setter
// sku商家编码
func (r *TaobaoOmniitemSkuGetAPIRequest) SetSkuOuterId(_skuOuterId string) error {
	r._skuOuterId = _skuOuterId
	r.Set("sku_outer_id", _skuOuterId)
	return nil
}

// GetSkuOuterId SkuOuterId Getter
func (r TaobaoOmniitemSkuGetAPIRequest) GetSkuOuterId() string {
	return r._skuOuterId
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoOmniitemSkuGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoOmniitemSkuGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// skuId
func (r *TaobaoOmniitemSkuGetAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoOmniitemSkuGetAPIRequest) GetSkuId() int64 {
	return r._skuId
}

var poolTaobaoOmniitemSkuGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniitemSkuGetRequest()
	},
}

// GetTaobaoOmniitemSkuGetRequest 从 sync.Pool 获取 TaobaoOmniitemSkuGetAPIRequest
func GetTaobaoOmniitemSkuGetAPIRequest() *TaobaoOmniitemSkuGetAPIRequest {
	return poolTaobaoOmniitemSkuGetAPIRequest.Get().(*TaobaoOmniitemSkuGetAPIRequest)
}

// ReleaseTaobaoOmniitemSkuGetAPIRequest 将 TaobaoOmniitemSkuGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniitemSkuGetAPIRequest(v *TaobaoOmniitemSkuGetAPIRequest) {
	v.Reset()
	poolTaobaoOmniitemSkuGetAPIRequest.Put(v)
}
