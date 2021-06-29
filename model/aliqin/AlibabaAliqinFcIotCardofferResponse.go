package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询物联网卡上订购的offer APIResponse
alibaba.aliqin.fc.iot.cardoffer

查询物联网卡上订购的offer
*/
type AlibabaAliqinFcIotCardofferAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcIotCardofferResponse
}

type AlibabaAliqinFcIotCardofferResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_cardoffer_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果对象
    
    Result   *AlibabaAliqinFcIotCardofferResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
