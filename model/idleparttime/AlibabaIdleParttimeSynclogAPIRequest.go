package idleparttime

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleparttimesynclogAPIRequest 兼职同步日志 API请求
// alibaba.idle.parttime.synclog
//
// 提供给供应商查询的接口
type AlibabaidleparttimesynclogAPIRequest struct {
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

// NewAlibabaidleparttimesynclogRequest 初始化AlibabaidleparttimesynclogAPIRequest对象
func NewAlibabaidleparttimesynclogRequest() *AlibabaidleparttimesynclogAPIRequest {
	return &AlibabaidleparttimesynclogAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleparttimesynclogAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.parttime.synclog"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleparttimesynclogAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleparttimesynclogAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncIds is SyncIds Setter
// 同步的id
func (r *AlibabaidleparttimesynclogAPIRequest) SetSyncIds(_syncIds []string) error {
	r._syncIds = _syncIds
	r.Set("sync_ids", _syncIds)
	return nil
}

// GetSyncIds SyncIds Getter
func (r AlibabaidleparttimesynclogAPIRequest) GetSyncIds() []string {
	return r._syncIds
}

// SetStartTime is StartTime Setter
// 查询岗位同步开始时间
func (r *AlibabaidleparttimesynclogAPIRequest) SetStartTime(_startTime int64) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AlibabaidleparttimesynclogAPIRequest) GetStartTime() int64 {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 查询岗位同步结束时间
func (r *AlibabaidleparttimesynclogAPIRequest) SetEndTime(_endTime int64) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabaidleparttimesynclogAPIRequest) GetEndTime() int64 {
	return r._endTime
}

// SetType is Type Setter
// 查询的类型, 0:岗位
func (r *AlibabaidleparttimesynclogAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaidleparttimesynclogAPIRequest) GetType() int64 {
	return r._type
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaidleparttimesynclogAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaidleparttimesynclogAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 第几页, 从0开始
func (r *AlibabaidleparttimesynclogAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabaidleparttimesynclogAPIRequest) GetPageNum() int64 {
	return r._pageNum
}
