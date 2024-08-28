package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceBuildingGetbycampusid 根据园区ID获取楼宇
// alibaba.campus.space.building.getbycampusid
//
// 根据园区ID获取楼宇
// HSF接口名称：com.alibaba.campus.api.space.service.top.BuildingApiTopService
// HSF方法名称：getBuildingList
func AlibabaCampusSpaceBuildingGetbycampusid(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusSpaceBuildingGetbycampusidAPIRequest, resp *campus.AlibabaCampusSpaceBuildingGetbycampusidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
