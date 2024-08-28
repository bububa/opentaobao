package baichuan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountRegistercodeCheck 百川检查注册验证码
// taobao.baichuan.openaccount.registercode.check
//
// 百川检查注册验证码
func TaobaoBaichuanOpenaccountRegistercodeCheck(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
