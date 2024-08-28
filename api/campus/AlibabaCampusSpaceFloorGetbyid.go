package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceFloorGetbyid 根据id获取楼层
// alibaba.campus.space.floor.getbyid
//
// 根据id获取楼层
func AlibabaCampusSpaceFloorGetbyid(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusSpaceFloorGetbyidAPIRequest, resp *campus.AlibabaCampusSpaceFloorGetbyidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
