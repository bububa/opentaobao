package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoAppstoreSubscribeGet 查询appstore应用订购关系
// taobao.appstore.subscribe.get
//
// 查询appstore应用订购关系(对于新上架的多版本应用，建议使用taobao.vas.subscribe.get)
func TaobaoAppstoreSubscribeGet(ctx context.Context, clt *core.SDKClient, req *user.TaobaoAppstoreSubscribeGetAPIRequest, resp *user.TaobaoAppstoreSubscribeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
