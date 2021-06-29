package eleenterprisecoupon

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户优惠券 API返回值 
alibaba.ele.enterprise.coupon.get

获取用户优惠券
*/
type AlibabaEleEnterpriseCouponGetAPIResponse struct {
    model.CommonResponse
    AlibabaEleEnterpriseCouponGetResponse
}

// 获取用户优惠券 成功返回结果
type AlibabaEleEnterpriseCouponGetResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_enterprise_coupon_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 响应code
    EnterpriseCode   string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
    // 响应信息
    EnterpriseMsg   string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
    // 返回值信息
    EnterpriseDatas   []MyCouponsResDTO `json:"enterprise_datas,omitempty" xml:"enterprise_datas>my_coupons_res_dto,omitempty"`
    // 请求id
    EnterpriseRequestid   string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
}
