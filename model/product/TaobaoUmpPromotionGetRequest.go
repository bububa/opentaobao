package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品优惠详情查询 API请求
taobao.ump.promotion.get

商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
*/
type TaobaoUmpPromotionGetAPIRequest struct {
    model.Params
    // 商品id
    _itemId   int64
}

// 初始化TaobaoUmpPromotionGetAPIRequest对象
func NewTaobaoUmpPromotionGetRequest() *TaobaoUmpPromotionGetAPIRequest{
    return &TaobaoUmpPromotionGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpPromotionGetAPIRequest) GetApiMethodName() string {
    return "taobao.ump.promotion.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpPromotionGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoUmpPromotionGetAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoUmpPromotionGetAPIRequest) GetItemId() int64 {
    return r._itemId
}
