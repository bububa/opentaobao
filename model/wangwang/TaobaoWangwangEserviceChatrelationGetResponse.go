package wangwang

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聊天关系获取接口 APIResponse
taobao.wangwang.eservice.chatrelation.get

获取指定时间区间内的聊天关系（查询账号，和谁，在哪天说过话）。如：
A 和 B 在2016-09-01 和 2016-09-02 都说过话。以A为查询账号，则该接口将返回：
2016-09-01， B
2016-09-02， B
*/
type TaobaoWangwangEserviceChatrelationGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWangwangEserviceChatrelationGetResponse `json:"wangwang_eservice_chatrelation_get_response,omitempty"` 
    TaobaoWangwangEserviceChatrelationGetResponse
}

/* model for simplify = false
type TaobaoWangwangEserviceChatrelationGetResponse struct {

    // result
    
    Result  *struct {
        ChatRelationResult  *ChatRelationResult `json:"chat_relation_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWangwangEserviceChatrelationGetResponse struct {

    // result
    Result   *ChatRelationResult `json:"result,omitempty"`

}
