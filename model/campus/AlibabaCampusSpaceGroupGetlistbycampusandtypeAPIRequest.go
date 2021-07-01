package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest
根据园区id及TypeId获取空间分组 API请求
alibaba.campus.space.group.getlistbycampusandtype

根据园区id及TypeId获取空间分组
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
HSF方法名称：getListByCampusAndType */
type AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *WorkBenchContext
	// 查询参数封装
	_param1 *SpaceGroupQuery
}

// NewAlibabaCampusSpaceGroupGetlistbycampusandtypeRequest 初始化AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest对象
func NewAlibabaCampusSpaceGroupGetlistbycampusandtypeRequest() *AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest {
	return &AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.group.getlistbycampusandtype"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 系统自动生成
func (r *AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// Set is Param1 Setter
// 查询参数封装
func (r *AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) SetParam1(_param1 *SpaceGroupQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) GetParam1() *SpaceGroupQuery {
	return r._param1
}
