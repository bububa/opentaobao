package baichuan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountLoginbytoken 百川TOKEN 登录
// taobao.baichuan.openaccount.loginbytoken
//
// 百川TOKEN 登录
func TaobaoBaichuanOpenaccountLoginbytoken(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountLoginbytokenAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountLoginbytokenAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
