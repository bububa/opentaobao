package aliqin

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物联卡信息查询 APIResponse
alibaba.aliqin.fc.iot.cardInfo

物联卡信息查询
*/
type AlibabaAliqinFcIotCardInfoAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFcIotCardInfoResponse `json:"alibaba_aliqin_fc_iot_cardInfo_response,omitempty"`
}

type AlibabaAliqinFcIotCardInfoResponse struct {

    // 结果对象
    Result   *AlibabaAliqinFcIotCardInfoResult `json:"result,omitempty"`

}
