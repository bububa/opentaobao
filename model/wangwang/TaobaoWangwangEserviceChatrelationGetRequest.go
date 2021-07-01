package wangwang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聊天关系获取接口 API请求
taobao.wangwang.eservice.chatrelation.get

获取指定时间区间内的聊天关系（查询账号，和谁，在哪天说过话）。如：
A 和 B 在2016-09-01 和 2016-09-02 都说过话。以A为查询账号，则该接口将返回：
2016-09-01， B
2016-09-02， B
*/
type TaobaoWangwangEserviceChatrelationGetAPIRequest struct {
    model.Params
    // 请求参数
    _chatRelationRequest   *ChatRelationRequest
}

// 初始化TaobaoWangwangEserviceChatrelationGetAPIRequest对象
func NewTaobaoWangwangEserviceChatrelationGetRequest() *TaobaoWangwangEserviceChatrelationGetAPIRequest{
    return &TaobaoWangwangEserviceChatrelationGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWangwangEserviceChatrelationGetAPIRequest) GetApiMethodName() string {
    return "taobao.wangwang.eservice.chatrelation.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWangwangEserviceChatrelationGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChatRelationRequest Setter
// 请求参数
func (r *TaobaoWangwangEserviceChatrelationGetAPIRequest) SetChatRelationRequest(_chatRelationRequest *ChatRelationRequest) error {
    r._chatRelationRequest = _chatRelationRequest
    r.Set("chat_relation_request", _chatRelationRequest)
    return nil
}

// ChatRelationRequest Getter
func (r TaobaoWangwangEserviceChatrelationGetAPIRequest) GetChatRelationRequest() *ChatRelationRequest {
    return r._chatRelationRequest
}
