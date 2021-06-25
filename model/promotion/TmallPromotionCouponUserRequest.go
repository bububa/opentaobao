package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户信息查询接口 APIRequest
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

func NewTmallPromotionCouponUserRequest() *TmallPromotionCouponUserRequest{
    return &TmallPromotionCouponUserRequest{
        Params: model.NewParams(),
    }
}

func (r TmallPromotionCouponUserRequest) GetApiMethodName() string {
    return "tmall.promotion.coupon.user"
}

func (r TmallPromotionCouponUserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallPromotionCouponUserRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TmallPromotionCouponUserRequest) GetBizType() string {
    return r.bizType
}

func (r *TmallPromotionCouponUserRequest) SetPayCode(payCode string) error {
    r.payCode = payCode
    r.Set("pay_code", payCode)
    return nil
}

func (r TmallPromotionCouponUserRequest) GetPayCode() string {
    return r.payCode
}

func (r *TmallPromotionCouponUserRequest) SetExtra(extra string) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

func (r TmallPromotionCouponUserRequest) GetExtra() string {
    return r.extra
}

