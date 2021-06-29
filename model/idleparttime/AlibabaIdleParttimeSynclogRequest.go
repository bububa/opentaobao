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
    startTime   int64
    // 查询岗位同步结束时间
    endTime   int64
    // 查询的类型, 0:岗位
    type   int64
    // 页大小
    pageSize   int64
    // 第几页, 从0开始
    pageNum   int64
    // 同步的id
    syncIds   []int64
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
func (r *AlibabaIdleParttimeSynclogRequest) SetStartTime(startTime int64) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r AlibabaIdleParttimeSynclogRequest) GetStartTime() int64 {
    return r.startTime
}
// EndTime Setter
// 查询岗位同步结束时间
func (r *AlibabaIdleParttimeSynclogRequest) SetEndTime(endTime int64) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r AlibabaIdleParttimeSynclogRequest) GetEndTime() int64 {
    return r.endTime
}
// Type Setter
// 查询的类型, 0:岗位
func (r *AlibabaIdleParttimeSynclogRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaIdleParttimeSynclogRequest) GetType() int64 {
    return r.type
}
// PageSize Setter
// 页大小
func (r *AlibabaIdleParttimeSynclogRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaIdleParttimeSynclogRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNum Setter
// 第几页, 从0开始
func (r *AlibabaIdleParttimeSynclogRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

// PageNum Getter
func (r AlibabaIdleParttimeSynclogRequest) GetPageNum() int64 {
    return r.pageNum
}
// SyncIds Setter
// 同步的id
func (r *AlibabaIdleParttimeSynclogRequest) SetSyncIds(syncIds []int64) error {
    r.syncIds = syncIds
    r.Set("sync_ids", syncIds)
    return nil
}

// SyncIds Getter
func (r AlibabaIdleParttimeSynclogRequest) GetSyncIds() []int64 {
    return r.syncIds
}
