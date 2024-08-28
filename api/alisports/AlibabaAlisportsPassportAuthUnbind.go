package alisports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAuthUnbind 三方关系解绑接口
// alibaba.alisports.passport.auth.unbind
//
// 解除阿里体育openId和三方合作方的openId的绑定关系
func AlibabaAlisportsPassportAuthUnbind(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAuthUnbindAPIRequest, resp *alisports.AlibabaAlisportsPassportAuthUnbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
