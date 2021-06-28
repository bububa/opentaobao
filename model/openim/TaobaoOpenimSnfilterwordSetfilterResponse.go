package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词过滤 APIResponse
taobao.openim.snfilterword.setfilter

设置openim关键词过滤
*/
type TaobaoOpenimSnfilterwordSetfilterAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"openim_snfilterword_setfilter_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 成功
    
    Errid   int64 `json:"errid,omitempty" xml:"