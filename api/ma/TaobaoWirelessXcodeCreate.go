package ma

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ma"
)

// TaobaoWirelessXcodeCreate 创建二维码/短连接
// taobao.wireless.xcode.create
//
// 创建码平台的普通二维码或者长连接转短连接服务
func TaobaoWirelessXcodeCreate(ctx context.Context, clt *core.SDKClient, req *ma.TaobaoWirelessXcodeCreateAPIRequest, resp *ma.TaobaoWirelessXcodeCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
