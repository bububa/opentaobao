package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlinkMessageHistoryCountAPIRequest
查询消息总数 API请求
alibaba.alink.message.history.count

查询消息总数 */
type AlibabaAlinkMessageHistoryCountAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
	// 消息类型 1:通知, 2:报警, 3:运营，5:语音控制机器人响应，6:语音控
	_type string
	// 消息状态，0：未读；1：已读
	_status string
	// 消息级别 1：普通；2：重要消息
	_level string
	// 查询多少条数据
	_limit string
	// 偏移量
	_offset string
}

// NewAlibabaAlinkMessageHistoryCountRequest 初始化AlibabaAlinkMessageHistoryCountAPIRequest对象
func NewAlibabaAlinkMessageHistoryCountRequest() *AlibabaAlinkMessageHistoryCountAPIRequest {
	return &AlibabaAlinkMessageHistoryCountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlinkMessageHistoryCountAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.message.history.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlinkMessageHistoryCountAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Uuid Setter
// 设备id
func (r *AlibabaAlinkMessageHistoryCountAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r AlibabaAlinkMessageHistoryCountAPIRequest) GetUuid() string {
	return r._uuid
}

// Set is Type Setter
// 消息类型 1:通知, 2:报警, 3:运营，5:语音控制机器人响应，6:语音控
func (r *AlibabaAlinkMessageHistoryCountAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r AlibabaAlinkMessageHistoryCountAPIRequest) GetType() string {
	return r._type
}

// Set is Status Setter
// 消息状态，0：未读；1：已读
func (r *AlibabaAlinkMessageHistoryCountAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r AlibabaAlinkMessageHistoryCountAPIRequest) GetStatus() string {
	return r._status
}

// Set is Level Setter
// 消息级别 1：普通；2：重要消息
func (r *AlibabaAlinkMessageHistoryCountAPIRequest) SetLevel(_level string) error {
	r._level = _level
	r.Set("level", _level)
	return nil
}

// Get Level Getter
func (r AlibabaAlinkMessageHistoryCountAPIRequest) GetLevel() string {
	return r._level
}

// Set is Limit Setter
// 查询多少条数据
func (r *AlibabaAlinkMessageHistoryCountAPIRequest) SetLimit(_limit string) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// Get Limit Getter
func (r AlibabaAlinkMessageHistoryCountAPIRequest) GetLimit() string {
	return r._limit
}

// Set is Offset Setter
// 偏移量
func (r *AlibabaAlinkMessageHistoryCountAPIRequest) SetOffset(_offset string) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// Get Offset Getter
func (r AlibabaAlinkMessageHistoryCountAPIRequest) GetOffset() string {
	return r._offset
}
