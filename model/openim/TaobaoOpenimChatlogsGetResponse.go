package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
openim聊天记录查询接口 APIResponse
taobao.openim.chatlogs.get

查询openim账号聊天记录
*/
type TaobaoOpenimChatlogsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimChatlogsGetResponse `json:"openim_chatlogs_get_response,omitempty"` 
    TaobaoOpenimChatlogsGetResponse
}

/* model for simplify = false
type TaobaoOpenimChatlogsGetResponse struct {

    // 聊天记录查询结果
    
    Result  *struct {
        RoamingMessageResult  *RoamingMessageResult `json:"roaming_message_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoOpenimChatlogsGetResponse struct {

    // 聊天记录查询结果
    Result   *RoamingMessageResult `json:"result,omitempty"`

}
