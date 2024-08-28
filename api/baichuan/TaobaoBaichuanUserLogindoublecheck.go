package baichuan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanUserLogindoublecheck 百川H5登录二次验证
// taobao.baichuan.user.logindoublecheck
//
// 百川H5登录二次验证
func TaobaoBaichuanUserLogindoublecheck(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanUserLogindoublecheckAPIRequest, resp *baichuan.TaobaoBaichuanUserLogindoublecheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
