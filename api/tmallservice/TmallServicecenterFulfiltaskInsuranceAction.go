package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterFulfiltaskInsuranceAction 供应链保险链路动作
// tmall.servicecenter.fulfiltask.insurance.action
//
// 服务供应链履约链路 保险类业务履约接口
func TmallServicecenterFulfiltaskInsuranceAction(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterFulfiltaskInsuranceActionAPIRequest, resp *tmallservice.TmallServicecenterFulfiltaskInsuranceActionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
