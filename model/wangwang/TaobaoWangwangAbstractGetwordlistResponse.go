package wangwang

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词列表 APIResponse
taobao.wangwang.abstract.getwordlist

获取关键词列表，只支持json返回
*/
type TaobaoWangwangAbstractGetwordlistAPIResponse struct {
    model.CommonResponse
    TaobaoWangwangAbstractGetwordlistResponse
}

type TaobaoWangwangAbstractGetwordlistResponse struct {
    XMLName xml.Name `xml:"wangwang_abstract_getwordlist_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 0或-1，表示错误或正确，错误时有错误信息
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 例如单词长度太长等，ret_code=-1才有
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
    // 关键词列表，ret_code=0才有
    
    WordLists   []WordList `json:"word_lists,omitempty" xml:"word_lists>word_list,omitempty"`
    
    
}
