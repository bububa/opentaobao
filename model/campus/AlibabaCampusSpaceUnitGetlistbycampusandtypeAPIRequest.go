package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspaceunitgetlistbycampusandtypeAPIRequest 根据园区id及TypeId获取空间单元 API请求
// alibaba.campus.space.unit.getlistbycampusandtype
//
// 根据园区id及TypeId获取空间单元
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getListByCampusAndType
type AlibabacampusspaceunitgetlistbycampusandtypeAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *WorkBenchContext
	// 查询参数封装
	_param1 *SpaceUnitQuery
}

// NewAlibabacampusspaceunitgetlistbycampusandtypeRequest 初始化AlibabacampusspaceunitgetlistbycampusandtypeAPIRequest对象
func NewAlibabacampusspaceunitgetlistbycampusandtypeRequest() *AlibabacampusspaceunitgetlistbycampusandtypeAPIRequest {
	return &AlibabacampusspaceunitgetlistbycampusandtypeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspaceunitgetlistbycampusandtypeAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getlistbycampusandtype"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspaceunitgetlistbycampusandtypeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspaceunitgetlistbycampusandtypeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabacampusspaceunitgetlistbycampusandtypeAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampusspaceunitgetlistbycampusandtypeAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 查询参数封装
func (r *AlibabacampusspaceunitgetlistbycampusandtypeAPIRequest) SetParam1(_param1 *SpaceUnitQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusspaceunitgetlistbycampusandtypeAPIRequest) GetParam1() *SpaceUnitQuery {
	return r._param1
}
