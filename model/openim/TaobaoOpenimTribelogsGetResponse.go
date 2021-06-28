package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
openim 群聊天记录导出接口 APIResponse
taobao.openim.tribelogs.get

获取openim账号的群聊天记录
*/
type TaobaoOpenimTribelogsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimTribelogsGetResponse `json:"openim_tribelogs_get_response,omitempty"` 
    TaobaoOpenimTribelogsGetResponse
}

/* model for simplify = false
type TaobaoOpenimTribelogsGetResponse struct {

    // 错误码
    
    RetCode   int64 `json:"retCode,omitempty"`
    

    // 返回结构
    
    Data  *struct {
        TribeMessageResult  *TribeMessageResult `json:"tribe_message_result,omitempty"`
    } `json:"data,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty"`
    

    // 错误原因
    
    Reason   string `json:"reason,omitempty"`
    

    // 是否成功
    
    Succ   bool `json:"succ,omitempty"`
    

}
*/

type TaobaoOpenimTribelogsGetResponse struct {

    // 错误码
    RetCode   int64 `json:"retCode,omitempty"`

    // 返回结构
    Data   *TribeMessageResult `json:"data,omitempty"`

    // 错误信息
    Message   string `json:"message,omitempty"`

    // 错误原因
    Reason   string `json:"reason,omitempty"`

    // 是否成功
    Succ   bool `json:"succ,omitempty"`

}
