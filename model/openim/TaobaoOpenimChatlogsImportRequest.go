package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openim单聊消息导入 APIRequest
taobao.openim.chatlogs.import

提供openim账号的聊天消息导入功能
*/
type TaobaoOpenimChatlogsImportRequest struct {
    model.Params

    // 消息序列
    messages   []TextMessage 

}

func NewTaobaoOpenimChatlogsImportRequest() *TaobaoOpenimChatlogsImportRequest{
    return &TaobaoOpenimChatlogsImportRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimChatlogsImportRequest) GetApiMethodName() string {
    return "taobao.openim.chatlogs.import"
}

func (r TaobaoOpenimChatlogsImportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimChatlogsImportRequest) SetMessages(messages []TextMessage) error {
    r.messages = messages
    r.Set("messages", messages)
    return nil
}

func (r TaobaoOpenimChatlogsImportRequest) GetMessages() []TextMessage {
    return r.messages
}

