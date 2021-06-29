package middleclaims

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
国际化中台服务域接收理赔账单 API返回值 
alibaba.middle.claimsbill.receive

国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔打款的处理后，将该打款结果返回至服务域
*/
type AlibabaMiddleClaimsbillReceiveAPIResponse struct {
    model.CommonResponse
    AlibabaMiddleClaimsbillReceiveResponse
}

// 国际化中台服务域接收理赔账单 成功返回结果
type AlibabaMiddleClaimsbillReceiveResponse struct {
    XMLName xml.Name `xml:"alibaba_middle_claimsbill_receive_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果实体类
    Result   *AlibabaMiddleClaimsbillReceiveResult `json:"result,omitempty" xml:"result,omitempty"`
}
