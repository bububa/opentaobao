package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspacegroupgetspacegroupwithattrAPIRequest 空间分组id查业务属性实例 API请求
// alibaba.campus.space.group.getspacegroupwithattr
//
// 空间分组id查业务属性实例
type AlibabacampusspacegroupgetspacegroupwithattrAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 空间单元id
	_groupId int64
}

// NewAlibabacampusspacegroupgetspacegroupwithattrRequest 初始化AlibabacampusspacegroupgetspacegroupwithattrAPIRequest对象
func NewAlibabacampusspacegroupgetspacegroupwithattrRequest() *AlibabacampusspacegroupgetspacegroupwithattrAPIRequest {
	return &AlibabacampusspacegroupgetspacegroupwithattrAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspacegroupgetspacegroupwithattrAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.group.getspacegroupwithattr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspacegroupgetspacegroupwithattrAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspacegroupgetspacegroupwithattrAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 操作用户上下文
func (r *AlibabacampusspacegroupgetspacegroupwithattrAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabacampusspacegroupgetspacegroupwithattrAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// SetGroupId is GroupId Setter
// 空间单元id
func (r *AlibabacampusspacegroupgetspacegroupwithattrAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabacampusspacegroupgetspacegroupwithattrAPIRequest) GetGroupId() int64 {
	return r._groupId
}
