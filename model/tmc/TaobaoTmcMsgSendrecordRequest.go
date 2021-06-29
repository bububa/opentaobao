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
    _groupName   string
    // TOPIC名称
    _topicName   string
    // 消息主键ID
    _dataId   string
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
func (r *TaobaoTmcMsgSendrecordRequest) SetGroupName(_groupName string) error {
    r._groupName = _groupName
    r.Set("group_name", _groupName)
    return nil
}

// GroupName Getter
func (r TaobaoTmcMsgSendrecordRequest) GetGroupName() string {
    return r._groupName
}
// TopicName Setter
// TOPIC名称
func (r *TaobaoTmcMsgSendrecordRequest) SetTopicName(_topicName string) error {
    r._topicName = _topicName
    r.Set("topic_name", _topicName)
    return nil
}

// TopicName Getter
func (r TaobaoTmcMsgSendrecordRequest) GetTopicName() string {
    return r._topicName
}
// DataId Setter
// 消息主键ID
func (r *TaobaoTmcMsgSendrecordRequest) SetDataId(_dataId string) error {
    r._dataId = _dataId
    r.Set("data_id", _dataId)
    return nil
}

// DataId Getter
func (r TaobaoTmcMsgSendrecordRequest) GetDataId() string {
    return r._dataId
}
