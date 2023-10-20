package idleparttime

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleparttimenotifyAPIRequest 兼职通知接口 API请求
// alibaba.idle.parttime.notify
//
// 服务商侧有岗位状态变更对我们进行通知(岗位关闭, 录取状态)
type AlibabaidleparttimenotifyAPIRequest struct {
	model.Params
	// 通知消息
	_message string
	// 实时同步类型, 0: 岗位状态, 1: 录取状态
	_type int64
	// 岗位: 0关闭 ; 录取: 0不录取, 1录取
	_status int64
	// 岗位id
	_jobId int64
	// 用户id
	_userId int64
	// 同步时间
	_syncTime int64
	// 报名id
	_applyId int64
}

// NewAlibabaidleparttimenotifyRequest 初始化AlibabaidleparttimenotifyAPIRequest对象
func NewAlibabaidleparttimenotifyRequest() *AlibabaidleparttimenotifyAPIRequest {
	return &AlibabaidleparttimenotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleparttimenotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.parttime.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleparttimenotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleparttimenotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMessage is Message Setter
// 通知消息
func (r *AlibabaidleparttimenotifyAPIRequest) SetMessage(_message string) error {
	r._message = _message
	r.Set("message", _message)
	return nil
}

// GetMessage Message Getter
func (r AlibabaidleparttimenotifyAPIRequest) GetMessage() string {
	return r._message
}

// SetType is Type Setter
// 实时同步类型, 0: 岗位状态, 1: 录取状态
func (r *AlibabaidleparttimenotifyAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaidleparttimenotifyAPIRequest) GetType() int64 {
	return r._type
}

// SetStatus is Status Setter
// 岗位: 0关闭 ; 录取: 0不录取, 1录取
func (r *AlibabaidleparttimenotifyAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaidleparttimenotifyAPIRequest) GetStatus() int64 {
	return r._status
}

// SetJobId is JobId Setter
// 岗位id
func (r *AlibabaidleparttimenotifyAPIRequest) SetJobId(_jobId int64) error {
	r._jobId = _jobId
	r.Set("job_id", _jobId)
	return nil
}

// GetJobId JobId Getter
func (r AlibabaidleparttimenotifyAPIRequest) GetJobId() int64 {
	return r._jobId
}

// SetUserId is UserId Setter
// 用户id
func (r *AlibabaidleparttimenotifyAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaidleparttimenotifyAPIRequest) GetUserId() int64 {
	return r._userId
}

// SetSyncTime is SyncTime Setter
// 同步时间
func (r *AlibabaidleparttimenotifyAPIRequest) SetSyncTime(_syncTime int64) error {
	r._syncTime = _syncTime
	r.Set("sync_time", _syncTime)
	return nil
}

// GetSyncTime SyncTime Getter
func (r AlibabaidleparttimenotifyAPIRequest) GetSyncTime() int64 {
	return r._syncTime
}

// SetApplyId is ApplyId Setter
// 报名id
func (r *AlibabaidleparttimenotifyAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlibabaidleparttimenotifyAPIRequest) GetApplyId() int64 {
	return r._applyId
}
