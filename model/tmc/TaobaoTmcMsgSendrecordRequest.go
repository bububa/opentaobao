package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息发送记录查询 APIRequest
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

func NewTaobaoTmcMsgSendrecordRequest() *TaobaoTmcMsgSendrecordRequest{
    return &TaobaoTmcMsgSendrecordRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTmcMsgSendrecordRequest) GetApiMethodName() string {
    return "taobao.tmc.msg.sendrecord"
}

func (r TaobaoTmcMsgSendrecordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTmcMsgSendrecordRequest) SetGroupName(groupName string) error {
    r.groupName = groupName
    r.Set("group_name", groupName)
    return nil
}

func (r TaobaoTmcMsgSendrecordRequest) GetGroupName() string {
    return r.groupName
}

func (r *TaobaoTmcMsgSendrecordRequest) SetTopicName(topicName string) error {
    r.topicName = topicName
    r.Set("topic_name", topicName)
    return nil
}

func (r TaobaoTmcMsgSendrecordRequest) GetTopicName() string {
    return r.topicName
}

func (r *TaobaoTmcMsgSendrecordRequest) SetDataId(dataId string) error {
    r.dataId = dataId
    r.Set("data_id", dataId)
    return nil
}

func (r TaobaoTmcMsgSendrecordRequest) GetDataId() string {
    return r.dataId
}

