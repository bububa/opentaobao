package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest 空间分组id查业务属性实例 API请求
// alibaba.campus.space.group.getspacegroupwithattr
//
// 空间分组id查业务属性实例
type AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 空间单元id
	_groupId int64
}

// NewAlibabaCampusSpaceGroupGetspacegroupwithattrRequest 初始化AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest对象
func NewAlibabaCampusSpaceGroupGetspacegroupwithattrRequest() *AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest {
	return &AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.group.getspacegroupwithattr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// Get Context Getter
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// Set is GroupId Setter
// 空间单元id
func (r *AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// Get GroupId Getter
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) GetGroupId() int64 {
	return r._groupId
}
