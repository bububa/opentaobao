package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交专车优惠券活动 API请求
alibaba.alihouse.newhome.coupon.submit

提交专车优惠券活动
*/
type AlibabaAlihouseNewhomeCouponSubmitAPIRequest struct {
    model.Params
    // 打车券活动对象
    _couponDto   *MarketingCouponDTO
}

// 初始化AlibabaAlihouseNewhomeCouponSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeCouponSubmitRequest() *AlibabaAlihouseNewhomeCouponSubmitAPIRequest{
    return &AlibabaAlihouseNewhomeCouponSubmitAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeCouponSubmitAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.coupon.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeCouponSubmitAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CouponDto Setter
// 打车券活动对象
func (r *AlibabaAlihouseNewhomeCouponSubmitAPIRequest) SetCouponDto(_couponDto *MarketingCouponDTO) error {
    r._couponDto = _couponDto
    r.Set("coupon_dto", _couponDto)
    return nil
}

// CouponDto Getter
func (r AlibabaAlihouseNewhomeCouponSubmitAPIRequest) GetCouponDto() *MarketingCouponDTO {
    return r._couponDto
}
