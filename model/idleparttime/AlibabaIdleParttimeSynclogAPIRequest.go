package idleparttime

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleParttimeSynclogAPIRequest 兼职同步日志 API请求
// alibaba.idle.parttime.synclog
//
// 提供给供应商查询的接口
type AlibabaIdleParttimeSynclogAPIRequest struct {
	model.Params
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
	// 同步的id
	_syncIds []int64
}

// NewAlibabaIdleParttimeSynclogRequest 初始化AlibabaIdleParttimeSynclogAPIRequest对象
func NewAlibabaIdleParttimeSynclogRequest() *AlibabaIdleParttimeSynclogAPIRequest {
	return &AlibabaIdleParttimeSynclogAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleParttimeSynclogAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.parttime.synclog"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleParttimeSynclogAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StartTime Setter
// 查询岗位同步开始时间
func (r *AlibabaIdleParttimeSynclogAPIRequest) SetStartTime(_startTime int64) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r AlibabaIdleParttimeSynclogAPIRequest) GetStartTime() int64 {
	return r._startTime
}

// Set is EndTime Setter
// 查询岗位同步结束时间
func (r *AlibabaIdleParttimeSynclogAPIRequest) SetEndTime(_endTime int64) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r AlibabaIdleParttimeSynclogAPIRequest) GetEndTime() int64 {
	return r._endTime
}

// Set is Type Setter
// 查询的类型, 0:岗位
func (r *AlibabaIdleParttimeSynclogAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r AlibabaIdleParttimeSynclogAPIRequest) GetType() int64 {
	return r._type
}

// Set is PageSize Setter
// 页大小
func (r *AlibabaIdleParttimeSynclogAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaIdleParttimeSynclogAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is PageNum Setter
// 第几页, 从0开始
func (r *AlibabaIdleParttimeSynclogAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// Get PageNum Getter
func (r AlibabaIdleParttimeSynclogAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// Set is SyncIds Setter
// 同步的id
func (r *AlibabaIdleParttimeSynclogAPIRequest) SetSyncIds(_syncIds []int64) error {
	r._syncIds = _syncIds
	r.Set("sync_ids", _syncIds)
	return nil
}

// Get SyncIds Getter
func (r AlibabaIdleParttimeSynclogAPIRequest) GetSyncIds() []int64 {
	return r._syncIds
}
