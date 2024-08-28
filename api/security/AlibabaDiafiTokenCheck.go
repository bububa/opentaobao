package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaDiafiTokenCheck 天朗token校验API
// alibaba.diafi.token.check
//
// 天朗token校验
func AlibabaDiafiTokenCheck(ctx context.Context, clt *core.SDKClient, req *security.AlibabaDiafiTokenCheckAPIRequest, resp *security.AlibabaDiafiTokenCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
