package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceGetbyidsAPIRequest 根据ids和类型查询空间列表 API请求
// alibaba.campus.space.getbyids
//
// 根据ids和类型查询空间列表
type AlibabaCampusSpaceGetbyidsAPIRequest struct {
	model.Params
	// 上下文
	_context *WorkBenchContext
	// 查询条件
	_query *SpaceIdsQuery
}

// NewAlibabaCampusSpaceGetbyidsRequest 初始化AlibabaCampusSpaceGetbyidsAPIRequest对象
func NewAlibabaCampusSpaceGetbyidsRequest() *AlibabaCampusSpaceGetbyidsAPIRequest {
	return &AlibabaCampusSpaceGetbyidsAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusSpaceGetbyidsAPIRequest) Reset() {
	r._context = nil
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGetbyidsAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.getbyids"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGetbyidsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceGetbyidsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 上下文
func (r *AlibabaCampusSpaceGetbyidsAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabaCampusSpaceGetbyidsAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// SetQuery is Query Setter
// 查询条件
func (r *AlibabaCampusSpaceGetbyidsAPIRequest) SetQuery(_query *SpaceIdsQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaCampusSpaceGetbyidsAPIRequest) GetQuery() *SpaceIdsQuery {
	return r._query
}

var poolAlibabaCampusSpaceGetbyidsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusSpaceGetbyidsRequest()
	},
}

// GetAlibabaCampusSpaceGetbyidsRequest 从 sync.Pool 获取 AlibabaCampusSpaceGetbyidsAPIRequest
func GetAlibabaCampusSpaceGetbyidsAPIRequest() *AlibabaCampusSpaceGetbyidsAPIRequest {
	return poolAlibabaCampusSpaceGetbyidsAPIRequest.Get().(*AlibabaCampusSpaceGetbyidsAPIRequest)
}

// ReleaseAlibabaCampusSpaceGetbyidsAPIRequest 将 AlibabaCampusSpaceGetbyidsAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusSpaceGetbyidsAPIRequest(v *AlibabaCampusSpaceGetbyidsAPIRequest) {
	v.Reset()
	poolAlibabaCampusSpaceGetbyidsAPIRequest.Put(v)
}
