package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetlistmapbygroupidAPIRequest 新增查询多个分组ID各自相关的空间单元信息 API请求
// alibaba.campus.space.unit.getlistmapbygroupid
//
// 新增查询多个分组ID各自相关的空间单元信息
// HSF接口名称：	com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：	getListMapByGroupIds
type AlibabaCampusSpaceUnitGetlistmapbygroupidAPIRequest struct {
	model.Params
	// 用户环境
	_param0 *WorkBenchContext
	// 查询封装
	_param1 *SpaceUnitQuery
}

// NewAlibabaCampusSpaceUnitGetlistmapbygroupidRequest 初始化AlibabaCampusSpaceUnitGetlistmapbygroupidAPIRequest对象
func NewAlibabaCampusSpaceUnitGetlistmapbygroupidRequest() *AlibabaCampusSpaceUnitGetlistmapbygroupidAPIRequest {
	return &AlibabaCampusSpaceUnitGetlistmapbygroupidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetlistmapbygroupidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getlistmapbygroupid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetlistmapbygroupidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 用户环境
func (r *AlibabaCampusSpaceUnitGetlistmapbygroupidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusSpaceUnitGetlistmapbygroupidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 查询封装
func (r *AlibabaCampusSpaceUnitGetlistmapbygroupidAPIRequest) SetParam1(_param1 *SpaceUnitQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusSpaceUnitGetlistmapbygroupidAPIRequest) GetParam1() *SpaceUnitQuery {
	return r._param1
}
