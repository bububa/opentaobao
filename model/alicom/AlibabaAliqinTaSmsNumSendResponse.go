package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
短信发送 APIResponse
alibaba.aliqin.ta.sms.num.send

短信发送
*/
type AlibabaAliqinTaSmsNumSendAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_ta_sms_num_send_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *AlibabaAliqinTaSmsNumSendBizResult `json:"result,omitempty" xml:"