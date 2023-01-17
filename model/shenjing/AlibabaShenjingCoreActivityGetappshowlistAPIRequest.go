package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaShenjingCoreActivityGetappshowlistAPIRequest 获取神鲸活动列表 API请求
// alibaba.shenjing.core.activity.getappshowlist
//
// 获取神鲸活动列表
type AlibabaShenjingCoreActivityGetappshowlistAPIRequest struct {
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

// NewAlibabaShenjingCoreActivityGetappshowlistRequest 初始化AlibabaShenjingCoreActivityGetappshowlistAPIRequest对象
func NewAlibabaShenjingCoreActivityGetappshowlistRequest() *AlibabaShenjingCoreActivityGetappshowlistAPIRequest {
	return &AlibabaShenjingCoreActivityGetappshowlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaShenjingCoreActivityGetappshowlistAPIRequest) GetApiMethodName() string {
	return "alibaba.shenjing.core.activity.getappshowlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaShenjingCoreActivityGetappshowlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaShenjingCoreActivityGetappshowlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkBenchContext is WorkBenchContext Setter
// 验权对象
func (r *AlibabaShenjingCoreActivityGetappshowlistAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabaShenjingCoreActivityGetappshowlistAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetTimestamp1 is Timestamp1 Setter
// 时间戳
func (r *AlibabaShenjingCoreActivityGetappshowlistAPIRequest) SetTimestamp1(_timestamp1 int64) error {
	r._timestamp1 = _timestamp1
	r.Set("timestamp1", _timestamp1)
	return nil
}

// GetTimestamp1 Timestamp1 Getter
func (r AlibabaShenjingCoreActivityGetappshowlistAPIRequest) GetTimestamp1() int64 {
	return r._timestamp1
}

// SetPage is Page Setter
// 页码
func (r *AlibabaShenjingCoreActivityGetappshowlistAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaShenjingCoreActivityGetappshowlistAPIRequest) GetPage() int64 {
	return r._page
}

// SetSize is Size Setter
// 一页行数
func (r *AlibabaShenjingCoreActivityGetappshowlistAPIRequest) SetSize(_size int64) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r AlibabaShenjingCoreActivityGetappshowlistAPIRequest) GetSize() int64 {
	return r._size
}
