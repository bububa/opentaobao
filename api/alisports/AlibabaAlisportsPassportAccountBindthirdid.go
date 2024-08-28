package alisports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAccountBindthirdid 阿里体育三方ID绑定接口
// alibaba.alisports.passport.account.bindthirdid
//
// 阿里体育三方ID绑定接口
func AlibabaAlisportsPassportAccountBindthirdid(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountBindthirdidAPIRequest, resp *alisports.AlibabaAlisportsPassportAccountBindthirdidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
