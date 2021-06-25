package wangwang

import (
    "github.com/bububa/opentaobao/model"
)

/* 
增加关键词 APIResponse
taobao.wangwang.abstract.addword

增加关键词，只支持json返回
*/
type TaobaoWangwangAbstractAddwordAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWangwangAbstractAddwordResponse `json:"taobao_wangwang_abstract_addword_response,omitempty"`
}

type TaobaoWangwangAbstractAddwordResponse struct {

    // 0或-1，表示错误或正确，错误时有错误信息
    RetCode   int64 `json:"ret_code,omitempty"`

    // 例如单词长度太长等，当ret_code=-1时才有这项
    ErrorMsg   string `json:"error_msg,omitempty"`

}
