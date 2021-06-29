package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询消息总数 API请求
alibaba.alink.message.history.count

查询消息总数
*/
type AlibabaAlinkMessageHistoryCountRequest struct {
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

// 初始化AlibabaAlinkMessageHistoryCountRequest对象
func NewAlibabaAlinkMessageHistoryCountRequest() *AlibabaAlinkMessageHistoryCountRequest{
    return &AlibabaAlinkMessageHistoryCountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlinkMessageHistoryCountRequest) GetApiMethodName() string {
    return "alibaba.alink.message.history.count"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlinkMessageHistoryCountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uuid Setter
// 设备id
func (r *AlibabaAlinkMessageHistoryCountRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAlinkMessageHistoryCountRequest) GetUuid() string {
    return r.uuid
}
// Type Setter
// 消息类型 1:通知, 2:报警, 3:运营，5:语音控制机器人响应，6:语音控
func (r *AlibabaAlinkMessageHistoryCountRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaAlinkMessageHistoryCountRequest) GetType() string {
    return r.type
}
// Status Setter
// 消息状态，0：未读；1：已读
func (r *AlibabaAlinkMessageHistoryCountRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaAlinkMessageHistoryCountRequest) GetStatus() string {
    return r.status
}
// Level Setter
// 消息级别 1：普通；2：重要消息
func (r *AlibabaAlinkMessageHistoryCountRequest) SetLevel(level string) error {
    r.level = level
    r.Set("level", level)
    return nil
}

// Level Getter
func (r AlibabaAlinkMessageHistoryCountRequest) GetLevel() string {
    return r.level
}
// Limit Setter
// 查询多少条数据
func (r *AlibabaAlinkMessageHistoryCountRequest) SetLimit(limit string) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

// Limit Getter
func (r AlibabaAlinkMessageHistoryCountRequest) GetLimit() string {
    return r.limit
}
// Offset Setter
// 偏移量
func (r *AlibabaAlinkMessageHistoryCountRequest) SetOffset(offset string) error {
    r.offset = offset
    r.Set("offset", offset)
    return nil
}

// Offset Getter
func (r AlibabaAlinkMessageHistoryCountRequest) GetOffset() string {
    return r.offset
}
