package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceGetbyids 根据ids和类型查询空间列表
// alibaba.campus.space.getbyids
//
// 根据ids和类型查询空间列表
func AlibabaCampusSpaceGetbyids(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusSpaceGetbyidsAPIRequest, resp *campus.AlibabaCampusSpaceGetbyidsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
