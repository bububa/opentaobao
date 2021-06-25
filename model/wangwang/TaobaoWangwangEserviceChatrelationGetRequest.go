package wangwang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聊天关系获取接口 APIRequest
taobao.wangwang.eservice.chatrelation.get

获取指定时间区间内的聊天关系（查询账号，和谁，在哪天说过话）。如：
A 和 B 在2016-09-01 和 2016-09-02 都说过话。以A为查询账号，则该接口将返回：
2016-09-01， B
2016-09-02， B
*/
type TaobaoWangwangEserviceChatrelationGetRequest struct {
    model.Params

    // 请求参数
    chatRelationRequest   *ChatRelationRequest 

}

func NewTaobaoWangwangEserviceChatrelationGetRequest() *TaobaoWangwangEserviceChatrelationGetRequest{
    return &TaobaoWangwangEserviceChatrelationGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWangwangEserviceChatrelationGetRequest) GetApiMethodName() string {
    return "taobao.wangwang.eservice.chatrelation.get"
}

func (r TaobaoWangwangEserviceChatrelationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWangwangEserviceChatrelationGetRequest) SetChatRelationRequest(chatRelationRequest *ChatRelationRequest) error {
    r.chatRelationRequest = chatRelationRequest
    r.Set("chat_relation_request", chatRelationRequest)
    return nil
}

func (r TaobaoWangwangEserviceChatrelationGetRequest) GetChatRelationRequest() *ChatRelationRequest {
    return r.chatRelationRequest
}

