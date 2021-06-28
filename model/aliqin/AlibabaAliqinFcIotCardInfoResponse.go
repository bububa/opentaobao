package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物联卡信息查询 APIResponse
alibaba.aliqin.fc.iot.cardInfo

物联卡信息查询
*/
type AlibabaAliqinFcIotCardInfoAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcIotCardInfoResponse
}

type AlibabaAliqinFcIotCardInfoResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_cardInfo_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果对象
    
    Result   *AlibabaAliqinFcIotCardInfoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
