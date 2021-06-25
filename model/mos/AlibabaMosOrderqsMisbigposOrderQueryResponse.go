package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
大pos新选单退 APIResponse
alibaba.mos.orderqs.misbigpos.order.query

大pos新选单退
*/
type AlibabaMosOrderqsMisbigposOrderQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMosOrderqsMisbigposOrderQueryResponse `json:"alibaba_mos_orderqs_misbigpos_order_query_response,omitempty"`
}

type AlibabaMosOrderqsMisbigposOrderQueryResponse struct {

    // 错误信息
    ErrsMsg   string `json:"errs_msg,omitempty"`

    // 错误码
    ErrsCode   int64 `json:"errs_code,omitempty"`

    // 调用是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 小票基本数据
    Sumstrs   []SumStr `json:"sumstrs,omitempty"`

    // 支付方式数据
    Payments   []Payment `json:"payments,omitempty"`

    // 商品信息
    Goods   []Goods `json:"goods,omitempty"`

    // 定向礼券相关
    Coupons   []Coupon `json:"coupons,omitempty"`

    // 会员相关
    Vipinfo   *Vipinfo `json:"vipinfo,omitempty"`

    // 退款相关
    Refunds   []Refund `json:"refunds,omitempty"`

    // 分摊相关
    Products   []Product `json:"products,omitempty"`

    // 券相关
    CouponExterns   []Couponextern `json:"coupon_externs,omitempty"`

    // 满返接口使用情况
    CouponServiceValid   bool `json:"coupon_service_valid,omitempty"`

}
