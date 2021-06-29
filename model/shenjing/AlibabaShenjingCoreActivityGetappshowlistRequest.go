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
    _workBenchContext   *WorkBenchContext
    // 时间戳
    _timestamp1   int64
    // 页码
    _page   int64
    // 一页行数
    _size   int64
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
func (r *AlibabaShenjingCoreActivityGetappshowlistRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
    r._workBenchContext = _workBenchContext
    r.Set("work_bench_context", _workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaShenjingCoreActivityGetappshowlistRequest) GetWorkBenchContext() *WorkBenchContext {
    return r._workBenchContext
}
// Timestamp1 Setter
// 时间戳
func (r *AlibabaShenjingCoreActivityGetappshowlistRequest) SetTimestamp1(_timestamp1 int64) error {
    r._timestamp1 = _timestamp1
    r.Set("timestamp1", _timestamp1)
    return nil
}

// Timestamp1 Getter
func (r AlibabaShenjingCoreActivityGetappshowlistRequest) GetTimestamp1() int64 {
    return r._timestamp1
}
// Page Setter
// 页码
func (r *AlibabaShenjingCoreActivityGetappshowlistRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaShenjingCoreActivityGetappshowlistRequest) GetPage() int64 {
    return r._page
}
// Size Setter
// 一页行数
func (r *AlibabaShenjingCoreActivityGetappshowlistRequest) SetSize(_size int64) error {
    r._size = _size
    r.Set("size", _size)
    return nil
}

// Size Getter
func (r AlibabaShenjingCoreActivityGetappshowlistRequest) GetSize() int64 {
    return r._size
}
