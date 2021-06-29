package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建分销和后端商品映射关系 API请求
taobao.fenxiao.product.map.add

创建分销和供应链商品映射关系。
*/
type TaobaoFenxiaoProductMapAddRequest struct {
    model.Params
    // 分销产品id。
    productId   int64
    // 后端商品id（如果当前分销产品没有sku和后端商品时需要指定）。
    scItemId   int64
    // 分销产品的sku id。逗号分隔，顺序需要保证与sc_item_ids一致（没有sku就不传）。
    skuIds   string
    // 在有sku的情况下，与各个sku对应的后端商品id列表。逗号分隔，顺序需要保证与sku_ids一致。
    scItemIds   string
    // 是否需要校验商家编码，true不校验，false校验。
    notCheckOuterCode   bool
}

// 初始化TaobaoFenxiaoProductMapAddRequest对象
func NewTaobaoFenxiaoProductMapAddRequest() *TaobaoFenxiaoProductMapAddRequest{
    return &TaobaoFenxiaoProductMapAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductMapAddRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.map.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductMapAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 分销产品id。
func (r *TaobaoFenxiaoProductMapAddRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TaobaoFenxiaoProductMapAddRequest) GetProductId() int64 {
    return r.productId
}
// ScItemId Setter
// 后端商品id（如果当前分销产品没有sku和后端商品时需要指定）。
func (r *TaobaoFenxiaoProductMapAddRequest) SetScItemId(scItemId int64) error {
    r.scItemId = scItemId
    r.Set("sc_item_id", scItemId)
    return nil
}

// ScItemId Getter
func (r TaobaoFenxiaoProductMapAddRequest) GetScItemId() int64 {
    return r.scItemId
}
// SkuIds Setter
// 分销产品的sku id。逗号分隔，顺序需要保证与sc_item_ids一致（没有sku就不传）。
func (r *TaobaoFenxiaoProductMapAddRequest) SetSkuIds(skuIds string) error {
    r.skuIds = skuIds
    r.Set("sku_ids", skuIds)
    return nil
}

// SkuIds Getter
func (r TaobaoFenxiaoProductMapAddRequest) GetSkuIds() string {
    return r.skuIds
}
// ScItemIds Setter
// 在有sku的情况下，与各个sku对应的后端商品id列表。逗号分隔，顺序需要保证与sku_ids一致。
func (r *TaobaoFenxiaoProductMapAddRequest) SetScItemIds(scItemIds string) error {
    r.scItemIds = scItemIds
    r.Set("sc_item_ids", scItemIds)
    return nil
}

// ScItemIds Getter
func (r TaobaoFenxiaoProductMapAddRequest) GetScItemIds() string {
    return r.scItemIds
}
// NotCheckOuterCode Setter
// 是否需要校验商家编码，true不校验，false校验。
func (r *TaobaoFenxiaoProductMapAddRequest) SetNotCheckOuterCode(notCheckOuterCode bool) error {
    r.notCheckOuterCode = notCheckOuterCode
    r.Set("not_check_outer_code", notCheckOuterCode)
    return nil
}

// NotCheckOuterCode Getter
func (r TaobaoFenxiaoProductMapAddRequest) GetNotCheckOuterCode() bool {
    return r.notCheckOuterCode
}
