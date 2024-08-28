package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceFloorGetbybuildingid 根据楼宇ID获取楼层
// alibaba.campus.space.floor.getbybuildingid
//
// 根据楼宇ID获取楼层
// HSF接口名称：com.alibaba.campus.api.space.service.top.FloorApiTopService
// HSF方法名称：getFloorList
func AlibabaCampusSpaceFloorGetbybuildingid(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusSpaceFloorGetbybuildingidAPIRequest, resp *campus.AlibabaCampusSpaceFloorGetbybuildingidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
