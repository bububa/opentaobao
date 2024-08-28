package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtMiaolingThirdLogin 喵零第三方免登
// tmall.nrt.miaoling.third.login
//
// 喵零第三方免登
func TmallNrtMiaolingThirdLogin(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtMiaolingThirdLoginAPIRequest, resp *nrt.TmallNrtMiaolingThirdLoginAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
