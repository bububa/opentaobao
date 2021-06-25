package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物联网绑定/换绑API APIResponse
alibaba.aliqin.fc.iot.modbind

支持用户的设备的换绑和解绑操作
*/
type AlibabaAliqinFcIotModbindAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFcIotModbindResponse `json:"alibaba_aliqin_fc_iot_modbind_response,omitempty"`
}

type AlibabaAliqinFcIotModbindResponse struct {

    // result
    Result   *AlibabaAliqinFcIotModbindResult `json:"result,omitempty"`

}
