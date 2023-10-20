package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest 空间单元列表带业务属性实例 API请求
// alibaba.campus.space.unit.getspaceunitlistwithattr
//
// 空间单元列表带业务属性实例
type AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 空间单元查询对象
	_unitQuery *SpaceUnitQuery
}

// NewAlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest 初始化AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest对象
func NewAlibabaCampusSpaceUnitGetspaceunitlistwithattrRequest() *AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest {
	return &AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getspaceunitlistwithattr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// SetUnitQuery is UnitQuery Setter
// 空间单元查询对象
func (r *AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest) SetUnitQuery(_unitQuery *SpaceUnitQuery) error {
	r._unitQuery = _unitQuery
	r.Set("unit_query", _unitQuery)
	return nil
}

// GetUnitQuery UnitQuery Getter
func (r AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest) GetUnitQuery() *SpaceUnitQuery {
	return r._unitQuery
}
