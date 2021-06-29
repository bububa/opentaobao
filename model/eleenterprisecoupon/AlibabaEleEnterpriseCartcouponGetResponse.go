package eleenterprisecoupon

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取下单可用的优惠券 APIResponse
alibaba.ele.enterprise.cartcoupon.get

获取下单可用的优惠券
*/
type AlibabaEleEnterpriseCartcouponGetAPIResponse struct {
    model.CommonResponse
    AlibabaEleEnterpriseCartcouponGetResponse
}

type AlibabaEleEnterpriseCartcouponGetResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_enterprise_cartcoupon_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应code
    
    EnterpriseCode   string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`

    
    // 响应信息
    
    EnterpriseMsg   string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`

    
    // 返回值信息
    
    EnterpriseDatas   *MyCouponsResDto `json:"enterprise_datas,omitempty" xml:"enterprise_datas,omitempty"`

    
    // 请求id
    
    EnterpriseRequestid   string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`

    
}
