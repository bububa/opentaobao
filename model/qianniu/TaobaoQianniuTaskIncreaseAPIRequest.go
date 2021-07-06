package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskIncreaseAPIRequest 增加任务接收人接口 API请求
// taobao.qianniu.task.increase
//
// 根据任务元id增加任务接收人
type TaobaoQianniuTaskIncreaseAPIRequest struct {
	model.Params
	// 任务列表，JSON格式，例如： tasks =[{ "receiver_uid" : 123, "receiver_nick" : "nick"}, { "receiver_uid" : 456, "receiver_nick" : "nick2"} ]
	_tasks string
	// 任务元id
	_metadataId int64
}

// NewTaobaoQianniuTaskIncreaseRequest 初始化TaobaoQianniuTaskIncreaseAPIRequest对象
func NewTaobaoQianniuTaskIncreaseRequest() *TaobaoQianniuTaskIncreaseAPIRequest {
	return &TaobaoQianniuTaskIncreaseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskIncreaseAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.increase"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskIncreaseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTasks is Tasks Setter
// 任务列表，JSON格式，例如： tasks =[{ "receiver_uid" : 123, "receiver_nick" : "nick"}, { "receiver_uid" : 456, "receiver_nick" : "nick2"} ]
func (r *TaobaoQianniuTaskIncreaseAPIRequest) SetTasks(_tasks string) error {
	r._tasks = _tasks
	r.Set("tasks", _tasks)
	return nil
}

// GetTasks Tasks Getter
func (r TaobaoQianniuTaskIncreaseAPIRequest) GetTasks() string {
	return r._tasks
}

// SetMetadataId is MetadataId Setter
// 任务元id
func (r *TaobaoQianniuTaskIncreaseAPIRequest) SetMetadataId(_metadataId int64) error {
	r._metadataId = _metadataId
	r.Set("metadata_id", _metadataId)
	return nil
}

// GetMetadataId MetadataId Getter
func (r TaobaoQianniuTaskIncreaseAPIRequest) GetMetadataId() int64 {
	return r._metadataId
}
