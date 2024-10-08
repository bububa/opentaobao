package alisports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAuthAccountinfo 授权账号信息
// alibaba.alisports.passport.auth.accountinfo
//
// 获取体育用户OpenId\nick\avatar 三个信息
func AlibabaAlisportsPassportAuthAccountinfo(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAuthAccountinfoAPIRequest, resp *alisports.AlibabaAlisportsPassportAuthAccountinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
