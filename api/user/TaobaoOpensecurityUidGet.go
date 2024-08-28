package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoOpensecurityUidGet 淘宝open security uid 获取接口
// taobao.opensecurity.uid.get
//
// 根据明文 taobao user id 换取 app的 open_uid
func TaobaoOpensecurityUidGet(ctx context.Context, clt *core.SDKClient, req *user.TaobaoOpensecurityUidGetAPIRequest, resp *user.TaobaoOpensecurityUidGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
