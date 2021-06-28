package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔新短信发送接口 APIResponse
taobao.jst.sms.message.direct.batchsend

聚石塔所见即所得的短信发送接口
*/
type TaobaoJstSmsMessageDirectBatchsendAPIResponse struct {
    model.CommonResponse
    TaobaoJstSmsMessageDirectBatchsendResponse
}

type TaobaoJstSmsMessageDirectBatchsendResponse struct {
    XMLName xml.Name `xml:"jst_sms_message_direct_batchsend_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 短信回执码
    
    Module   string `json:"module,omitempty" xml:"module,omitempty"`

    
}
