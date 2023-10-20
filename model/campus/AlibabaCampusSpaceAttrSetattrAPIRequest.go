package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspaceattrsetattrAPIRequest 新增业务属性实例接口 API请求
// alibaba.campus.space.attr.setattr
//
// 新增业务属性实例接口
type AlibabacampusspaceattrsetattrAPIRequest struct {
	model.Params
	// 业务属性实例集合
	_list []TypeAttrInstanceRequest
	// 操作用户上下文
	_context *WorkBenchContext
}

// NewAlibabacampusspaceattrsetattrRequest 初始化AlibabacampusspaceattrsetattrAPIRequest对象
func NewAlibabacampusspaceattrsetattrRequest() *AlibabacampusspaceattrsetattrAPIRequest {
	return &AlibabacampusspaceattrsetattrAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspaceattrsetattrAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.attr.setattr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspaceattrsetattrAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspaceattrsetattrAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetList is List Setter
// 业务属性实例集合
func (r *AlibabacampusspaceattrsetattrAPIRequest) SetList(_list []TypeAttrInstanceRequest) error {
	r._list = _list
	r.Set("list", _list)
	return nil
}

// GetList List Getter
func (r AlibabacampusspaceattrsetattrAPIRequest) GetList() []TypeAttrInstanceRequest {
	return r._list
}

// SetContext is Context Setter
// 操作用户上下文
func (r *AlibabacampusspaceattrsetattrAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabacampusspaceattrsetattrAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}
