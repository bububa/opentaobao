package campus

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) Reset() {
	r._context = nil
	r._groupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.group.getspacegroupwithattr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// SetGroupId is GroupId Setter
// 空间单元id
func (r *AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) GetGroupId() int64 {
	return r._groupId
}

var poolAlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusSpaceGroupGetspacegroupwithattrRequest()
	},
}

// GetAlibabaCampusSpaceGroupGetspacegroupwithattrRequest 从 sync.Pool 获取 AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest
func GetAlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest() *AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest {
	return poolAlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest.Get().(*AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest)
}

// ReleaseAlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest 将 AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest(v *AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest) {
	v.Reset()
	poolAlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest.Put(v)
}
