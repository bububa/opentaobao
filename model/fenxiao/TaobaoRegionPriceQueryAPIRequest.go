package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRegionPriceQueryAPIRequest 区域价格查询 API请求
// taobao.region.price.query
//
// 区域价格查询
type TaobaoRegionPriceQueryAPIRequest struct {
	model.Params
	// 不传则返回所有设置的区域价格
	_regionalPriceDtos []RegionalPriceDto
	// 商品id
	_itemId int64
	// 无sku可传0
	_skuId int64
}

// NewTaobaoRegionPriceQueryRequest 初始化TaobaoRegionPriceQueryAPIRequest对象
func NewTaobaoRegionPriceQueryRequest() *TaobaoRegionPriceQueryAPIRequest {
	return &TaobaoRegionPriceQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRegionPriceQueryAPIRequest) Reset() {
	r._regionalPriceDtos = r._regionalPriceDtos[:0]
	r._itemId = 0
	r._skuId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRegionPriceQueryAPIRequest) GetApiMethodName() string {
	return "taobao.region.price.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRegionPriceQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRegionPriceQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRegionalPriceDtos is RegionalPriceDtos Setter
// 不传则返回所有设置的区域价格
func (r *TaobaoRegionPriceQueryAPIRequest) SetRegionalPriceDtos(_regionalPriceDtos []RegionalPriceDto) error {
	r._regionalPriceDtos = _regionalPriceDtos
	r.Set("regional_price_dtos", _regionalPriceDtos)
	return nil
}

// GetRegionalPriceDtos RegionalPriceDtos Getter
func (r TaobaoRegionPriceQueryAPIRequest) GetRegionalPriceDtos() []RegionalPriceDto {
	return r._regionalPriceDtos
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoRegionPriceQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoRegionPriceQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// 无sku可传0
func (r *TaobaoRegionPriceQueryAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoRegionPriceQueryAPIRequest) GetSkuId() int64 {
	return r._skuId
}

var poolTaobaoRegionPriceQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRegionPriceQueryRequest()
	},
}

// GetTaobaoRegionPriceQueryRequest 从 sync.Pool 获取 TaobaoRegionPriceQueryAPIRequest
func GetTaobaoRegionPriceQueryAPIRequest() *TaobaoRegionPriceQueryAPIRequest {
	return poolTaobaoRegionPriceQueryAPIRequest.Get().(*TaobaoRegionPriceQueryAPIRequest)
}

// ReleaseTaobaoRegionPriceQueryAPIRequest 将 TaobaoRegionPriceQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoRegionPriceQueryAPIRequest(v *TaobaoRegionPriceQueryAPIRequest) {
	v.Reset()
	poolTaobaoRegionPriceQueryAPIRequest.Put(v)
}
