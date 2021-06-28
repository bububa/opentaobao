package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
短信查询 APIResponse
alibaba.aliqin.ta.sms.num.query

查询短信发送揭露
*/
type AlibabaAliqinTaSmsNumQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_ta_sms_num_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 当前页码
    
    CurrentPage   int64 `json:"current_page,omitempty" xml:"