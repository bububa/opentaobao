package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceTypeGetbycode 根据类别编码查询类别
// alibaba.campus.space.type.getbycode
//
// 根据类别编码查询类别
// HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
// HSF方法名称：getByCode
func AlibabaCampusSpaceTypeGetbycode(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusSpaceTypeGetbycodeAPIRequest, resp *campus.AlibabaCampusSpaceTypeGetbycodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
