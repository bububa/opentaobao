package wangwang

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词列表 APIResponse
taobao.wangwang.abstract.getwordlist

获取关键词列表，只支持json返回
*/
type TaobaoWangwangAbstractGetwordlistAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWangwangAbstractGetwordlistResponse `json:"wangwang_abstract_getwordlist_response,omitempty"` 
    TaobaoWangwangAbstractGetwordlistResponse
}

/* model for simplify = false
type TaobaoWangwangAbstractGetwordlistResponse struct {

    // 0或-1，表示错误或正确，错误时有错误信息
    
    RetCode   int64 `json:"ret_code,omitempty"`
    

    // 例如单词长度太长等，ret_code=-1才有
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 关键词列表，ret_code=0才有
    
    WordLists  struct {
        WordList  []WordList `json:"word_list,omitempty"`
    } `json:"word_lists,omitempty"`
    

}
*/

type TaobaoWangwangAbstractGetwordlistResponse struct {

    // 0或-1，表示错误或正确，错误时有错误信息
    RetCode   int64 `json:"ret_code,omitempty"`

    // 例如单词长度太长等，ret_code=-1才有
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 关键词列表，ret_code=0才有
    WordLists   []WordList `json:"word_lists,omitempty"`

}
