package alisports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAccountCheckmobile 阿里体育会员系统--手机号验证接口
// alibaba.alisports.passport.account.checkmobile
//
// 验证三方用户的手机号，获取对应的阿里体育会员id
func AlibabaAlisportsPassportAccountCheckmobile(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountCheckmobileAPIRequest, resp *alisports.AlibabaAlisportsPassportAccountCheckmobileAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
