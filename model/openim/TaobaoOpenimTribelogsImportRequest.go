package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openim群聊天记录导入 APIRequest
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

func NewTaobaoOpenimTribelogsImportRequest() *TaobaoOpenimTribelogsImportRequest{
    return &TaobaoOpenimTribelogsImportRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimTribelogsImportRequest) GetApiMethodName() string {
    return "taobao.openim.tribelogs.import"
}

func (r TaobaoOpenimTribelogsImportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimTribelogsImportRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

func (r TaobaoOpenimTribelogsImportRequest) GetTribeId() int64 {
    return r.tribeId
}

func (r *TaobaoOpenimTribelogsImportRequest) SetMessages(messages []TribeTextMessage) error {
    r.messages = messages
    r.Set("messages", messages)
    return nil
}

func (r TaobaoOpenimTribelogsImportRequest) GetMessages() []TribeTextMessage {
    return r.messages
}

