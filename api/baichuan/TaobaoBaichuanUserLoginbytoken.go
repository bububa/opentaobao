package baichuan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanUserLoginbytoken 百川手淘信任登录
// taobao.baichuan.user.loginbytoken
//
// 百川手淘信任登录
func TaobaoBaichuanUserLoginbytoken(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanUserLoginbytokenAPIRequest, resp *baichuan.TaobaoBaichuanUserLoginbytokenAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
