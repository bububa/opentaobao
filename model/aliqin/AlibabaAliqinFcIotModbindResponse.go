package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物联网绑定/换绑API APIResponse
alibaba.aliqin.fc.iot.modbind

支持用户的设备的换绑和解绑操作
*/
type AlibabaAliqinFcIotModbindAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcIotModbindResponse
}

type AlibabaAliqinFcIotModbindResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_modbind_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaAliqinFcIotModbindResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
