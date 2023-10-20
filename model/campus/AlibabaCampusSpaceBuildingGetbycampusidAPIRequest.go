package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspacebuildinggetbycampusidAPIRequest 根据园区ID获取楼宇 API请求
// alibaba.campus.space.building.getbycampusid
//
// 根据园区ID获取楼宇
// HSF接口名称：com.alibaba.campus.api.space.service.top.BuildingApiTopService
// HSF方法名称：getBuildingList
type AlibabacampusspacebuildinggetbycampusidAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *WorkBenchContext
	// 园区封装
	_param1 *BuildingQuery
}

// NewAlibabacampusspacebuildinggetbycampusidRequest 初始化AlibabacampusspacebuildinggetbycampusidAPIRequest对象
func NewAlibabacampusspacebuildinggetbycampusidRequest() *AlibabacampusspacebuildinggetbycampusidAPIRequest {
	return &AlibabacampusspacebuildinggetbycampusidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspacebuildinggetbycampusidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.building.getbycampusid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspacebuildinggetbycampusidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspacebuildinggetbycampusidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabacampusspacebuildinggetbycampusidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampusspacebuildinggetbycampusidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 园区封装
func (r *AlibabacampusspacebuildinggetbycampusidAPIRequest) SetParam1(_param1 *BuildingQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusspacebuildinggetbycampusidAPIRequest) GetParam1() *BuildingQuery {
	return r._param1
}
