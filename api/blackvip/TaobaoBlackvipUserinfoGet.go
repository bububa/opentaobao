package blackvip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/blackvip"
)

// TaobaoBlackvipUserinfoGet 88VIP用户信息查询
// taobao.blackvip.userinfo.get
//
// 查询88VIP用户信息，比如用户是否是88VIP，88VIP的失效时间等
func TaobaoBlackvipUserinfoGet(ctx context.Context, clt *core.SDKClient, req *blackvip.TaobaoBlackvipUserinfoGetAPIRequest, resp *blackvip.TaobaoBlackvipUserinfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
