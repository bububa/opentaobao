package wangwang

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
模糊聊天记录查询 APIResponse
taobao.wangwang.abstract.logquery

模糊聊天记录查询
*/
type TaobaoWangwangAbstractLogqueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wangwang_abstract_logquery_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 0或-1，表示错误或正确，错误时有错误信息.<br/>为-1时就只有error_msg字段是有效的。
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"