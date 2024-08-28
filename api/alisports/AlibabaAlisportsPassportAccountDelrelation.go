package alisports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAccountDelrelation 阿里体育会员系统--取消三方关联接口
// alibaba.alisports.passport.account.delrelation
//
// 阿里体育会员系统--取消三方关联接口
func AlibabaAlisportsPassportAccountDelrelation(ctx context.Context, clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountDelrelationAPIRequest, resp *alisports.AlibabaAlisportsPassportAccountDelrelationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
