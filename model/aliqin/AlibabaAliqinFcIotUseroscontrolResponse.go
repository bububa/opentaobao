package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物联卡用户管理停开机功能 APIResponse
alibaba.aliqin.fc.iot.useroscontrol

物联网针对用户级管理停开支持
*/
type AlibabaAliqinFcIotUseroscontrolAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcIotUseroscontrolResponse
}

type AlibabaAliqinFcIotUseroscontrolResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_useroscontrol_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaAliqinFcIotUseroscontrolResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
