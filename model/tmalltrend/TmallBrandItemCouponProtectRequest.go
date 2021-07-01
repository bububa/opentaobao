package tmalltrend

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全域新品店铺优惠券免除 API请求
tmall.brand.item.coupon.protect

全域新品店铺优惠券免除申请打标接口
*/
type TmallBrandItemCouponProtectAPIRequest struct {
    model.Params
    // 天猫商品id
    _itemId   int64
    // 店铺优惠券新品保护期档次:PERIOD_0D("0天"),     PERIOD_7D("7天"),     PERIOD_14D("14天"),     PERIOD_21D("21天")
    _protectionPeriod   string
    // 天猫品牌id
    _brandId   int64
}

// 初始化TmallBrandItemCouponProtectAPIRequest对象
func NewTmallBrandItemCouponProtectRequest() *TmallBrandItemCouponProtectAPIRequest{
    return &TmallBrandItemCouponProtectAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallBrandItemCouponProtectAPIRequest) GetApiMethodName() string {
    return "tmall.brand.item.coupon.protect"
}

// IRequest interface 方法, 获取API参数
func (r TmallBrandItemCouponProtectAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 天猫商品id
func (r *TmallBrandItemCouponProtectAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallBrandItemCouponProtectAPIRequest) GetItemId() int64 {
    return r._itemId
}
// ProtectionPeriod Setter
// 店铺优惠券新品保护期档次:PERIOD_0D("0天"),     PERIOD_7D("7天"),     PERIOD_14D("14天"),     PERIOD_21D("21天")
func (r *TmallBrandItemCouponProtectAPIRequest) SetProtectionPeriod(_protectionPeriod string) error {
    r._protectionPeriod = _protectionPeriod
    r.Set("protection_period", _protectionPeriod)
    return nil
}

// ProtectionPeriod Getter
func (r TmallBrandItemCouponProtectAPIRequest) GetProtectionPeriod() string {
    return r._protectionPeriod
}
// BrandId Setter
// 天猫品牌id
func (r *TmallBrandItemCouponProtectAPIRequest) SetBrandId(_brandId int64) error {
    r._brandId = _brandId
    r.Set("brand_id", _brandId)
    return nil
}

// BrandId Getter
func (r TmallBrandItemCouponProtectAPIRequest) GetBrandId() int64 {
    return r._brandId
}
