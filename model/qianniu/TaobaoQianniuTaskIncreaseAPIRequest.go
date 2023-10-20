package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqianniutaskincreaseAPIRequest 增加任务接收人接口 API请求
// taobao.qianniu.task.increase
//
// 根据任务元id增加任务接收人
type TaobaoqianniutaskincreaseAPIRequest struct {
	model.Params
	// 任务列表，JSON格式，例如： tasks =[{ "receiver_uid" : 123, "receiver_nick" : "nick"}, { "receiver_uid" : 456, "receiver_nick" : "nick2"} ]
	_tasks string
	// 任务元id
	_metadataId int64
}

// NewTaobaoqianniutaskincreaseRequest 初始化TaobaoqianniutaskincreaseAPIRequest对象
func NewTaobaoqianniutaskincreaseRequest() *TaobaoqianniutaskincreaseAPIRequest {
	return &TaobaoqianniutaskincreaseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqianniutaskincreaseAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.increase"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqianniutaskincreaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqianniutaskincreaseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTasks is Tasks Setter
// 任务列表，JSON格式，例如： tasks =[{ &#34;receiver_uid&#34; : 123, &#34;receiver_nick&#34; : &#34;nick&#34;}, { &#34;receiver_uid&#34; : 456, &#34;receiver_nick&#34; : &#34;nick2&#34;} ]
func (r *TaobaoqianniutaskincreaseAPIRequest) SetTasks(_tasks string) error {
	r._tasks = _tasks
	r.Set("tasks", _tasks)
	return nil
}

// GetTasks Tasks Getter
func (r TaobaoqianniutaskincreaseAPIRequest) GetTasks() string {
	return r._tasks
}

// SetMetadataId is MetadataId Setter
// 任务元id
func (r *TaobaoqianniutaskincreaseAPIRequest) SetMetadataId(_metadataId int64) error {
	r._metadataId = _metadataId
	r.Set("metadata_id", _metadataId)
	return nil
}

// GetMetadataId MetadataId Getter
func (r TaobaoqianniutaskincreaseAPIRequest) GetMetadataId() int64 {
	return r._metadataId
}
