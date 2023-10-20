package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspaceunitgetspaceunitlistwithattrAPIRequest 空间单元列表带业务属性实例 API请求
// alibaba.campus.space.unit.getspaceunitlistwithattr
//
// 空间单元列表带业务属性实例
type AlibabacampusspaceunitgetspaceunitlistwithattrAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 空间单元查询对象
	_unitQuery *SpaceUnitQuery
}

// NewAlibabacampusspaceunitgetspaceunitlistwithattrRequest 初始化AlibabacampusspaceunitgetspaceunitlistwithattrAPIRequest对象
func NewAlibabacampusspaceunitgetspaceunitlistwithattrRequest() *AlibabacampusspaceunitgetspaceunitlistwithattrAPIRequest {
	return &AlibabacampusspaceunitgetspaceunitlistwithattrAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspaceunitgetspaceunitlistwithattrAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getspaceunitlistwithattr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspaceunitgetspaceunitlistwithattrAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspaceunitgetspaceunitlistwithattrAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 操作用户上下文
func (r *AlibabacampusspaceunitgetspaceunitlistwithattrAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabacampusspaceunitgetspaceunitlistwithattrAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// SetUnitQuery is UnitQuery Setter
// 空间单元查询对象
func (r *AlibabacampusspaceunitgetspaceunitlistwithattrAPIRequest) SetUnitQuery(_unitQuery *SpaceUnitQuery) error {
	r._unitQuery = _unitQuery
	r.Set("unit_query", _unitQuery)
	return nil
}

// GetUnitQuery UnitQuery Getter
func (r AlibabacampusspaceunitgetspaceunitlistwithattrAPIRequest) GetUnitQuery() *SpaceUnitQuery {
	return r._unitQuery
}
