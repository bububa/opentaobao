package wangwang

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除关键词 APIResponse
taobao.wangwang.abstract.deleteword

删除关键词，只支持json返回
*/
type TaobaoWangwangAbstractDeletewordAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWangwangAbstractDeletewordResponse `json:"wangwang_abstract_deleteword_response,omitempty"` 
    TaobaoWangwangAbstractDeletewordResponse
}

/* model for simplify = false
type TaobaoWangwangAbstractDeletewordResponse struct {

    // 0或-1，表示错误或正确，错误时有错误信息
    
    RetCode   int64 `json:"ret_code,omitempty"`
    

    // 例如单词长度太长等
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

type TaobaoWangwangAbstractDeletewordResponse struct {

    // 0或-1，表示错误或正确，错误时有错误信息
    RetCode   int64 `json:"ret_code,omitempty"`

    // 例如单词长度太长等
    ErrorMsg   string `json:"error_msg,omitempty"`

}
