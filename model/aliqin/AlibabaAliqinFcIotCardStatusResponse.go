package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物联卡状态查询 API返回值 
alibaba.aliqin.fc.iot.cardStatus

物联卡状态查询
*/
type AlibabaAliqinFcIotCardStatusAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcIotCardStatusResponse
}

// 物联卡状态查询 成功返回结果
type AlibabaAliqinFcIotCardStatusResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_cardStatus_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果对象
    Result   *AlibabaAliqinFcIotCardStatusResult `json:"result,omitempty" xml:"result,omitempty"`
}
