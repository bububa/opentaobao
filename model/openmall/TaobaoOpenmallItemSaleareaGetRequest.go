package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品可售区域 API请求
taobao.openmall.item.salearea.get

获取openmall商品的可售区域
*/
type TaobaoOpenmallItemSaleareaGetRequest struct {
    model.Params
    // 商品SKU
    _skuIds   string
    // 商品ID
    _itemId   int64
}

// 初始化TaobaoOpenmallItemSaleareaGetRequest对象
func NewTaobaoOpenmallItemSaleareaGetRequest() *TaobaoOpenmallItemSaleareaGetRequest{
    return &TaobaoOpenmallItemSaleareaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallItemSaleareaGetRequest) GetApiMethodName() string {
    return "taobao.openmall.item.salearea.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallItemSaleareaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuIds Setter
// 商品SKU
func (r *TaobaoOpenmallItemSaleareaGetRequest) SetSkuIds(_skuIds string) error {
    r._skuIds = _skuIds
    r.Set("sku_ids", _skuIds)
    return nil
}

// SkuIds Getter
func (r TaobaoOpenmallItemSaleareaGetRequest) GetSkuIds() string {
    return r._skuIds
}
// ItemId Setter
// 商品ID
func (r *TaobaoOpenmallItemSaleareaGetRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpenmallItemSaleareaGetRequest) GetItemId() int64 {
    return r._itemId
}
