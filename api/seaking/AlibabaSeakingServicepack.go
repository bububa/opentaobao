package seaking

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// AlibabaSeakingServicepack 获取海王用户权限包
// alibaba.seaking.servicepack
//
// 获取海王用户权限包
func AlibabaSeakingServicepack(ctx context.Context, clt *core.SDKClient, req *seaking.AlibabaSeakingServicepackAPIRequest, resp *seaking.AlibabaSeakingServicepackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
