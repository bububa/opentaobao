package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物联卡状态查询 APIResponse
alibaba.aliqin.fc.iot.cardStatus

物联卡状态查询
*/
type AlibabaAliqinFcIotCardStatusAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_fc_iot_cardStatus_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果对象
    
    Result   *AlibabaAliqinFcIotCardStatusResult `json:"result,omitempty" xml:"