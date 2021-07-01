package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceUnitGetlistAPIRequest
多条件查询空间单元信息 API请求
alibaba.campus.space.unit.getlist

多条件查询空间单元信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getList */
type AlibabaCampusSpaceUnitGetlistAPIRequest struct {
	model.Params
	// 查询条件封装
	_param0 *WorkBenchContext
	// 查询参数封装
	_param1 *SpaceUnitQuery
}

// NewAlibabaCampusSpaceUnitGetlistRequest 初始化AlibabaCampusSpaceUnitGetlistAPIRequest对象
func NewAlibabaCampusSpaceUnitGetlistRequest() *AlibabaCampusSpaceUnitGetlistAPIRequest {
	return &AlibabaCampusSpaceUnitGetlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetlistAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetlistAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 查询条件封装
func (r *AlibabaCampusSpaceUnitGetlistAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaCampusSpaceUnitGetlistAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// Set is Param1 Setter
// 查询参数封装
func (r *AlibabaCampusSpaceUnitGetlistAPIRequest) SetParam1(_param1 *SpaceUnitQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r AlibabaCampusSpaceUnitGetlistAPIRequest) GetParam1() *SpaceUnitQuery {
	return r._param1
}
