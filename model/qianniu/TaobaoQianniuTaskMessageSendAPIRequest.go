package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqianniutaskmessagesendAPIRequest 发送任务提醒消息 API请求
// taobao.qianniu.task.message.send
//
// 如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。消息会以任务消息的形式发给客户端。
type TaobaoqianniutaskmessagesendAPIRequest struct {
	model.Params
	// 任务ID。如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。
	_taskId int64
	// 任务元id，如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。
	_metadataId int64
}

// NewTaobaoqianniutaskmessagesendRequest 初始化TaobaoqianniutaskmessagesendAPIRequest对象
func NewTaobaoqianniutaskmessagesendRequest() *TaobaoqianniutaskmessagesendAPIRequest {
	return &TaobaoqianniutaskmessagesendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqianniutaskmessagesendAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.message.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqianniutaskmessagesendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqianniutaskmessagesendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaskId is TaskId Setter
// 任务ID。如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。
func (r *TaobaoqianniutaskmessagesendAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r TaobaoqianniutaskmessagesendAPIRequest) GetTaskId() int64 {
	return r._taskId
}

// SetMetadataId is MetadataId Setter
// 任务元id，如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。
func (r *TaobaoqianniutaskmessagesendAPIRequest) SetMetadataId(_metadataId int64) error {
	r._metadataId = _metadataId
	r.Set("metadata_id", _metadataId)
	return nil
}

// GetMetadataId MetadataId Getter
func (r TaobaoqianniutaskmessagesendAPIRequest) GetMetadataId() int64 {
	return r._metadataId
}
