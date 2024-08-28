package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaServicecenterWorkcardCreate 服务平台工单创建接口
// alibaba.servicecenter.workcard.create
//
// 创建服务平台工单
func AlibabaServicecenterWorkcardCreate(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaServicecenterWorkcardCreateAPIRequest, resp *tmallservice.AlibabaServicecenterWorkcardCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
