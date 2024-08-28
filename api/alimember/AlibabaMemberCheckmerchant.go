package alimember

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberCheckmerchant 校验商家身份
// alibaba.member.checkmerchant
//
// 校验商家身份
func AlibabaMemberCheckmerchant(ctx context.Context, clt *core.SDKClient, req *alimember.AlibabaMemberCheckmerchantAPIRequest, resp *alimember.AlibabaMemberCheckmerchantAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
