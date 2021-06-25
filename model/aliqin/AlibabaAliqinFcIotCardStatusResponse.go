package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物联卡状态查询 APIResponse
alibaba.aliqin.fc.iot.cardStatus

物联卡状态查询
*/
type AlibabaAliqinFcIotCardStatusAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFcIotCardStatusResponse `json:"alibaba_aliqin_fc_iot_cardStatus_response,omitempty"`
}

type AlibabaAliqinFcIotCardStatusResponse struct {

    // 结果对象
    Result   *AlibabaAliqinFcIotCardStatusResult `json:"result,omitempty"`

}
