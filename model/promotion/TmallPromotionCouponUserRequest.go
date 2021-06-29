package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户信息查询接口 API请求
tmall.promotion.coupon.user

开发给外部合作商（例如：苏宁），通过会员付款码获得会员nick
*/
type TmallPromotionCouponUserRequest struct {
    model.Params
    // 例如：suning
    bizType   string
    // 会员付款码
    payCode   string
    // 扩展字段
    extra   string
}

// 初始化TmallPromotionCouponUserRequest对象
func NewTmallPromotionCouponUserRequest() *TmallPromotionCouponUserRequest{
    return &TmallPromotionCouponUserRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallPromotionCouponUserRequest) GetApiMethodName() string {
    return "tmall.promotion.coupon.user"
}

// IRequest interface 方法, 获取API参数
func (r TmallPromotionCouponUserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 例如：suning
func (r *TmallPromotionCouponUserRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TmallPromotionCouponUserRequest) GetBizType() string {
    return r.bizType
}
// PayCode Setter
// 会员付款码
func (r *TmallPromotionCouponUserRequest) SetPayCode(payCode string) error {
    r.payCode = payCode
    r.Set("pay_code", payCode)
    return nil
}

// PayCode Getter
func (r TmallPromotionCouponUserRequest) GetPayCode() string {
    return r.payCode
}
// Extra Setter
// 扩展字段
func (r *TmallPromotionCouponUserRequest) SetExtra(extra string) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

// Extra Getter
func (r TmallPromotionCouponUserRequest) GetExtra() string {
    return r.extra
}
