package shop

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取优惠券信息 APIResponse
alibaba.data.coupon.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
type AlibabaDataCouponGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_data_coupon_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // unnamed
    
    Unnamed   string `json:"unnamed,omitempty" xml:"