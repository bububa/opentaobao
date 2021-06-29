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
type TaobaoTbkSkuBestCouponRequest struct {
    model.Params
    // 商品ID
    _itemId   int64
    // 商品对应的skuId
    _skuId   int64
}

// 初始化TaobaoTbkSkuBestCouponRequest对象
func NewTaobaoTbkSkuBestCouponRequest() *TaobaoTbkSkuBestCouponRequest{
    return &TaobaoTbkSkuBestCouponRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkSkuBestCouponRequest) GetApiMethodName() string {
    return "taobao.tbk.sku.best.coupon"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkSkuBestCouponRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *TaobaoTbkSkuBestCouponRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoTbkSkuBestCouponRequest) GetItemId() int64 {
    return r._itemId
}
// SkuId Setter
// 商品对应的skuId
func (r *TaobaoTbkSkuBestCouponRequest) SetSkuId(_skuId int64) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TaobaoTbkSkuBestCouponRequest) GetSkuId() int64 {
    return r._skuId
}
