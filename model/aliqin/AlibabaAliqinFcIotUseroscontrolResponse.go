package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物联卡用户管理停开机功能 APIResponse
alibaba.aliqin.fc.iot.useroscontrol

物联网针对用户级管理停开支持
*/
type AlibabaAliqinFcIotUseroscontrolAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFcIotUseroscontrolResponse `json:"alibaba_aliqin_fc_iot_useroscontrol_response,omitempty"`
}

type AlibabaAliqinFcIotUseroscontrolResponse struct {

    // result
    Result   *AlibabaAliqinFcIotUseroscontrolResult `json:"result,omitempty"`

}
