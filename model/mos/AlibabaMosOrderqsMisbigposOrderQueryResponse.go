package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大pos新选单退 API返回值 
alibaba.mos.orderqs.misbigpos.order.query

大pos新选单退
*/
type AlibabaMosOrderqsMisbigposOrderQueryAPIResponse struct {
    model.CommonResponse
    AlibabaMosOrderqsMisbigposOrderQueryResponse
}

// 大pos新选单退 成功返回结果
type AlibabaMosOrderqsMisbigposOrderQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_orderqs_misbigpos_order_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误信息
    ErrsMsg   string `json:"errs_msg,omitempty" xml:"errs_msg,omitempty"`
    // 错误码
    ErrsCode   int64 `json:"errs_code,omitempty" xml:"errs_code,omitempty"`
    // 调用是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 小票基本数据
    Sumstrs   []SumStr `json:"sumstrs,omitempty" xml:"sumstrs>sum_str,omitempty"`
    // 支付方式数据
    Payments   []Payment `json:"payments,omitempty" xml:"payments>payment,omitempty"`
    // 商品信息
    Goods   []Goods `json:"goods,omitempty" xml:"goods>goods,omitempty"`
    // 定向礼券相关
    Coupons   []Coupon `json:"coupons,omitempty" xml:"coupons>coupon,omitempty"`
    // 会员相关
    Vipinfo   *Vipinfo `json:"vipinfo,omitempty" xml:"vipinfo,omitempty"`
    // 退款相关
    Refunds   []Refund `json:"refunds,omitempty" xml:"refunds>refund,omitempty"`
    // 分摊相关
    Products   []Product `json:"products,omitempty" xml:"products>product,omitempty"`
    // 券相关
    CouponExterns   []Couponextern `json:"coupon_externs,omitempty" xml:"coupon_externs>couponextern,omitempty"`
    // 满返接口使用情况
    CouponServiceValid   bool `json:"coupon_service_valid,omitempty" xml:"coupon_service_valid,omitempty"`
}
