package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripitfarequerytaskAPIRequest 【国际机票自有政策】批量操作结果查询 API请求
// taobao.alitrip.it.fare.querytask
//
// 批量操作同步返回任务id，后台生成异步任务，通过此接口查询批量操作的执行结果
type TaobaoalitripitfarequerytaskAPIRequest struct {
	model.Params
	// json格式的字符串，扩展属性，预留
	_extendAttributes string
	// 任务id
	_taskId int64
}

// NewTaobaoalitripitfarequerytaskRequest 初始化TaobaoalitripitfarequerytaskAPIRequest对象
func NewTaobaoalitripitfarequerytaskRequest() *TaobaoalitripitfarequerytaskAPIRequest {
	return &TaobaoalitripitfarequerytaskAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripitfarequerytaskAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.fare.querytask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripitfarequerytaskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripitfarequerytaskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtendAttributes is ExtendAttributes Setter
// json格式的字符串，扩展属性，预留
func (r *TaobaoalitripitfarequerytaskAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extendAttributes", _extendAttributes)
	return nil
}

// GetExtendAttributes ExtendAttributes Getter
func (r TaobaoalitripitfarequerytaskAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// SetTaskId is TaskId Setter
// 任务id
func (r *TaobaoalitripitfarequerytaskAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("taskId", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r TaobaoalitripitfarequerytaskAPIRequest) GetTaskId() int64 {
	return r._taskId
}
