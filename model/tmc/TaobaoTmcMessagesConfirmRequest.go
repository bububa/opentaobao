package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认消费消息的状态 APIRequest
taobao.tmc.messages.confirm

确认消费消息的状态
*/
type TaobaoTmcMessagesConfirmRequest struct {
    model.Params

    // 分组名称，不传代表默认分组
    groupName   string 

    // 处理成功的消息ID列表 最大 200个ID
    sMessageIds   []Number 

    // 处理失败的消息ID列表--已废弃，无需传此字段
    fMessageIds   []Number 

}

func NewTaobaoTmcMessagesConfirmRequest() *TaobaoTmcMessagesConfirmRequest{
    return &TaobaoTmcMessagesConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTmcMessagesConfirmRequest) GetApiMethodName() string {
    return "taobao.tmc.messages.confirm"
}

func (r TaobaoTmcMessagesConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTmcMessagesConfirmRequest) SetGroupName(groupName string) error {
    r.groupName = groupName
    r.Set("group_name", groupName)
    return nil
}

func (r TaobaoTmcMessagesConfirmRequest) GetGroupName() string {
    return r.groupName
}

func (r *TaobaoTmcMessagesConfirmRequest) SetSMessageIds(sMessageIds []Number) error {
    r.sMessageIds = sMessageIds
    r.Set("s_message_ids", sMessageIds)
    return nil
}

func (r TaobaoTmcMessagesConfirmRequest) GetSMessageIds() []Number {
    return r.sMessageIds
}

func (r *TaobaoTmcMessagesConfirmRequest) SetFMessageIds(fMessageIds []Number) error {
    r.fMessageIds = fMessageIds
    r.Set("f_message_ids", fMessageIds)
    return nil
}

func (r TaobaoTmcMessagesConfirmRequest) GetFMessageIds() []Number {
    return r.fMessageIds
}

