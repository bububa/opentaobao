package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
openim群聊天记录导入 APIResponse
taobao.openim.tribelogs.import

openim群聊天记录导入
*/
type TaobaoOpenimTribelogsImportAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimTribelogsImportResponse `json:"openim_tribelogs_import_response,omitempty"` 
    TaobaoOpenimTribelogsImportResponse
}

/* model for simplify = false
type TaobaoOpenimTribelogsImportResponse struct {

    // 错误码
    
    Ret   int64 `json:"ret,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty"`
    

    // 是否成功
    
    Succ   bool `json:"succ,omitempty"`
    

}
*/

type TaobaoOpenimTribelogsImportResponse struct {

    // 错误码
    Ret   int64 `json:"ret,omitempty"`

    // 错误信息
    Message   string `json:"message,omitempty"`

    // 是否成功
    Succ   bool `json:"succ,omitempty"`

}
