package qianniu

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQianniuTaskIncreaseAPIRequest) Reset() {
	r._tasks = ""
	r._metadataId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskIncreaseAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.task.increase"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskIncreaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQianniuTaskIncreaseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTasks is Tasks Setter
// 任务列表，JSON格式，例如： tasks =[{ &#34;receiver_uid&#34; : 123, &#34;receiver_nick&#34; : &#34;nick&#34;}, { &#34;receiver_uid&#34; : 456, &#34;receiver_nick&#34; : &#34;nick2&#34;} ]
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

var poolTaobaoQianniuTaskIncreaseAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQianniuTaskIncreaseRequest()
	},
}

// GetTaobaoQianniuTaskIncreaseRequest 从 sync.Pool 获取 TaobaoQianniuTaskIncreaseAPIRequest
func GetTaobaoQianniuTaskIncreaseAPIRequest() *TaobaoQianniuTaskIncreaseAPIRequest {
	return poolTaobaoQianniuTaskIncreaseAPIRequest.Get().(*TaobaoQianniuTaskIncreaseAPIRequest)
}

// ReleaseTaobaoQianniuTaskIncreaseAPIRequest 将 TaobaoQianniuTaskIncreaseAPIRequest 放入 sync.Pool
func ReleaseTaobaoQianniuTaskIncreaseAPIRequest(v *TaobaoQianniuTaskIncreaseAPIRequest) {
	v.Reset()
	poolTaobaoQianniuTaskIncreaseAPIRequest.Put(v)
}
