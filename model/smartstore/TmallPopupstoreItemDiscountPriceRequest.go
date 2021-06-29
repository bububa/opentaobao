package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品优惠价格查询 API请求
tmall.popupstore.item.discount.price

商品优惠价格查询
*/
type TmallPopupstoreItemDiscountPriceRequest struct {
    model.Params
    // 商品id列表
    itemIds   []int64
}

// 初始化TmallPopupstoreItemDiscountPriceRequest对象
func NewTmallPopupstoreItemDiscountPriceRequest() *TmallPopupstoreItemDiscountPriceRequest{
    return &TmallPopupstoreItemDiscountPriceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallPopupstoreItemDiscountPriceRequest) GetApiMethodName() string {
    return "tmall.popupstore.item.discount.price"
}

// IRequest interface 方法, 获取API参数
func (r TmallPopupstoreItemDiscountPriceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemIds Setter
// 商品id列表
func (r *TmallPopupstoreItemDiscountPriceRequest) SetItemIds(itemIds []int64) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

// ItemIds Getter
func (r TmallPopupstoreItemDiscountPriceRequest) GetItemIds() []int64 {
    return r.itemIds
}
