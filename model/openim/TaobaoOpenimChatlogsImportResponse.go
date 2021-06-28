package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
openim单聊消息导入 APIResponse
taobao.openim.chatlogs.import

提供openim账号的聊天消息导入功能
*/
type TaobaoOpenimChatlogsImportAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimChatlogsImportResponse `json:"openim_chatlogs_import_response,omitempty"` 
    TaobaoOpenimChatlogsImportResponse
}

/* model for simplify = false
type TaobaoOpenimChatlogsImportResponse struct {

    // 错误码
    
    Ret   int64 `json:"ret,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty"`
    

    // 是否成功
    
    Succ   bool `json:"succ,omitempty"`
    

}
*/

type TaobaoOpenimChatlogsImportResponse struct {

    // 错误码
    Ret   int64 `json:"ret,omitempty"`

    // 错误信息
    Message   string `json:"message,omitempty"`

    // 是否成功
    Succ   bool `json:"succ,omitempty"`

}
