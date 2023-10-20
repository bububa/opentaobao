package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabashenjingcoreactivitygetappshowlistAPIRequest 获取神鲸活动列表 API请求
// alibaba.shenjing.core.activity.getappshowlist
//
// 获取神鲸活动列表
type AlibabashenjingcoreactivitygetappshowlistAPIRequest struct {
	model.Params
	// 验权对象
	_workBenchContext *WorkBenchContext
	// 时间戳
	_timestamp1 int64
	// 页码
	_page int64
	// 一页行数
	_size int64
}

// NewAlibabashenjingcoreactivitygetappshowlistRequest 初始化AlibabashenjingcoreactivitygetappshowlistAPIRequest对象
func NewAlibabashenjingcoreactivitygetappshowlistRequest() *AlibabashenjingcoreactivitygetappshowlistAPIRequest {
	return &AlibabashenjingcoreactivitygetappshowlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabashenjingcoreactivitygetappshowlistAPIRequest) GetApiMethodName() string {
	return "alibaba.shenjing.core.activity.getappshowlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabashenjingcoreactivitygetappshowlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabashenjingcoreactivitygetappshowlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkBenchContext is WorkBenchContext Setter
// 验权对象
func (r *AlibabashenjingcoreactivitygetappshowlistAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabashenjingcoreactivitygetappshowlistAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetTimestamp1 is Timestamp1 Setter
// 时间戳
func (r *AlibabashenjingcoreactivitygetappshowlistAPIRequest) SetTimestamp1(_timestamp1 int64) error {
	r._timestamp1 = _timestamp1
	r.Set("timestamp1", _timestamp1)
	return nil
}

// GetTimestamp1 Timestamp1 Getter
func (r AlibabashenjingcoreactivitygetappshowlistAPIRequest) GetTimestamp1() int64 {
	return r._timestamp1
}

// SetPage is Page Setter
// 页码
func (r *AlibabashenjingcoreactivitygetappshowlistAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabashenjingcoreactivitygetappshowlistAPIRequest) GetPage() int64 {
	return r._page
}

// SetSize is Size Setter
// 一页行数
func (r *AlibabashenjingcoreactivitygetappshowlistAPIRequest) SetSize(_size int64) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r AlibabashenjingcoreactivitygetappshowlistAPIRequest) GetSize() int64 {
	return r._size
}
