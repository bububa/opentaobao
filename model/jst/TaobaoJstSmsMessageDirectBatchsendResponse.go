package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔新短信发送接口 APIResponse
taobao.jst.sms.message.direct.batchsend

聚石塔所见即所得的短信发送接口
*/
type TaobaoJstSmsMessageDirectBatchsendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJstSmsMessageDirectBatchsendResponse `json:"jst_sms_message_direct_batchsend_response,omitempty"` 
    TaobaoJstSmsMessageDirectBatchsendResponse
}

/* model for simplify = false
type TaobaoJstSmsMessageDirectBatchsendResponse struct {

    // 短信回执码
    
    Module   string `json:"module,omitempty"`
    

}
*/

type TaobaoJstSmsMessageDirectBatchsendResponse struct {

    // 短信回执码
    Module   string `json:"module,omitempty"`

}
