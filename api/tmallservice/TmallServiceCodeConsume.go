package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServiceCodeConsume 天猫服务平台服务核销
// tmall.service.code.consume
//
// 天猫服务平台－服务核销
func TmallServiceCodeConsume(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServiceCodeConsumeAPIRequest, resp *tmallservice.TmallServiceCodeConsumeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
