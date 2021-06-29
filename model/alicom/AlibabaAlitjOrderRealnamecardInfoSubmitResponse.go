package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里实人认证卡片信息回传 APIResponse
alibaba.alitj.order.realnamecard.info.submit

阿里实人认证卡片信息回传。ISP相关商家在线对接阿里通信的实人认证功能，在线提交订单对应运营商的合约订购相关信息，以便完成在线使用实人认证功能。
*/
type AlibabaAlitjOrderRealnamecardInfoSubmitAPIResponse struct {
    model.CommonResponse
    AlibabaAlitjOrderRealnamecardInfoSubmitResponse
}

type AlibabaAlitjOrderRealnamecardInfoSubmitResponse struct {
    XMLName xml.Name `xml:"alibaba_alitj_order_realnamecard_info_submit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
