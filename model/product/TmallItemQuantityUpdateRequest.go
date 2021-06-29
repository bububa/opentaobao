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
    itemId   int64
    // 更新SKU库存时候的SKU库存对象；如果没有SKU或者不更新SKU库存，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！
    skuQuantities   []UpdateSkuQuantity
    // 商品库存更新时候的可选参数
    options   *UpdateItemQuantityOption
    // 商品库存数；增量编辑方式支持正数、负数。（无SKU商品使用这个字段）
    itemQuantity   int64
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
func (r *TmallItemQuantityUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TmallItemQuantityUpdateRequest) GetItemId() int64 {
    return r.itemId
}
// SkuQuantities Setter
// 更新SKU库存时候的SKU库存对象；如果没有SKU或者不更新SKU库存，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！
func (r *TmallItemQuantityUpdateRequest) SetSkuQuantities(skuQuantities []UpdateSkuQuantity) error {
    r.skuQuantities = skuQuantities
    r.Set("sku_quantities", skuQuantities)
    return nil
}

// SkuQuantities Getter
func (r TmallItemQuantityUpdateRequest) GetSkuQuantities() []UpdateSkuQuantity {
    return r.skuQuantities
}
// Options Setter
// 商品库存更新时候的可选参数
func (r *TmallItemQuantityUpdateRequest) SetOptions(options *UpdateItemQuantityOption) error {
    r.options = options
    r.Set("options", options)
    return nil
}

// Options Getter
func (r TmallItemQuantityUpdateRequest) GetOptions() *UpdateItemQuantityOption {
    return r.options
}
// ItemQuantity Setter
// 商品库存数；增量编辑方式支持正数、负数。（无SKU商品使用这个字段）
func (r *TmallItemQuantityUpdateRequest) SetItemQuantity(itemQuantity int64) error {
    r.itemQuantity = itemQuantity
    r.Set("item_quantity", itemQuantity)
    return nil
}

// ItemQuantity Getter
func (r TmallItemQuantityUpdateRequest) GetItemQuantity() int64 {
    return r.itemQuantity
}
