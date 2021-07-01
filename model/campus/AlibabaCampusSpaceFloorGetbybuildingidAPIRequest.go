package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceFloorGetbybuildingidAPIRequest
根据楼宇ID获取楼层 API请求
alibaba.campus.space.floor.getbybuildingid

根据楼宇ID获取楼层
HSF接口名称：com.alibaba.campus.api.space.service.top.FloorApiTopService
HSF方法名称：getFloorList */
type AlibabaCampusSpaceFloorGetbybuildingidAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *WorkBenchContext
	// 楼宇iD封装
	_param1 *FloorQuery
}

// New
