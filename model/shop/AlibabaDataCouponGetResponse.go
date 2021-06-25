package shop

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取优惠券信息 APIResponse
alibaba.data.coupon.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
type AlibabaDataCouponGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaDataCouponGetResponse `json:"alibaba_data_coupon_get_response,omitempty"`
}

type AlibabaDataCouponGetResponse struct {

    // unnamed
    Unnamed   string `json:"unnamed,omitempty"`

}
