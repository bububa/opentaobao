package aliqin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
短信发送记录查询 APIResponse
alibaba.aliqin.fc.sms.num.query

短信发送记录查询。
*/
type AlibabaAliqinFcSmsNumQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_fc_sms_num_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 当前页码
    
    CurrentPage   int64 `json:"current_page,omitempty" xml:"