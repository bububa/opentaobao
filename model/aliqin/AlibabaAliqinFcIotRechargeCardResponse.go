package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
按终端号订购增值业务 APIResponse
alibaba.aliqin.fc.iot.rechargeCard

按终端号订购增值业务
*/
type AlibabaAliqinFcIotRechargeCardAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcIotRechargeCardResponse
}

type AlibabaAliqinFcIotRechargeCardResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_rechargeCard_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaAliqinFcIotRechargeCardResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
