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
type TmallBrandItemCouponProtectRequest struct {
    model.Params
    // 天猫商品id
    itemId   int64
    // 店铺优惠券新品保护期档次:PERIOD_0D("0天"),     PERIOD_7D("7天"),     PERIOD_14D("14天"),     PERIOD_21D("21天")
    protectionPeriod   string
    // 天猫品牌id
    brandId   int64
}

// 初始化TmallBrandItemCouponProtectRequest对象
func NewTmallBrandItemCouponProtectRequest() *TmallBrandItemCouponProtectRequest{
    return &TmallBrandItemCouponProtectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallBrandItemCouponProtectRequest) GetApiMethodName() string {
    return "tmall.brand.item.coupon.protect"
}

// IRequest interface 方法, 获取API参数
func (r TmallBrandItemCouponProtectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 天猫商品id
func (r *TmallBrandItemCouponProtectRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TmallBrandItemCouponProtectRequest) GetItemId() int64 {
    return r.itemId
}
// ProtectionPeriod Setter
// 店铺优惠券新品保护期档次:PERIOD_0D("0天"),     PERIOD_7D("7天"),     PERIOD_14D("14天"),     PERIOD_21D("21天")
func (r *TmallBrandItemCouponProtectRequest) SetProtectionPeriod(protectionPeriod string) error {
    r.protectionPeriod = protectionPeriod
    r.Set("protection_period", protectionPeriod)
    return nil
}

// ProtectionPeriod Getter
func (r TmallBrandItemCouponProtectRequest) GetProtectionPeriod() string {
    return r.protectionPeriod
}
// BrandId Setter
// 天猫品牌id
func (r *TmallBrandItemCouponProtectRequest) SetBrandId(brandId int64) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

// BrandId Getter
func (r TmallBrandItemCouponProtectRequest) GetBrandId() int64 {
    return r.brandId
}
