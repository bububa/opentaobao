package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫商品/SKU库存更新接口 API请求
tmall.item.quantity.update

天猫商品/SKU库存更新接口；支持商品库存更新；支持同一商品下的SKU批量更新。
*/
type TmallItemQuantityUpdateRequest struct {
    model.Params
    // 商品id
    _itemId   int64
    // 更新SKU库存时候的SKU库存对象；如果没有SKU或者不更新SKU库存，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！
    _skuQuantities   []UpdateSkuQuantity
    // 商品库存更新时候的可选参数
    _options   *UpdateItemQuantityOption
    // 商品库存数；增量编辑方式支持正数、负数。（无SKU商品使用这个字段）
    _itemQuantity   int64
}

// 初始化TmallItemQuantityUpdateRequest对象
func NewTmallItemQuantityUpdateRequest() *TmallItemQuantityUpdateRequest{
    return &TmallItemQuantityUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemQuantityUpdateRequest) GetApiMethodName() string {
    return "tmall.item.quantity.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemQuantityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TmallItemQuantityUpdateRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallItemQuantityUpdateRequest) GetItemId() int64 {
    return r._itemId
}
// SkuQuantities Setter
// 更新SKU库存时候的SKU库存对象；如果没有SKU或者不更新SKU库存，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！
func (r *TmallItemQuantityUpdateRequest) SetSkuQuantities(_skuQuantities []UpdateSkuQuantity) error {
    r._skuQuantities = _skuQuantities
    r.Set("sku_quantities", _skuQuantities)
    return nil
}

// SkuQuantities Getter
func (r TmallItemQuantityUpdateRequest) GetSkuQuantities() []UpdateSkuQuantity {
    return r._skuQuantities
}
// Options Setter
// 商品库存更新时候的可选参数
func (r *TmallItemQuantityUpdateRequest) SetOptions(_options *UpdateItemQuantityOption) error {
    r._options = _options
    r.Set("options", _options)
    return nil
}

// Options Getter
func (r TmallItemQuantityUpdateRequest) GetOptions() *UpdateItemQuantityOption {
    return r._options
}
// ItemQuantity Setter
// 商品库存数；增量编辑方式支持正数、负数。（无SKU商品使用这个字段）
func (r *TmallItemQuantityUpdateRequest) SetItemQuantity(_itemQuantity int64) error {
    r._itemQuantity = _itemQuantity
    r.Set("item_quantity", _itemQuantity)
    return nil
}

// ItemQuantity Getter
func (r TmallItemQuantityUpdateRequest) GetItemQuantity() int64 {
    return r._itemQuantity
}
