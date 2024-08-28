package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaServicecenterIdentifytaskCreate 创建核销单
// alibaba.servicecenter.identifytask.create
//
// 创建核销单
func AlibabaServicecenterIdentifytaskCreate(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaServicecenterIdentifytaskCreateAPIRequest, resp *tmallservice.AlibabaServicecenterIdentifytaskCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
