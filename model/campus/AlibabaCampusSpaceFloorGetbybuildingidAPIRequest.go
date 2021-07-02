package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceFloorGetbybuildingidAPIRequest 根据楼宇ID获取楼层 API请求
// alibaba.campus.space.floor.getbybuildingid
//
// 根据楼宇ID获取楼层
// HSF接口名称：com.alibaba.campus.api.space.service.top.FloorApiTopService
// HSF方法名称：getFloorList
type AlibabaCampusSpaceFloorGetbybuildingidAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *WorkBenchContext
	// 楼宇iD封装
	_param1 *FloorQuery
}

// NewAlibabaCampusSpaceFloorGetbybuildingidRequest 初始化AlibabaCampusSpaceFloorGetbybuildingidAPIRequest对象
func NewAlibabaCampusSpaceFloorGetbybuildingidRequest() *AlibabaCampusSpaceFloorGetbybuildingidAPIRequest {
	return &AlibabaCampusSpaceFloorGetbybuildingidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.floor.getbybuildingid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 楼宇iD封装
func (r *AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) SetParam1(_param1 *FloorQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusSpaceFloorGetbybuildingidAPIRequest) GetParam1() *FloorQuery {
	return r._param1
}
