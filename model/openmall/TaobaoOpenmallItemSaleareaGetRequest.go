package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品可售区域 APIRequest
taobao.openmall.item.salearea.get

获取openmall商品的可售区域
*/
type TaobaoOpenmallItemSaleareaGetRequest struct {
    model.Params

    // 商品SKU
    skuIds   string 

    // 商品ID
    itemId   int64 

}

func NewTaobaoOpenmallItemSaleareaGetRequest() *TaobaoOpenmallItemSaleareaGetRequest{
    return &TaobaoOpenmallItemSaleareaGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallItemSaleareaGetRequest) GetApiMethodName() string {
    return "taobao.openmall.item.salearea.get"
}

func (r TaobaoOpenmallItemSaleareaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallItemSaleareaGetRequest) SetSkuIds(skuIds string) error {
    r.skuIds = skuIds
    r.Set("sku_ids", skuIds)
    return nil
}

func (r TaobaoOpenmallItemSaleareaGetRequest) GetSkuIds() string {
    return r.skuIds
}

func (r *TaobaoOpenmallItemSaleareaGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoOpenmallItemSaleareaGetRequest) GetItemId() int64 {
    return r.itemId
}

