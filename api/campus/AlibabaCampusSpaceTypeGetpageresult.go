package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceTypeGetpageresult 分页查询空间类别接口
// alibaba.campus.space.type.getpageresult
//
// 分页查询空间类别接口
// HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
// HSF方法名称：getPageResult
func AlibabaCampusSpaceTypeGetpageresult(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusSpaceTypeGetpageresultAPIRequest, resp *campus.AlibabaCampusSpaceTypeGetpageresultAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
