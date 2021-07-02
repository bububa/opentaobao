package itpolicy

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItFareQuerytaskAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.fare.querytask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItFareQuerytaskAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ExtendAttributes Setter
// json格式的字符串，扩展属性，预留
func (r *TaobaoAlitripItFareQuerytaskAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extendAttributes", _extendAttributes)
	return nil
}

// Get ExtendAttributes Getter
func (r TaobaoAlitripItFareQuerytaskAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// Set is TaskId Setter
// 任务id
func (r *TaobaoAlitripItFareQuerytaskAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("taskId", _taskId)
	return nil
}

// Get TaskId Getter
func (r TaobaoAlitripItFareQuerytaskAPIRequest) GetTaskId() int64 {
	return r._taskId
}
