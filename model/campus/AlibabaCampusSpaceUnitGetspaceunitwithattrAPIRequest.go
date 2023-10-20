package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest 空间单元id查业务属性实例 API请求
// alibaba.campus.space.unit.getspaceunitwithattr
//
// 空间单元id查业务属性实例
type AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 空间单元id
	_spaceUnitId int64
}

// NewAlibabaCampusSpaceUnitGetspaceunitwithattrRequest 初始化AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest对象
func NewAlibabaCampusSpaceUnitGetspaceunitwithattrRequest() *AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest {
	return &AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) Reset() {
	r._context = nil
	r._spaceUnitId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getspaceunitwithattr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// SetSpaceUnitId is SpaceUnitId Setter
// 空间单元id
func (r *AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) SetSpaceUnitId(_spaceUnitId int64) error {
	r._spaceUnitId = _spaceUnitId
	r.Set("space_unit_id", _spaceUnitId)
	return nil
}

// GetSpaceUnitId SpaceUnitId Getter
func (r AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) GetSpaceUnitId() int64 {
	return r._spaceUnitId
}

var poolAlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusSpaceUnitGetspaceunitwithattrRequest()
	},
}

// GetAlibabaCampusSpaceUnitGetspaceunitwithattrRequest 从 sync.Pool 获取 AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest
func GetAlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest() *AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest {
	return poolAlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest.Get().(*AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest)
}

// ReleaseAlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest 将 AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest(v *AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) {
	v.Reset()
	poolAlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest.Put(v)
}
