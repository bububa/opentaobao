package idleparttime

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleParttimeSynclogAPIRequest 兼职同步日志 API请求
// alibaba.idle.parttime.synclog
//
// 提供给供应商查询的接口
type AlibabaIdleParttimeSynclogAPIRequest struct {
	model.Params
	// 同步的id
	_syncIds []string
	// 查询岗位同步开始时间
	_startTime int64
	// 查询岗位同步结束时间
	_endTime int64
	// 查询的类型, 0:岗位
	_type int64
	// 页大小
	_pageSize int64
	// 第几页, 从0开始
	_pageNum int64
}

// NewAlibabaIdleParttimeSynclogRequest 初始化AlibabaIdleParttimeSynclogAPIRequest对象
func NewAlibabaIdleParttimeSynclogRequest() *AlibabaIdleParttimeSynclogAPIRequest {
	return &AlibabaIdleParttimeSynclogAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleParttimeSynclogAPIRequest) Reset() {
	r._syncIds = r._syncIds[:0]
	r._startTime = 0
	r._endTime = 0
	r._type = 0
	r._pageSize = 0
	r._pageNum = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleParttimeSynclogAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.parttime.synclog"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleParttimeSynclogAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleParttimeSynclogAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncIds is SyncIds Setter
// 同步的id
func (r *AlibabaIdleParttimeSynclogAPIRequest) SetSyncIds(_syncIds []string) error {
	r._syncIds = _syncIds
	r.Set("sync_ids", _syncIds)
	return nil
}

// GetSyncIds SyncIds Getter
func (r AlibabaIdleParttimeSynclogAPIRequest) GetSyncIds() []string {
	return r._syncIds
}

// SetStartTime is StartTime Setter
// 查询岗位同步开始时间
func (r *AlibabaIdleParttimeSynclogAPIRequest) SetStartTime(_startTime int64) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AlibabaIdleParttimeSynclogAPIRequest) GetStartTime() int64 {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 查询岗位同步结束时间
func (r *AlibabaIdleParttimeSynclogAPIRequest) SetEndTime(_endTime int64) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabaIdleParttimeSynclogAPIRequest) GetEndTime() int64 {
	return r._endTime
}

// SetType is Type Setter
// 查询的类型, 0:岗位
func (r *AlibabaIdleParttimeSynclogAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaIdleParttimeSynclogAPIRequest) GetType() int64 {
	return r._type
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaIdleParttimeSynclogAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaIdleParttimeSynclogAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 第几页, 从0开始
func (r *AlibabaIdleParttimeSynclogAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabaIdleParttimeSynclogAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

var poolAlibabaIdleParttimeSynclogAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleParttimeSynclogRequest()
	},
}

// GetAlibabaIdleParttimeSynclogRequest 从 sync.Pool 获取 AlibabaIdleParttimeSynclogAPIRequest
func GetAlibabaIdleParttimeSynclogAPIRequest() *AlibabaIdleParttimeSynclogAPIRequest {
	return poolAlibabaIdleParttimeSynclogAPIRequest.Get().(*AlibabaIdleParttimeSynclogAPIRequest)
}

// ReleaseAlibabaIdleParttimeSynclogAPIRequest 将 AlibabaIdleParttimeSynclogAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleParttimeSynclogAPIRequest(v *AlibabaIdleParttimeSynclogAPIRequest) {
	v.Reset()
	poolAlibabaIdleParttimeSynclogAPIRequest.Put(v)
}
