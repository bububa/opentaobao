package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceGroupGetbyid 根据分组ID查询相关的空间分组信息
// alibaba.campus.space.group.getbyid
//
// 根据分组ID查询相关的空间分组信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
// HSF方法名称：getById
func AlibabaCampusSpaceGroupGetbyid(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusSpaceGroupGetbyidAPIRequest, resp *campus.AlibabaCampusSpaceGroupGetbyidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
