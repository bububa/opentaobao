package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaunitcampusspacebookinfoqueryAPIRequest 环路资源信息查询单元环境 API请求
// alibaba.unit.campus.space.bookinfo.query
//
// 环路资源信息查询单元环境
type AlibabaunitcampusspacebookinfoqueryAPIRequest struct {
	model.Params
	// work_bench_context
	_workBenchContext *WorkBenchContext
	// get_resource_book_info_request
	_getResourceBookInfoRequest *GetResourceBookInfoRequest
}

// NewAlibabaunitcampusspacebookinfoqueryRequest 初始化AlibabaunitcampusspacebookinfoqueryAPIRequest对象
func NewAlibabaunitcampusspacebookinfoqueryRequest() *AlibabaunitcampusspacebookinfoqueryAPIRequest {
	return &AlibabaunitcampusspacebookinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaunitcampusspacebookinfoqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.unit.campus.space.bookinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaunitcampusspacebookinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaunitcampusspacebookinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkBenchContext is WorkBenchContext Setter
// work_bench_context
func (r *AlibabaunitcampusspacebookinfoqueryAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabaunitcampusspacebookinfoqueryAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetGetResourceBookInfoRequest is GetResourceBookInfoRequest Setter
// get_resource_book_info_request
func (r *AlibabaunitcampusspacebookinfoqueryAPIRequest) SetGetResourceBookInfoRequest(_getResourceBookInfoRequest *GetResourceBookInfoRequest) error {
	r._getResourceBookInfoRequest = _getResourceBookInfoRequest
	r.Set("get_resource_book_info_request", _getResourceBookInfoRequest)
	return nil
}

// GetGetResourceBookInfoRequest GetResourceBookInfoRequest Getter
func (r AlibabaunitcampusspacebookinfoqueryAPIRequest) GetGetResourceBookInfoRequest() *GetResourceBookInfoRequest {
	return r._getResourceBookInfoRequest
}
