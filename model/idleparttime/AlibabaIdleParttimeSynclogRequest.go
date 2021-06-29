package idleparttime

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
兼职同步日志 APIRequest
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

func NewAlibabaIdleParttimeSynclogRequest() *AlibabaIdleParttimeSynclogRequest{
    return &AlibabaIdleParttimeSynclogRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleParttimeSynclogRequest) GetApiMethodName() string {
    return "alibaba.idle.parttime.synclog"
}

func (r AlibabaIdleParttimeSynclogRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleParttimeSynclogRequest) SetStartTime(startTime int64) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r AlibabaIdleParttimeSynclogRequest) GetStartTime() int64 {
    return r.startTime
}

func (r *AlibabaIdleParttimeSynclogRequest) SetEndTime(endTime int64) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r AlibabaIdleParttimeSynclogRequest) GetEndTime() int64 {
    return r.endTime
}

func (r *AlibabaIdleParttimeSynclogRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaIdleParttimeSynclogRequest) GetType() int64 {
    return r.type
}

func (r *AlibabaIdleParttimeSynclogRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaIdleParttimeSynclogRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaIdleParttimeSynclogRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

func (r AlibabaIdleParttimeSynclogRequest) GetPageNum() int64 {
    return r.pageNum
}

func (r *AlibabaIdleParttimeSynclogRequest) SetSyncIds(syncIds []int64) error {
    r.syncIds = syncIds
    r.Set("sync_ids", syncIds)
    return nil
}

func (r AlibabaIdleParttimeSynclogRequest) GetSyncIds() []int64 {
    return r.syncIds
}

