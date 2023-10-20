package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspaceunitgetlistmapbygroupidAPIRequest 新增查询多个分组ID各自相关的空间单元信息 API请求
// alibaba.campus.space.unit.getlistmapbygroupid
//
// 新增查询多个分组ID各自相关的空间单元信息
// HSF接口名称：	com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：	getListMapByGroupIds
type AlibabacampusspaceunitgetlistmapbygroupidAPIRequest struct {
	model.Params
	// 用户环境
	_param0 *WorkBenchContext
	// 查询封装
	_param1 *SpaceUnitQuery
}

// NewAlibabacampusspaceunitgetlistmapbygroupidRequest 初始化AlibabacampusspaceunitgetlistmapbygroupidAPIRequest对象
func NewAlibabacampusspaceunitgetlistmapbygroupidRequest() *AlibabacampusspaceunitgetlistmapbygroupidAPIRequest {
	return &AlibabacampusspaceunitgetlistmapbygroupidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspaceunitgetlistmapbygroupidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getlistmapbygroupid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspaceunitgetlistmapbygroupidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspaceunitgetlistmapbygroupidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 用户环境
func (r *AlibabacampusspaceunitgetlistmapbygroupidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampusspaceunitgetlistmapbygroupidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 查询封装
func (r *AlibabacampusspaceunitgetlistmapbygroupidAPIRequest) SetParam1(_param1 *SpaceUnitQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusspaceunitgetlistmapbygroupidAPIRequest) GetParam1() *SpaceUnitQuery {
	return r._param1
}
