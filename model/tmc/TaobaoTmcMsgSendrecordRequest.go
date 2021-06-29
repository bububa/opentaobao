package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息发送记录查询 API请求
taobao.tmc.msg.sendrecord

查询单条消息发送记录，只返回返回条数和时间。
*/
type TaobaoTmcMsgSendrecordRequest struct {
    model.Params
    // 消息分组名
    groupName   string
    // TOPIC名称
    topicName   string
    // 消息主键ID
    dataId   string
}

// 初始化TaobaoTmcMsgSendrecordRequest对象
func NewTaobaoTmcMsgSendrecordRequest() *TaobaoTmcMsgSendrecordRequest{
    return &TaobaoTmcMsgSendrecordRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmcMsgSendrecordRequest) GetApiMethodName() string {
    return "taobao.tmc.msg.sendrecord"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmcMsgSendrecordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupName Setter
// 消息分组名
func (r *TaobaoTmcMsgSendrecordRequest) SetGroupName(groupName string) error {
    r.groupName = groupName
    r.Set("group_name", groupName)
    return nil
}

// GroupName Getter
func (r TaobaoTmcMsgSendrecordRequest) GetGroupName() string {
    return r.groupName
}
// TopicName Setter
// TOPIC名称
func (r *TaobaoTmcMsgSendrecordRequest) SetTopicName(topicName string) error {
    r.topicName = topicName
    r.Set("topic_name", topicName)
    return nil
}

// TopicName Getter
func (r TaobaoTmcMsgSendrecordRequest) GetTopicName() string {
    return r.topicName
}
// DataId Setter
// 消息主键ID
func (r *TaobaoTmcMsgSendrecordRequest) SetDataId(dataId string) error {
    r.dataId = dataId
    r.Set("data_id", dataId)
    return nil
}

// DataId Getter
func (r TaobaoTmcMsgSendrecordRequest) GetDataId() string {
    return r.dataId
}
