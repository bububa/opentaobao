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
type AlibabaAlihouseNewhomeCouponSubmitRequest struct {
    model.Params
    // 打车券活动对象
    _couponDto   *MarketingCouponDto
}

// 初始化AlibabaAlihouseNewhomeCouponSubmitRequest对象
func NewAlibabaAlihouseNewhomeCouponSubmitRequest() *AlibabaAlihouseNewhomeCouponSubmitRequest{
    return &AlibabaAlihouseNewhomeCouponSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeCouponSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.coupon.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeCouponSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CouponDto Setter
// 打车券活动对象
func (r *AlibabaAlihouseNewhomeCouponSubmitRequest) SetCouponDto(_couponDto *MarketingCouponDto) error {
    r._couponDto = _couponDto
    r.Set("coupon_dto", _couponDto)
    return nil
}

// CouponDto Getter
func (r AlibabaAlihouseNewhomeCouponSubmitRequest) GetCouponDto() *MarketingCouponDto {
    return r._couponDto
}
