package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询消息列表 APIRequest
alibaba.alink.message.history.list

查询消息列表
*/
type AlibabaAlinkMessageHistoryListRequest struct {
    model.Params

    // 设备id
    uuid   string 

    // 消息类型 1:通知, 2:报警, 3:运营，5:语音控制机器人响应，6:语音控
    type   string 

    // 消息状态，0：未读；1：已读
    status   string 

    // 消息级别 1：普通；2：重要消息
    level   string 

    // 查询多少条数据
    limit   string 

    // 偏移量
    offset   string 

}

func NewAlibabaAlinkMessageHistoryListRequest() *AlibabaAlinkMessageHistoryListRequest{
    return &AlibabaAlinkMessageHistoryListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlinkMessageHistoryListRequest) GetApiMethodName() string {
    return "alibaba.alink.message.history.list"
}

func (r AlibabaAlinkMessageHistoryListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlinkMessageHistoryListRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaAlinkMessageHistoryListRequest) GetUuid() string {
    return r.uuid
}

func (r *AlibabaAlinkMessageHistoryListRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaAlinkMessageHistoryListRequest) GetType() string {
    return r.type
}

func (r *AlibabaAlinkMessageHistoryListRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaAlinkMessageHistoryListRequest) GetStatus() string {
    return r.status
}

func (r *AlibabaAlinkMessageHistoryListRequest) SetLevel(level string) error {
    r.level = level
    r.Set("level", level)
    return nil
}

func (r AlibabaAlinkMessageHistoryListRequest) GetLevel() string {
    return r.level
}

func (r *AlibabaAlinkMessageHistoryListRequest) SetLimit(limit string) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

func (r AlibabaAlinkMessageHistoryListRequest) GetLimit() string {
    return r.limit
}

func (r *AlibabaAlinkMessageHistoryListRequest) SetOffset(offset string) error {
    r.offset = offset
    r.Set("offset", offset)
    return nil
}

func (r AlibabaAlinkMessageHistoryListRequest) GetOffset() string {
    return r.offset
}

