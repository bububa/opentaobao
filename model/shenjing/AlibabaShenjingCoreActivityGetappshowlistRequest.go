package shenjing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取神鲸活动列表 API请求
alibaba.shenjing.core.activity.getappshowlist

获取神鲸活动列表
*/
type AlibabaShenjingCoreActivityGetappshowlistRequest struct {
    model.Params
    // 验权对象
    workBenchContext   *WorkBenchContext
    // 时间戳
    timestamp1   int64
    // 页码
    page   int64
    // 一页行数
    size   int64
}

// 初始化AlibabaShenjingCoreActivityGetappshowlistRequest对象
func NewAlibabaShenjingCoreActivityGetappshowlistRequest() *AlibabaShenjingCoreActivityGetappshowlistRequest{
    return &AlibabaShenjingCoreActivityGetappshowlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaShenjingCoreActivityGetappshowlistRequest) GetApiMethodName() string {
    return "alibaba.shenjing.core.activity.getappshowlist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaShenjingCoreActivityGetappshowlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkBenchContext Setter
// 验权对象
func (r *AlibabaShenjingCoreActivityGetappshowlistRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaShenjingCoreActivityGetappshowlistRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}
// Timestamp1 Setter
// 时间戳
func (r *AlibabaShenjingCoreActivityGetappshowlistRequest) SetTimestamp1(timestamp1 int64) error {
    r.timestamp1 = timestamp1
    r.Set("timestamp1", timestamp1)
    return nil
}

// Timestamp1 Getter
func (r AlibabaShenjingCoreActivityGetappshowlistRequest) GetTimestamp1() int64 {
    return r.timestamp1
}
// Page Setter
// 页码
func (r *AlibabaShenjingCoreActivityGetappshowlistRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaShenjingCoreActivityGetappshowlistRequest) GetPage() int64 {
    return r.page
}
// Size Setter
// 一页行数
func (r *AlibabaShenjingCoreActivityGetappshowlistRequest) SetSize(size int64) error {
    r.size = size
    r.Set("size", size)
    return nil
}

// Size Getter
func (r AlibabaShenjingCoreActivityGetappshowlistRequest) GetSize() int64 {
    return r.size
}
