package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstinteractiveactivitycreateAPIRequest 互动任务活动创建接口 API请求
// taobao.jst.interactive.activity.create
//
// 调用活动创建接口为小程序创建互动任务活动，任务列表仅在活动期间内返回
type TaobaojstinteractiveactivitycreateAPIRequest struct {
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

// NewTaobaojstinteractiveactivitycreateRequest 初始化TaobaojstinteractiveactivitycreateAPIRequest对象
func NewTaobaojstinteractiveactivitycreateRequest() *TaobaojstinteractiveactivitycreateAPIRequest {
	return &TaobaojstinteractiveactivitycreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstinteractiveactivitycreateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstinteractiveactivitycreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstinteractiveactivitycreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaojstinteractiveactivitycreateAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaojstinteractiveactivitycreateAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}

// SetActivityName is ActivityName Setter
// 活动名称
func (r *TaobaojstinteractiveactivitycreateAPIRequest) SetActivityName(_activityName string) error {
	r._activityName = _activityName
	r.Set("activity_name", _activityName)
	return nil
}

// GetActivityName ActivityName Getter
func (r TaobaojstinteractiveactivitycreateAPIRequest) GetActivityName() string {
	return r._activityName
}

// SetStartTime is StartTime Setter
// 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
func (r *TaobaojstinteractiveactivitycreateAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaojstinteractiveactivitycreateAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
func (r *TaobaojstinteractiveactivitycreateAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaojstinteractiveactivitycreateAPIRequest) GetEndTime() string {
	return r._endTime
}
