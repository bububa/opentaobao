package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspaceunitgetspaceunitwithattrAPIRequest 空间单元id查业务属性实例 API请求
// alibaba.campus.space.unit.getspaceunitwithattr
//
// 空间单元id查业务属性实例
type AlibabacampusspaceunitgetspaceunitwithattrAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 空间单元id
	_spaceUnitId int64
}

// NewAlibabacampusspaceunitgetspaceunitwithattrRequest 初始化AlibabacampusspaceunitgetspaceunitwithattrAPIRequest对象
func NewAlibabacampusspaceunitgetspaceunitwithattrRequest() *AlibabacampusspaceunitgetspaceunitwithattrAPIRequest {
	return &AlibabacampusspaceunitgetspaceunitwithattrAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspaceunitgetspaceunitwithattrAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getspaceunitwithattr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspaceunitgetspaceunitwithattrAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspaceunitgetspaceunitwithattrAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 操作用户上下文
func (r *AlibabacampusspaceunitgetspaceunitwithattrAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabacampusspaceunitgetspaceunitwithattrAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// SetSpaceUnitId is SpaceUnitId Setter
// 空间单元id
func (r *AlibabacampusspaceunitgetspaceunitwithattrAPIRequest) SetSpaceUnitId(_spaceUnitId int64) error {
	r._spaceUnitId = _spaceUnitId
	r.Set("space_unit_id", _spaceUnitId)
	return nil
}

// GetSpaceUnitId SpaceUnitId Getter
func (r AlibabacampusspaceunitgetspaceunitwithattrAPIRequest) GetSpaceUnitId() int64 {
	return r._spaceUnitId
}
