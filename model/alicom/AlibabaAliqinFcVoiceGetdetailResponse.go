package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取呼叫详情 APIResponse
alibaba.aliqin.fc.voice.getdetail

通过呼叫id获取呼叫相关的数据
*/
type AlibabaAliqinFcVoiceGetdetailAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_fc_voice_getdetail_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    AlicomCode   string `json:"alicom_code,omitempty" xml:"