package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据号码tts单呼 APIResponse
alibaba.aliqin.ta.number.singlecallbytts

将语音验证码和语音通知发布至聚石塔渠道
*/
type AlibabaAliqinTaNumberSinglecallbyttsAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinTaNumberSinglecallbyttsResponse `json:"alibaba_aliqin_ta_number_singlecallbytts_response,omitempty"`
}

type AlibabaAliqinTaNumberSinglecallbyttsResponse struct {

    // 接口返回model
    Result   *AlibabaAliqinTaNumberSinglecallbyttsResult `json:"result,omitempty"`

}
