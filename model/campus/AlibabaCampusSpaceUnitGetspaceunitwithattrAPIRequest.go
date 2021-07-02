package campus

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getspaceunitwithattr"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Context Setter
// 操作用户上下文
func (r *AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) SetContext(_context *WorkBenchContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// Get Context Getter
func (r AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) GetContext() *WorkBenchContext {
	return r._context
}

// Set is SpaceUnitId Setter
// 空间单元id
func (r *AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) SetSpaceUnitId(_spaceUnitId int64) error {
	r._spaceUnitId = _spaceUnitId
	r.Set("space_unit_id", _spaceUnitId)
	return nil
}

// Get SpaceUnitId Getter
func (r AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest) GetSpaceUnitId() int64 {
	return r._spaceUnitId
}
