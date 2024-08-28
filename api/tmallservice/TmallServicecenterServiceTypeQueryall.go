package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterServiceTypeQueryall 服务供应链服务类型
// tmall.servicecenter.service.type.queryall
//
// 查询天猫服务类型列表
func TmallServicecenterServiceTypeQueryall(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterServiceTypeQueryallAPIRequest, resp *tmallservice.TmallServicecenterServiceTypeQueryallAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
