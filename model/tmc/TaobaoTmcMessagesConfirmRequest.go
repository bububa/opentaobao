package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认消费消息的状态 API请求
taobao.tmc.messages.confirm

确认消费消息的状态
*/
type TaobaoTmcMessagesConfirmRequest struct {
    model.Params
    // 分组名称，不传代表默认分组
    groupName   string
    // 处理成功的消息ID列表 最大 200个ID
    sMessageIds   []int64
    // 处理失败的消息ID列表--已废弃，无需传此字段
    fMessageIds   []int64
}

// 初始化TaobaoTmcMessagesConfirmRequest对象
func NewTaobaoTmcMessagesConfirmRequest() *TaobaoTmcMessagesConfirmRequest{
    return &TaobaoTmcMessagesConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmcMessagesConfirmRequest) GetApiMethodName() string {
    return "taobao.tmc.messages.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmcMessagesConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupName Setter
// 分组名称，不传代表默认分组
func (r *TaobaoTmcMessagesConfirmRequest) SetGroupName(groupName string) error {
    r.groupName = groupName
    r.Set("group_name", groupName)
    return nil
}

// GroupName Getter
func (r TaobaoTmcMessagesConfirmRequest) GetGroupName() string {
    return r.groupName
}
// SMessageIds Setter
// 处理成功的消息ID列表 最大 200个ID
func (r *TaobaoTmcMessagesConfirmRequest) SetSMessageIds(sMessageIds []int64) error {
    r.sMessageIds = sMessageIds
    r.Set("s_message_ids", sMessageIds)
    return nil
}

// SMessageIds Getter
func (r TaobaoTmcMessagesConfirmRequest) GetSMessageIds() []int64 {
    return r.sMessageIds
}
// FMessageIds Setter
// 处理失败的消息ID列表--已废弃，无需传此字段
func (r *TaobaoTmcMessagesConfirmRequest) SetFMessageIds(fMessageIds []int64) error {
    r.fMessageIds = fMessageIds
    r.Set("f_message_ids", fMessageIds)
    return nil
}

// FMessageIds Getter
func (r TaobaoTmcMessagesConfirmRequest) GetFMessageIds() []int64 {
    return r.fMessageIds
}
