package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫商品/SKU库存更新接口 APIRequest
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

func NewTmallItemQuantityUpdateRequest() *TmallItemQuantityUpdateRequest{
    return &TmallItemQuantityUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemQuantityUpdateRequest) GetApiMethodName() string {
    return "tmall.item.quantity.update"
}

func (r TmallItemQuantityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemQuantityUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallItemQuantityUpdateRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TmallItemQuantityUpdateRequest) SetSkuQuantities(skuQuantities []UpdateSkuQuantity) error {
    r.skuQuantities = skuQuantities
    r.Set("sku_quantities", skuQuantities)
    return nil
}

func (r TmallItemQuantityUpdateRequest) GetSkuQuantities() []UpdateSkuQuantity {
    return r.skuQuantities
}

func (r *TmallItemQuantityUpdateRequest) SetOptions(options *UpdateItemQuantityOption) error {
    r.options = options
    r.Set("options", options)
    return nil
}

func (r TmallItemQuantityUpdateRequest) GetOptions() *UpdateItemQuantityOption {
    return r.options
}

func (r *TmallItemQuantityUpdateRequest) SetItemQuantity(itemQuantity int64) error {
    r.itemQuantity = itemQuantity
    r.Set("item_quantity", itemQuantity)
    return nil
}

func (r TmallItemQuantityUpdateRequest) GetItemQuantity() int64 {
    return r.itemQuantity
}

