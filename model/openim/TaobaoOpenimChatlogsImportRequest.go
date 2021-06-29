package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openim单聊消息导入 API请求
taobao.openim.chatlogs.import

提供openim账号的聊天消息导入功能
*/
type TaobaoOpenimChatlogsImportRequest struct {
    model.Params
    // 消息序列
    _messages   []TextMessage
}

// 初始化TaobaoOpenimChatlogsImportRequest对象
func NewTaobaoOpenimChatlogsImportRequest() *TaobaoOpenimChatlogsImportRequest{
    return &TaobaoOpenimChatlogsImportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimChatlogsImportRequest) GetApiMethodName() string {
    return "taobao.openim.chatlogs.import"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimChatlogsImportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Messages Setter
// 消息序列
func (r *TaobaoOpenimChatlogsImportRequest) SetMessages(_messages []TextMessage) error {
    r._messages = _messages
    r.Set("messages", _messages)
    return nil
}

// Messages Getter
func (r TaobaoOpenimChatlogsImportRequest) GetMessages() []TextMessage {
    return r._messages
}
