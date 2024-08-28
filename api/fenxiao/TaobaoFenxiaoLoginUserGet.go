package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoLoginUserGet 获取分销用户登录信息
// taobao.fenxiao.login.user.get
//
// 获取用户登录信息
func TaobaoFenxiaoLoginUserGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoLoginUserGetAPIRequest, resp *fenxiao.TaobaoFenxiaoLoginUserGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
