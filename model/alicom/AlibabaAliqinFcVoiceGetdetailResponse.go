package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取呼叫详情 APIResponse
alibaba.aliqin.fc.voice.getdetail

通过呼叫id获取呼叫相关的数据
*/
type AlibabaAliqinFcVoiceGetdetailAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFcVoiceGetdetailResponse `json:"alibaba_aliqin_fc_voice_getdetail_response,omitempty"` 
    AlibabaAliqinFcVoiceGetdetailResponse
}

/* model for simplify = false
type AlibabaAliqinFcVoiceGetdetailResponse struct {

    // 错误码
    
    AlicomCode   string `json:"alicom_code,omitempty"`
    

    // 错误信息
    
    AlicomMsg   string `json:"alicom_msg,omitempty"`
    

    // 请求是否成功
    
    AlicomSuccess   bool `json:"alicom_success,omitempty"`
    

    // 返回值，在没有结果时为空。recordFile表示的是录音文件地址
    
    Model   string `json:"model,omitempty"`
    

}
*/

type AlibabaAliqinFcVoiceGetdetailResponse struct {

    // 错误码
    AlicomCode   string `json:"alicom_code,omitempty"`

    // 错误信息
    AlicomMsg   string `json:"alicom_msg,omitempty"`

    // 请求是否成功
    AlicomSuccess   bool `json:"alicom_success,omitempty"`

    // 返回值，在没有结果时为空。recordFile表示的是录音文件地址
    Model   string `json:"model,omitempty"`

}
