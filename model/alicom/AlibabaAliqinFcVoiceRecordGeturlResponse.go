package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
录音文件下载 APIResponse
alibaba.aliqin.fc.voice.record.geturl

完成录音文件的下载地址获取操作
*/
type AlibabaAliqinFcVoiceRecordGeturlAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_fc_voice_record_geturl_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaAliqinFcVoiceRecordGeturlResult `json:"result,omitempty" xml:"