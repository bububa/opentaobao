package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaServicecenterFulfiltaskCreate 合单生成核销单
// alibaba.servicecenter.fulfiltask.create
//
// 服务对工单进行合单，合单的结果是生成核销单
func AlibabaServicecenterFulfiltaskCreate(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaServicecenterFulfiltaskCreateAPIRequest, resp *tmallservice.AlibabaServicecenterFulfiltaskCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
