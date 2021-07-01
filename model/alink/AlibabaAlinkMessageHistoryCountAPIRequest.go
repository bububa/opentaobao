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

// New
