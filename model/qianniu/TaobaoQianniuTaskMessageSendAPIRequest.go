package qianniu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskMessageSendAPIRequest 发送任务提醒消息 API请求
// taobao.qianniu.task.message.send
//
// 如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。消息会以任务消息的形式发给客户端。
type TaobaoQianniuTaskMessageSendAPIRequest struct {
	model.Params
	// 任务ID。如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。
	_taskId int64
	// 任务元id，如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。
	_metadataId int64
}

// NewTaobaoQianniuTaskMessageSendRequest 初始化TaobaoQianniuTaskMessageSendAPIRequest对象
func NewTaobaoQianniuTaskMessageSendRequest() *TaobaoQianniuTaskMessageSendAPIRequest {
	return &TaobaoQianniuTaskMessageSendAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQianniuTaskMessageSendAPIRequest) Reset() {
	r._taskId = 0
	r._metadataId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskMessageSendAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.message.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskMessageSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQianniuTaskMessageSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaskId is TaskId Setter
// 任务ID。如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。
func (r *TaobaoQianniuTaskMessageSendAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r TaobaoQianniuTaskMessageSendAPIRequest) GetTaskId() int64 {
	return r._taskId
}

// SetMetadataId is MetadataId Setter
// 任务元id，如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。
func (r *TaobaoQianniuTaskMessageSendAPIRequest) SetMetadataId(_metadataId int64) error {
	r._metadataId = _metadataId
	r.Set("metadata_id", _metadataId)
	return nil
}

// GetMetadataId MetadataId Getter
func (r TaobaoQianniuTaskMessageSendAPIRequest) GetMetadataId() int64 {
	return r._metadataId
}

var poolTaobaoQianniuTaskMessageSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQianniuTaskMessageSendRequest()
	},
}

// GetTaobaoQianniuTaskMessageSendRequest 从 sync.Pool 获取 TaobaoQianniuTaskMessageSendAPIRequest
func GetTaobaoQianniuTaskMessageSendAPIRequest() *TaobaoQianniuTaskMessageSendAPIRequest {
	return poolTaobaoQianniuTaskMessageSendAPIRequest.Get().(*TaobaoQianniuTaskMessageSendAPIRequest)
}

// ReleaseTaobaoQianniuTaskMessageSendAPIRequest 将 TaobaoQianniuTaskMessageSendAPIRequest 放入 sync.Pool
func ReleaseTaobaoQianniuTaskMessageSendAPIRequest(v *TaobaoQianniuTaskMessageSendAPIRequest) {
	v.Reset()
	poolTaobaoQianniuTaskMessageSendAPIRequest.Put(v)
}
