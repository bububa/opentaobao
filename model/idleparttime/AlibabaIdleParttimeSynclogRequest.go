package idleparttime

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
兼职同步日志 API请求
alibaba.idle.parttime.synclog

提供给供应商查询的接口
*/
type AlibabaIdleParttimeSynclogRequest struct {
    model.Params
    // 查询岗位同步开始时间
    _startTime   int64
    // 查询岗位同步结束时间
    _endTime   int64
    // 查询的类型, 0:岗位
    _type   int64
    // 页大小
    _pageSize   int64
    // 第几页, 从0开始
    _pageNum   int64
    // 同步的id
    _syncIds   []int64
}

// 初始化AlibabaIdleParttimeSynclogRequest对象
func NewAlibabaIdleParttimeSynclogRequest() *AlibabaIdleParttimeSynclogRequest{
    return &AlibabaIdleParttimeSynclogRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleParttimeSynclogRequest) GetApiMethodName() string {
    return "alibaba.idle.parttime.synclog"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleParttimeSynclogRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartTime Setter
// 查询岗位同步开始时间
func (r *AlibabaIdleParttimeSynclogRequest) SetStartTime(_startTime int64) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r AlibabaIdleParttimeSynclogRequest) GetStartTime() int64 {
    return r._startTime
}
// EndTime Setter
// 查询岗位同步结束时间
func (r *AlibabaIdleParttimeSynclogRequest) SetEndTime(_endTime int64) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r AlibabaIdleParttimeSynclogRequest) GetEndTime() int64 {
    return r._endTime
}
// Type Setter
// 查询的类型, 0:岗位
func (r *AlibabaIdleParttimeSynclogRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaIdleParttimeSynclogRequest) GetType() int64 {
    return r._type
}
// PageSize Setter
// 页大小
func (r *AlibabaIdleParttimeSynclogRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaIdleParttimeSynclogRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNum Setter
// 第几页, 从0开始
func (r *AlibabaIdleParttimeSynclogRequest) SetPageNum(_pageNum int64) error {
    r._pageNum = _pageNum
    r.Set("page_num", _pageNum)
    return nil
}

// PageNum Getter
func (r AlibabaIdleParttimeSynclogRequest) GetPageNum() int64 {
    return r._pageNum
}
// SyncIds Setter
// 同步的id
func (r *AlibabaIdleParttimeSynclogRequest) SetSyncIds(_syncIds []int64) error {
    r._syncIds = _syncIds
    r.Set("sync_ids", _syncIds)
    return nil
}

// SyncIds Getter
func (r AlibabaIdleParttimeSynclogRequest) GetSyncIds() []int64 {
    return r._syncIds
}
