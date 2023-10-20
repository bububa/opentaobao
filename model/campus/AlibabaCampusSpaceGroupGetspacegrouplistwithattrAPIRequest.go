package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest 分页查询空间分组业务属性 API请求
// alibaba.campus.space.group.getspacegrouplistwithattr
//
// 分页查询空间分组业务属性
type AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 查询对象
	_groupQuery *SpaceGroupQuery
}

// NewAlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest 初始化AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest对象
func NewAlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest() *AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest {
	return &AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) Reset() {
	r._context = nil
	r._groupQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.group.getspacegrouplistwithattr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// SetGroupQuery is GroupQuery Setter
// 查询对象
func (r *AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) SetGroupQuery(_groupQuery *SpaceGroupQuery) error {
	r._groupQuery = _groupQuery
	r.Set("group_query", _groupQuery)
	return nil
}

// GetGroupQuery GroupQuery Getter
func (r AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) GetGroupQuery() *SpaceGroupQuery {
	return r._groupQuery
}

var poolAlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest()
	},
}

// GetAlibabaCampusSpaceGroupGetspacegrouplistwithattrRequest 从 sync.Pool 获取 AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest
func GetAlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest() *AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest {
	return poolAlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest.Get().(*AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest)
}

// ReleaseAlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest 将 AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest(v *AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest) {
	v.Reset()
	poolAlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest.Put(v)
}
