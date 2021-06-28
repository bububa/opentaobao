package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询服务商下的商家是否开通了改地址 APIResponse
alibaba.modifyaddress.isv.bindseller.check

鉴权服务商和商家的绑定关系，看商家是否开通了改地址
1. 没有授权
2. 已与当前appkey签约
3. 没有签约
4. 与其他服务商软件签约，如果是同一个isv name，返回appkey，否则不返回。
*/
type AlibabaModifyaddressIsvBindsellerCheckAPIResponse struct {
    model.CommonResponse
    AlibabaModifyaddressIsvBindsellerCheckResponse
}

type AlibabaModifyaddressIsvBindsellerCheckResponse struct {
    XMLName xml.Name `xml:"alibaba_modifyaddress_isv_bindseller_check_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // model
    
    Model   *CheckSellerChooseErpResponse `json:"model,omitempty" xml:"model,omitempty"`

    
}
