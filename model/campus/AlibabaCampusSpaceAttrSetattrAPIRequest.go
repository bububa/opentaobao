package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceAttrSetattrAPIRequest 新增业务属性实例接口 API请求
// alibaba.campus.space.attr.setattr
//
// 新增业务属性实例接口
type AlibabaCampusSpaceAttrSetattrAPIRequest struct {
	model.Params
	// 业务属性实例集合
	_list []TypeAttrInstanceRequest
	// 操作用户上下文
	_context *WorkBenchContext
}

// NewAlibabaCampusSpaceAttrSetattrRequest 初始化AlibabaCampusSpaceAttrSetattrAPIRequest对象
func NewAlibabaCampusSpaceAttrSetattrRequest() *AlibabaCampusSpaceAttrSetattrAPIRequest {
	return &AlibabaCampusSpaceAttrSetattrAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceAttrSetattrAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.attr.setattr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceAttrSetattrAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceAttrSetattrAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetList is List Setter
// 业务属性实例集合
func (r *AlibabaCampusSpaceAttrSetattrAPIRequest) SetList(_list []TypeAttrInstanceRequest) error {
	r._list = _list
	r.Set("list", _list)
	return nil
}

// GetList List Getter
func (r AlibabaCampusSpaceAttrSetattrAPIRequest) GetList() []TypeAttrInstanceRequest {
	return r._list
}

// SetContext is Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceAttrSetattrAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabaCampusSpaceAttrSetattrAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}
