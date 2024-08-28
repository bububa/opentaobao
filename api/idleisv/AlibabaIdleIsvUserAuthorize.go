package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvUserAuthorize 用户授权接口
// alibaba.idle.isv.user.authorize
//
// 用户授权接口
// 接口调用相关参考文档
// https://www.yuque.com/docs/share/9cd991b7-e3a3-40b6-948c-1835422d0164?# 《闲鱼优品2.0API接入说明》
func AlibabaIdleIsvUserAuthorize(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleIsvUserAuthorizeAPIRequest, resp *idleisv.AlibabaIdleIsvUserAuthorizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
