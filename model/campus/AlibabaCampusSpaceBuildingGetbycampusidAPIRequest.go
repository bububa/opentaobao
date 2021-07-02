package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceBuildingGetbycampusidAPIRequest 根据园区ID获取楼宇 API请求
// alibaba.campus.space.building.getbycampusid
//
// 根据园区ID获取楼宇
// HSF接口名称：com.alibaba.campus.api.space.service.top.BuildingApiTopService
// HSF方法名称：getBuildingList
type AlibabaCampusSpaceBuildingGetbycampusidAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *WorkBenchContext
	// 园区封装
	_param1 *BuildingQuery
}

// NewAlibabaCampusSpaceBuildingGetbycampusidRequest 初始化AlibabaCampusSpaceBuildingGetbycampusidAPIRequest对象
func NewAlibabaCampusSpaceBuildingGetbycampusidRequest() *AlibabaCampusSpaceBuildingGetbycampusidAPIRequest {
	return &AlibabaCampusSpaceBuildingGetbycampusidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceBuildingGetbycampusidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.building.getbycampusid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceBuildingGetbycampusidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabaCampusSpaceBuildingGetbycampusidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusSpaceBuildingGetbycampusidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 园区封装
func (r *AlibabaCampusSpaceBuildingGetbycampusidAPIRequest) SetParam1(_param1 *BuildingQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusSpaceBuildingGetbycampusidAPIRequest) GetParam1() *BuildingQuery {
	return r._param1
}
