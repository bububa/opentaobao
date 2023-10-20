package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRegionPriceCancleAPIRequest 取消区域价格 API请求
// taobao.region.price.cancle
//
// 取消区域价格
type TaobaoRegionPriceCancleAPIRequest struct {
	model.Params
	// 商品
	_itemId int64
	// 无sku传0
	_skuId int64
}

// NewTaobaoRegionPriceCancleRequest 初始化TaobaoRegionPriceCancleAPIRequest对象
func NewTaobaoRegionPriceCancleRequest() *TaobaoRegionPriceCancleAPIRequest {
	return &TaobaoRegionPriceCancleAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRegionPriceCancleAPIRequest) Reset() {
	r._itemId = 0
	r._skuId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRegionPriceCancleAPIRequest) GetApiMethodName() string {
	return "taobao.region.price.cancle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRegionPriceCancleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRegionPriceCancleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品
func (r *TaobaoRegionPriceCancleAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoRegionPriceCancleAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// 无sku传0
func (r *TaobaoRegionPriceCancleAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoRegionPriceCancleAPIRequest) GetSkuId() int64 {
	return r._skuId
}

var poolTaobaoRegionPriceCancleAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRegionPriceCancleRequest()
	},
}

// GetTaobaoRegionPriceCancleRequest 从 sync.Pool 获取 TaobaoRegionPriceCancleAPIRequest
func GetTaobaoRegionPriceCancleAPIRequest() *TaobaoRegionPriceCancleAPIRequest {
	return poolTaobaoRegionPriceCancleAPIRequest.Get().(*TaobaoRegionPriceCancleAPIRequest)
}

// ReleaseTaobaoRegionPriceCancleAPIRequest 将 TaobaoRegionPriceCancleAPIRequest 放入 sync.Pool
func ReleaseTaobaoRegionPriceCancleAPIRequest(v *TaobaoRegionPriceCancleAPIRequest) {
	v.Reset()
	poolTaobaoRegionPriceCancleAPIRequest.Put(v)
}
