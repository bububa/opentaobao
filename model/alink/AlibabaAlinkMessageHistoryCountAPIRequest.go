package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalinkmessagehistorycountAPIRequest 查询消息总数 API请求
// alibaba.alink.message.history.count
//
// 查询消息总数
type AlibabaalinkmessagehistorycountAPIRequest struct {
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

// NewAlibabaalinkmessagehistorycountRequest 初始化AlibabaalinkmessagehistorycountAPIRequest对象
func NewAlibabaalinkmessagehistorycountRequest() *AlibabaalinkmessagehistorycountAPIRequest {
	return &AlibabaalinkmessagehistorycountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalinkmessagehistorycountAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.message.history.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalinkmessagehistorycountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalinkmessagehistorycountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *AlibabaalinkmessagehistorycountAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaalinkmessagehistorycountAPIRequest) GetUuid() string {
	return r._uuid
}

// SetType is Type Setter
// 消息类型 1:通知, 2:报警, 3:运营，5:语音控制机器人响应，6:语音控
func (r *AlibabaalinkmessagehistorycountAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaalinkmessagehistorycountAPIRequest) GetType() string {
	return r._type
}

// SetStatus is Status Setter
// 消息状态，0：未读；1：已读
func (r *AlibabaalinkmessagehistorycountAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaalinkmessagehistorycountAPIRequest) GetStatus() string {
	return r._status
}

// SetLevel is Level Setter
// 消息级别 1：普通；2：重要消息
func (r *AlibabaalinkmessagehistorycountAPIRequest) SetLevel(_level string) error {
	r._level = _level
	r.Set("level", _level)
	return nil
}

// GetLevel Level Getter
func (r AlibabaalinkmessagehistorycountAPIRequest) GetLevel() string {
	return r._level
}

// SetLimit is Limit Setter
// 查询多少条数据
func (r *AlibabaalinkmessagehistorycountAPIRequest) SetLimit(_limit string) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// GetLimit Limit Getter
func (r AlibabaalinkmessagehistorycountAPIRequest) GetLimit() string {
	return r._limit
}

// SetOffset is Offset Setter
// 偏移量
func (r *AlibabaalinkmessagehistorycountAPIRequest) SetOffset(_offset string) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r AlibabaalinkmessagehistorycountAPIRequest) GetOffset() string {
	return r._offset
}
