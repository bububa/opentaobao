package wangwang

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除关键词 APIResponse
taobao.wangwang.abstract.deleteword

删除关键词，只支持json返回
*/
type TaobaoWangwangAbstractDeletewordAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wangwang_abstract_deleteword_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 0或-1，表示错误或正确，错误时有错误信息
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"