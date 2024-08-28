package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenplatformAddressGet 【商旅】开放平台对外页面跳转
// alitrip.btrip.openplatform.address.get
//
// 获取类目预定页跳转地址
func AlitripBtripOpenplatformAddressGet(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripOpenplatformAddressGetAPIRequest, resp *btrip.AlitripBtripOpenplatformAddressGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
