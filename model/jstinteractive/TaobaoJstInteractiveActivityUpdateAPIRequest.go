package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstinteractiveactivityupdateAPIRequest 互动任务活动修改接口 API请求
// taobao.jst.interactive.activity.update
//
// 互动任务活动修改接口
type TaobaojstinteractiveactivityupdateAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
	// 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
	_startTime string
	// 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
	_endTime string
	// 活动状态，0=无效，1=有效，status设为0即代表删除此活动，需创建新的活动
	_status int64
}

// NewTaobaojstinteractiveactivityupdateRequest 初始化TaobaojstinteractiveactivityupdateAPIRequest对象
func NewTaobaojstinteractiveactivityupdateRequest() *TaobaojstinteractiveactivityupdateAPIRequest {
	return &TaobaojstinteractiveactivityupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstinteractiveactivityupdateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstinteractiveactivityupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstinteractiveactivityupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaojstinteractiveactivityupdateAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaojstinteractiveactivityupdateAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}

// SetStartTime is StartTime Setter
// 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
func (r *TaobaojstinteractiveactivityupdateAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaojstinteractiveactivityupdateAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
func (r *TaobaojstinteractiveactivityupdateAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaojstinteractiveactivityupdateAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStatus is Status Setter
// 活动状态，0=无效，1=有效，status设为0即代表删除此活动，需创建新的活动
func (r *TaobaojstinteractiveactivityupdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaojstinteractiveactivityupdateAPIRequest) GetStatus() int64 {
	return r._status
}
