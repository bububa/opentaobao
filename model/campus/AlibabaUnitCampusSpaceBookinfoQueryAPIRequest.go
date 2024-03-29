package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaUnitCampusSpaceBookinfoQueryAPIRequest 环路资源信息查询单元环境 API请求
// alibaba.unit.campus.space.bookinfo.query
//
// 环路资源信息查询单元环境
type AlibabaUnitCampusSpaceBookinfoQueryAPIRequest struct {
	model.Params
	// work_bench_context
	_workBenchContext *WorkBenchContext
	// get_resource_book_info_request
	_getResourceBookInfoRequest *GetResourceBookInfoRequest
}

// NewAlibabaUnitCampusSpaceBookinfoQueryRequest 初始化AlibabaUnitCampusSpaceBookinfoQueryAPIRequest对象
func NewAlibabaUnitCampusSpaceBookinfoQueryRequest() *AlibabaUnitCampusSpaceBookinfoQueryAPIRequest {
	return &AlibabaUnitCampusSpaceBookinfoQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaUnitCampusSpaceBookinfoQueryAPIRequest) Reset() {
	r._workBenchContext = nil
	r._getResourceBookInfoRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaUnitCampusSpaceBookinfoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.unit.campus.space.bookinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaUnitCampusSpaceBookinfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaUnitCampusSpaceBookinfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkBenchContext is WorkBenchContext Setter
// work_bench_context
func (r *AlibabaUnitCampusSpaceBookinfoQueryAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabaUnitCampusSpaceBookinfoQueryAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetGetResourceBookInfoRequest is GetResourceBookInfoRequest Setter
// get_resource_book_info_request
func (r *AlibabaUnitCampusSpaceBookinfoQueryAPIRequest) SetGetResourceBookInfoRequest(_getResourceBookInfoRequest *GetResourceBookInfoRequest) error {
	r._getResourceBookInfoRequest = _getResourceBookInfoRequest
	r.Set("get_resource_book_info_request", _getResourceBookInfoRequest)
	return nil
}

// GetGetResourceBookInfoRequest GetResourceBookInfoRequest Getter
func (r AlibabaUnitCampusSpaceBookinfoQueryAPIRequest) GetGetResourceBookInfoRequest() *GetResourceBookInfoRequest {
	return r._getResourceBookInfoRequest
}

var poolAlibabaUnitCampusSpaceBookinfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaUnitCampusSpaceBookinfoQueryRequest()
	},
}

// GetAlibabaUnitCampusSpaceBookinfoQueryRequest 从 sync.Pool 获取 AlibabaUnitCampusSpaceBookinfoQueryAPIRequest
func GetAlibabaUnitCampusSpaceBookinfoQueryAPIRequest() *AlibabaUnitCampusSpaceBookinfoQueryAPIRequest {
	return poolAlibabaUnitCampusSpaceBookinfoQueryAPIRequest.Get().(*AlibabaUnitCampusSpaceBookinfoQueryAPIRequest)
}

// ReleaseAlibabaUnitCampusSpaceBookinfoQueryAPIRequest 将 AlibabaUnitCampusSpaceBookinfoQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaUnitCampusSpaceBookinfoQueryAPIRequest(v *AlibabaUnitCampusSpaceBookinfoQueryAPIRequest) {
	v.Reset()
	poolAlibabaUnitCampusSpaceBookinfoQueryAPIRequest.Put(v)
}
