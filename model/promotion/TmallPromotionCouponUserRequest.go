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
    _bizType   string
    // 会员付款码
    _payCode   string
    // 扩展字段
    _extra   string
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
func (r *TmallPromotionCouponUserRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TmallPromotionCouponUserRequest) GetBizType() string {
    return r._bizType
}
// PayCode Setter
// 会员付款码
func (r *TmallPromotionCouponUserRequest) SetPayCode(_payCode string) error {
    r._payCode = _payCode
    r.Set("pay_code", _payCode)
    return nil
}

// PayCode Getter
func (r TmallPromotionCouponUserRequest) GetPayCode() string {
    return r._payCode
}
// Extra Setter
// 扩展字段
func (r *TmallPromotionCouponUserRequest) SetExtra(_extra string) error {
    r._extra = _extra
    r.Set("extra", _extra)
    return nil
}

// Extra Getter
func (r TmallPromotionCouponUserRequest) GetExtra() string {
    return r._extra
}
