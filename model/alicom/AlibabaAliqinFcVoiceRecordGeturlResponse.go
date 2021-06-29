package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
录音文件下载 API返回值 
alibaba.aliqin.fc.voice.record.geturl

完成录音文件的下载地址获取操作
*/
type AlibabaAliqinFcVoiceRecordGeturlAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFcVoiceRecordGeturlResponse
}

// 录音文件下载 成功返回结果
type AlibabaAliqinFcVoiceRecordGeturlResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_fc_voice_record_geturl_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaAliqinFcVoiceRecordGeturlResult `json:"result,omitempty" xml:"result,omitempty"`
}
