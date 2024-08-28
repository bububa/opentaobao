package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// AlibabaBenefitSend 发奖接口
// alibaba.benefit.send
//
// 发奖接口
func AlibabaBenefitSend(ctx context.Context, clt *core.SDKClient, req *user.AlibabaBenefitSendAPIRequest, resp *user.AlibabaBenefitSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
