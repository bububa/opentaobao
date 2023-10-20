package jstinteractive

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveActivityCreateAPIRequest 互动任务活动创建接口 API请求
// taobao.jst.interactive.activity.create
//
// 调用活动创建接口为小程序创建互动任务活动，任务列表仅在活动期间内返回
type TaobaoJstInteractiveActivityCreateAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
	// 活动名称
	_activityName string
	// 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
	_startTime string
	// 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
	_endTime string
}

// NewTaobaoJstInteractiveActivityCreateRequest 初始化TaobaoJstInteractiveActivityCreateAPIRequest对象
func NewTaobaoJstInteractiveActivityCreateRequest() *TaobaoJstInteractiveActivityCreateAPIRequest {
	return &TaobaoJstInteractiveActivityCreateAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstInteractiveActivityCreateAPIRequest) Reset() {
	r._miniAppId = ""
	r._activityName = ""
	r._startTime = ""
	r._endTime = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveActivityCreateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveActivityCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstInteractiveActivityCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveActivityCreateAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaoJstInteractiveActivityCreateAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}

// SetActivityName is ActivityName Setter
// 活动名称
func (r *TaobaoJstInteractiveActivityCreateAPIRequest) SetActivityName(_activityName string) error {
	r._activityName = _activityName
	r.Set("activity_name", _activityName)
	return nil
}

// GetActivityName ActivityName Getter
func (r TaobaoJstInteractiveActivityCreateAPIRequest) GetActivityName() string {
	return r._activityName
}

// SetStartTime is StartTime Setter
// 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
func (r *TaobaoJstInteractiveActivityCreateAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoJstInteractiveActivityCreateAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
func (r *TaobaoJstInteractiveActivityCreateAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoJstInteractiveActivityCreateAPIRequest) GetEndTime() string {
	return r._endTime
}

var poolTaobaoJstInteractiveActivityCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstInteractiveActivityCreateRequest()
	},
}

// GetTaobaoJstInteractiveActivityCreateRequest 从 sync.Pool 获取 TaobaoJstInteractiveActivityCreateAPIRequest
func GetTaobaoJstInteractiveActivityCreateAPIRequest() *TaobaoJstInteractiveActivityCreateAPIRequest {
	return poolTaobaoJstInteractiveActivityCreateAPIRequest.Get().(*TaobaoJstInteractiveActivityCreateAPIRequest)
}

// ReleaseTaobaoJstInteractiveActivityCreateAPIRequest 将 TaobaoJstInteractiveActivityCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstInteractiveActivityCreateAPIRequest(v *TaobaoJstInteractiveActivityCreateAPIRequest) {
	v.Reset()
	poolTaobaoJstInteractiveActivityCreateAPIRequest.Put(v)
}
