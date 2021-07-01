package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceBuildingGetbycampusidAPIRequest
根据园区ID获取楼宇 API请求
alibaba.campus.space.building.getbycampusid

根据园区ID获取楼宇
HSF接口名称：com.alibaba.campus.api.space.service.top.BuildingApiTopService
HSF方法名称：getBuildingList */
type AlibabaCampusSpaceBuildingGetbycampusidAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *WorkBenchContext
	// 园区封装
	_param1 *BuildingQuery
}

// New
