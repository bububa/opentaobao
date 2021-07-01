package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
sku维度最优优惠券信息 API请求
taobao.tbk.sku.best.coupon

根据itemid和skuid查询最优优惠券信息
*/
type TaobaoTbkSkuBestCouponAPIRequest struct {
    model.Params
    // 商品ID
    _itemId   int64
    // 商品对应的skuId
    _skuId   int64
}

// 初始化TaobaoTbkSkuBestCouponAPIRequest对象
func NewTaobaoTbkSkuBestCouponRequest() *TaobaoTbkSkuBestCouponAPIRequest{
    return &TaobaoTbkSkuBestCouponAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkSkuBestCouponAPIRequest) GetApiMethodName() string {
    return "taobao.tbk.sku.best.coupon"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkSkuBestCouponAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *TaobaoTbkSkuBestCouponAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoTbkSkuBestCouponAPIRequest) GetItemId() int64 {
    return r._itemId
}
// SkuId Setter
// 商品对应的skuId
func (r *TaobaoTbkSkuBestCouponAPIRequest) SetSkuId(_skuId int64) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TaobaoTbkSkuBestCouponAPIRequest) GetSkuId() int64 {
    return r._skuId
}
