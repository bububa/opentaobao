package jstinteractive

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveActivityUpdateAPIRequest 互动任务活动修改接口 API请求
// taobao.jst.interactive.activity.update
//
// 互动任务活动修改接口
type TaobaoJstInteractiveActivityUpdateAPIRequest struct {
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

// NewTaobaoJstInteractiveActivityUpdateRequest 初始化TaobaoJstInteractiveActivityUpdateAPIRequest对象
func NewTaobaoJstInteractiveActivityUpdateRequest() *TaobaoJstInteractiveActivityUpdateAPIRequest {
	return &TaobaoJstInteractiveActivityUpdateAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstInteractiveActivityUpdateAPIRequest) Reset() {
	r._miniAppId = ""
	r._startTime = ""
	r._endTime = ""
	r._status = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveActivityUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveActivityUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstInteractiveActivityUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniAppId is MiniAppId Setter
// 小程序id
func (r *TaobaoJstInteractiveActivityUpdateAPIRequest) SetMiniAppId(_miniAppId string) error {
	r._miniAppId = _miniAppId
	r.Set("mini_app_id", _miniAppId)
	return nil
}

// GetMiniAppId MiniAppId Getter
func (r TaobaoJstInteractiveActivityUpdateAPIRequest) GetMiniAppId() string {
	return r._miniAppId
}

// SetStartTime is StartTime Setter
// 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
func (r *TaobaoJstInteractiveActivityUpdateAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoJstInteractiveActivityUpdateAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
func (r *TaobaoJstInteractiveActivityUpdateAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoJstInteractiveActivityUpdateAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStatus is Status Setter
// 活动状态，0=无效，1=有效，status设为0即代表删除此活动，需创建新的活动
func (r *TaobaoJstInteractiveActivityUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoJstInteractiveActivityUpdateAPIRequest) GetStatus() int64 {
	return r._status
}

var poolTaobaoJstInteractiveActivityUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstInteractiveActivityUpdateRequest()
	},
}

// GetTaobaoJstInteractiveActivityUpdateRequest 从 sync.Pool 获取 TaobaoJstInteractiveActivityUpdateAPIRequest
func GetTaobaoJstInteractiveActivityUpdateAPIRequest() *TaobaoJstInteractiveActivityUpdateAPIRequest {
	return poolTaobaoJstInteractiveActivityUpdateAPIRequest.Get().(*TaobaoJstInteractiveActivityUpdateAPIRequest)
}

// ReleaseTaobaoJstInteractiveActivityUpdateAPIRequest 将 TaobaoJstInteractiveActivityUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstInteractiveActivityUpdateAPIRequest(v *TaobaoJstInteractiveActivityUpdateAPIRequest) {
	v.Reset()
	poolTaobaoJstInteractiveActivityUpdateAPIRequest.Put(v)
}
