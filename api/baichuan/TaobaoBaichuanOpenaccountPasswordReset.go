package baichuan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountPasswordReset 百川找回密码
// taobao.baichuan.openaccount.password.reset
//
// 百川找回密码
func TaobaoBaichuanOpenaccountPasswordReset(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountPasswordResetAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountPasswordResetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
