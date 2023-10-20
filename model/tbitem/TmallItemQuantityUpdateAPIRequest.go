package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemQuantityUpdateAPIRequest 天猫商品/SKU库存更新接口 API请求
// tmall.item.quantity.update
//
// 天猫商品/SKU库存更新接口；支持商品库存更新；支持同一商品下的SKU批量更新。
type TmallItemQuantityUpdateAPIRequest struct {
	model.Params
	// 更新SKU库存时候的SKU库存对象；如果没有SKU或者不更新SKU库存，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！
	_skuQuantities []UpdateSkuQuantity
	// 商品id
	_itemId int64
	// 商品库存数；增量编辑方式支持正数、负数。（无SKU商品使用这个字段）
	_itemQuantity int64
	// 商品库存更新时候的可选参数
	_options *UpdateItemQuantityOption
}

// NewTmallItemQuantityUpdateRequest 初始化TmallItemQuantityUpdateAPIRequest对象
func NewTmallItemQuantityUpdateRequest() *TmallItemQuantityUpdateAPIRequest {
	return &TmallItemQuantityUpdateAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemQuantityUpdateAPIRequest) Reset() {
	r._skuQuantities = r._skuQuantities[:0]
	r._itemId = 0
	r._itemQuantity = 0
	r._options = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemQuantityUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.quantity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemQuantityUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemQuantityUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuQuantities is SkuQuantities Setter
// 更新SKU库存时候的SKU库存对象；如果没有SKU或者不更新SKU库存，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！
func (r *TmallItemQuantityUpdateAPIRequest) SetSkuQuantities(_skuQuantities []UpdateSkuQuantity) error {
	r._skuQuantities = _skuQuantities
	r.Set("sku_quantities", _skuQuantities)
	return nil
}

// GetSkuQuantities SkuQuantities Getter
func (r TmallItemQuantityUpdateAPIRequest) GetSkuQuantities() []UpdateSkuQuantity {
	return r._skuQuantities
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallItemQuantityUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemQuantityUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetItemQuantity is ItemQuantity Setter
// 商品库存数；增量编辑方式支持正数、负数。（无SKU商品使用这个字段）
func (r *TmallItemQuantityUpdateAPIRequest) SetItemQuantity(_itemQuantity int64) error {
	r._itemQuantity = _itemQuantity
	r.Set("item_quantity", _itemQuantity)
	return nil
}

// GetItemQuantity ItemQuantity Getter
func (r TmallItemQuantityUpdateAPIRequest) GetItemQuantity() int64 {
	return r._itemQuantity
}

// SetOptions is Options Setter
// 商品库存更新时候的可选参数
func (r *TmallItemQuantityUpdateAPIRequest) SetOptions(_options *UpdateItemQuantityOption) error {
	r._options = _options
	r.Set("options", _options)
	return nil
}

// GetOptions Options Getter
func (r TmallItemQuantityUpdateAPIRequest) GetOptions() *UpdateItemQuantityOption {
	return r._options
}

var poolTmallItemQuantityUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemQuantityUpdateRequest()
	},
}

// GetTmallItemQuantityUpdateRequest 从 sync.Pool 获取 TmallItemQuantityUpdateAPIRequest
func GetTmallItemQuantityUpdateAPIRequest() *TmallItemQuantityUpdateAPIRequest {
	return poolTmallItemQuantityUpdateAPIRequest.Get().(*TmallItemQuantityUpdateAPIRequest)
}

// ReleaseTmallItemQuantityUpdateAPIRequest 将 TmallItemQuantityUpdateAPIRequest 放入 sync.Pool
func ReleaseTmallItemQuantityUpdateAPIRequest(v *TmallItemQuantityUpdateAPIRequest) {
	v.Reset()
	poolTmallItemQuantityUpdateAPIRequest.Put(v)
}
