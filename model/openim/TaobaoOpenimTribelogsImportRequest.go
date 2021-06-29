package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openim群聊天记录导入 API请求
taobao.openim.tribelogs.import

openim群聊天记录导入
*/
type TaobaoOpenimTribelogsImportRequest struct {
    model.Params
    // 群号。必须为已存在的群，且群主属于本app
    tribeId   int64
    // 消息列表
    messages   []TribeTextMessage
}

// 初始化TaobaoOpenimTribelogsImportRequest对象
func NewTaobaoOpenimTribelogsImportRequest() *TaobaoOpenimTribelogsImportRequest{
    return &TaobaoOpenimTribelogsImportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribelogsImportRequest) GetApiMethodName() string {
    return "taobao.openim.tribelogs.import"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribelogsImportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TribeId Setter
// 群号。必须为已存在的群，且群主属于本app
func (r *TaobaoOpenimTribelogsImportRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribelogsImportRequest) GetTribeId() int64 {
    return r.tribeId
}
// Messages Setter
// 消息列表
func (r *TaobaoOpenimTribelogsImportRequest) SetMessages(messages []TribeTextMessage) error {
    r.messages = messages
    r.Set("messages", messages)
    return nil
}

// Messages Getter
func (r TaobaoOpenimTribelogsImportRequest) GetMessages() []TribeTextMessage {
    return r.messages
}
