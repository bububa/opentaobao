package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
录音文件下载 APIResponse
alibaba.aliqin.fc.voice.record.geturl

完成录音文件的下载地址获取操作
*/
type AlibabaAliqinFcVoiceRecordGeturlAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFcVoiceRecordGeturlResponse `json:"alibaba_aliqin_fc_voice_record_geturl_response,omitempty"`
}

type AlibabaAliqinFcVoiceRecordGeturlResponse struct {

    // result
    Result   *AlibabaAliqinFcVoiceRecordGeturlResult `json:"result,omitempty"`

}
