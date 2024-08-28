package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceUnitGetlist 多条件查询空间单元信息
// alibaba.campus.space.unit.getlist
//
// 多条件查询空间单元信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getList
func AlibabaCampusSpaceUnitGetlist(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusSpaceUnitGetlistAPIRequest, resp *campus.AlibabaCampusSpaceUnitGetlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
