package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspacegroupgetspacegrouplistwithattrAPIRequest 分页查询空间分组业务属性 API请求
// alibaba.campus.space.group.getspacegrouplistwithattr
//
// 分页查询空间分组业务属性
type AlibabacampusspacegroupgetspacegrouplistwithattrAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 查询对象
	_groupQuery *SpaceGroupQuery
}

// NewAlibabacampusspacegroupgetspacegrouplistwithattrRequest 初始化AlibabacampusspacegroupgetspacegrouplistwithattrAPIRequest对象
func NewAlibabacampusspacegroupgetspacegrouplistwithattrRequest() *AlibabacampusspacegroupgetspacegrouplistwithattrAPIRequest {
	return &AlibabacampusspacegroupgetspacegrouplistwithattrAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspacegroupgetspacegrouplistwithattrAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.group.getspacegrouplistwithattr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspacegroupgetspacegrouplistwithattrAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspacegroupgetspacegrouplistwithattrAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 操作用户上下文
func (r *AlibabacampusspacegroupgetspacegrouplistwithattrAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabacampusspacegroupgetspacegrouplistwithattrAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// SetGroupQuery is GroupQuery Setter
// 查询对象
func (r *AlibabacampusspacegroupgetspacegrouplistwithattrAPIRequest) SetGroupQuery(_groupQuery *SpaceGroupQuery) error {
	r._groupQuery = _groupQuery
	r.Set("group_query", _groupQuery)
	return nil
}

// GetGroupQuery GroupQuery Getter
func (r AlibabacampusspacegroupgetspacegrouplistwithattrAPIRequest) GetGroupQuery() *SpaceGroupQuery {
	return r._groupQuery
}
