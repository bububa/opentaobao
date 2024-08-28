package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbpromoUnbind 自促活动解绑接口
// taobao.xhotel.bnbpromo.unbind
//
// 自促活动解绑接口
func TaobaoXhotelBnbpromoUnbind(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbpromoUnbindAPIRequest, resp *xhotelitem.TaobaoXhotelBnbpromoUnbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
