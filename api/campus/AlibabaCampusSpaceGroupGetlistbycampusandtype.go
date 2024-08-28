package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceGroupGetlistbycampusandtype 根据园区id及TypeId获取空间分组
// alibaba.campus.space.group.getlistbycampusandtype
//
// 根据园区id及TypeId获取空间分组
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
// HSF方法名称：getListByCampusAndType
func AlibabaCampusSpaceGroupGetlistbycampusandtype(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest, resp *campus.AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
