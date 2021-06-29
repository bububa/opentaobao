package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品优惠价格查询 APIRequest
tmall.popupstore.item.discount.price

商品优惠价格查询
*/
type TmallPopupstoreItemDiscountPriceRequest struct {
    model.Params

    // 商品id列表
    itemIds   []int64 

}

func NewTmallPopupstoreItemDiscountPriceRequest() *TmallPopupstoreItemDiscountPriceRequest{
    return &TmallPopupstoreItemDiscountPriceRequest{
        Params: model.NewParams(),
    }
}

func (r TmallPopupstoreItemDiscountPriceRequest) GetApiMethodName() string {
    return "tmall.popupstore.item.discount.price"
}

func (r TmallPopupstoreItemDiscountPriceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallPopupstoreItemDiscountPriceRequest) SetItemIds(itemIds []int64) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

func (r TmallPopupstoreItemDiscountPriceRequest) GetItemIds() []int64 {
    return r.itemIds
}

