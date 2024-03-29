package itpolicy

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItFareQuerytaskAPIRequest 【国际机票自有政策】批量操作结果查询 API请求
// taobao.alitrip.it.fare.querytask
//
// 批量操作同步返回任务id，后台生成异步任务，通过此接口查询批量操作的执行结果
type TaobaoAlitripItFareQuerytaskAPIRequest struct {
	model.Params
	// json格式的字符串，扩展属性，预留
	_extendAttributes string
	// 任务id
	_taskId int64
}

// NewTaobaoAlitripItFareQuerytaskRequest 初始化TaobaoAlitripItFareQuerytaskAPIRequest对象
func NewTaobaoAlitripItFareQuerytaskRequest() *TaobaoAlitripItFareQuerytaskAPIRequest {
	return &TaobaoAlitripItFareQuerytaskAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripItFareQuerytaskAPIRequest) Reset() {
	r._extendAttributes = ""
	r._taskId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItFareQuerytaskAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.fare.querytask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItFareQuerytaskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripItFareQuerytaskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtendAttributes is ExtendAttributes Setter
// json格式的字符串，扩展属性，预留
func (r *TaobaoAlitripItFareQuerytaskAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extendAttributes", _extendAttributes)
	return nil
}

// GetExtendAttributes ExtendAttributes Getter
func (r TaobaoAlitripItFareQuerytaskAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// SetTaskId is TaskId Setter
// 任务id
func (r *TaobaoAlitripItFareQuerytaskAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("taskId", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r TaobaoAlitripItFareQuerytaskAPIRequest) GetTaskId() int64 {
	return r._taskId
}

var poolTaobaoAlitripItFareQuerytaskAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripItFareQuerytaskRequest()
	},
}

// GetTaobaoAlitripItFareQuerytaskRequest 从 sync.Pool 获取 TaobaoAlitripItFareQuerytaskAPIRequest
func GetTaobaoAlitripItFareQuerytaskAPIRequest() *TaobaoAlitripItFareQuerytaskAPIRequest {
	return poolTaobaoAlitripItFareQuerytaskAPIRequest.Get().(*TaobaoAlitripItFareQuerytaskAPIRequest)
}

// ReleaseTaobaoAlitripItFareQuerytaskAPIRequest 将 TaobaoAlitripItFareQuerytaskAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripItFareQuerytaskAPIRequest(v *TaobaoAlitripItFareQuerytaskAPIRequest) {
	v.Reset()
	poolTaobaoAlitripItFareQuerytaskAPIRequest.Put(v)
}
