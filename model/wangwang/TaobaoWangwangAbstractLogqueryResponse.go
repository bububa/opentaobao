package wangwang

import (
    "github.com/bububa/opentaobao/model"
)

/* 
模糊聊天记录查询 APIResponse
taobao.wangwang.abstract.logquery

模糊聊天记录查询
*/
type TaobaoWangwangAbstractLogqueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWangwangAbstractLogqueryResponse `json:"wangwang_abstract_logquery_response,omitempty"` 
    TaobaoWangwangAbstractLogqueryResponse
}

/* model for simplify = false
type TaobaoWangwangAbstractLogqueryResponse struct {

    // 0或-1，表示错误或正确，错误时有错误信息.<br/>为-1时就只有error_msg字段是有效的。
    
    RetCode   int64 `json:"ret_code,omitempty"`
    

    // 例如单词长度太长等。<br/>当ret_code不为0时这个值存在。
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 卖家id
    
    FromId   string `json:"from_id,omitempty"`
    

    // 买家id
    
    ToId   string `json:"to_id,omitempty"`
    

    // 当direction为1时有效，url列表
    
    UrlLists  struct {
        UrlList  []UrlList `json:"url_list,omitempty"`
    } `json:"url_lists,omitempty"`
    

    // 当direction为1时有效，关键词统计列表
    
    WordCountLists  struct {
        WordCountList  []WordCountList `json:"word_count_list,omitempty"`
    } `json:"word_count_lists,omitempty"`
    

    // 0或1<br/>其他返回时，是由于用户名等参数设置有误等引起的远端服务错误
    
    IsSliced   int64 `json:"is_sliced,omitempty"`
    

    // 当is_sliced为1时有效
    
    NextKey   string `json:"next_key,omitempty"`
    

    // 消息列表
    
    MsgLists  struct {
        MsgList  []MsgList `json:"msg_list,omitempty"`
    } `json:"msg_lists,omitempty"`
    

}
*/

type TaobaoWangwangAbstractLogqueryResponse struct {

    // 0或-1，表示错误或正确，错误时有错误信息.<br/>为-1时就只有error_msg字段是有效的。
    RetCode   int64 `json:"ret_code,omitempty"`

    // 例如单词长度太长等。<br/>当ret_code不为0时这个值存在。
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 卖家id
    FromId   string `json:"from_id,omitempty"`

    // 买家id
    ToId   string `json:"to_id,omitempty"`

    // 当direction为1时有效，url列表
    UrlLists   []UrlList `json:"url_lists,omitempty"`

    // 当direction为1时有效，关键词统计列表
    WordCountLists   []WordCountList `json:"word_count_lists,omitempty"`

    // 0或1<br/>其他返回时，是由于用户名等参数设置有误等引起的远端服务错误
    IsSliced   int64 `json:"is_sliced,omitempty"`

    // 当is_sliced为1时有效
    NextKey   string `json:"next_key,omitempty"`

    // 消息列表
    MsgLists   []MsgList `json:"msg_lists,omitempty"`

}
