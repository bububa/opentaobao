package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterIdentifytaskCreate 服务商创建核销单
// tmall.servicecenter.identifytask.create
//
// 服务商调用该接口进行创建核销单操作
func TmallServicecenterIdentifytaskCreate(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterIdentifytaskCreateAPIRequest, resp *tmallservice.TmallServicecenterIdentifytaskCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
